package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	authorizationURL = "https://lucid.app/oauth2/authorize"
	tokenURL         = "https://api.lucid.co/oauth2/token"
	refreshBuffer    = 5 * time.Minute
)

type TokenStore struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	TokenType    string    `json:"token_type"`
}

type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scopes       []string
	TokenFile    string
}

type OAuthClient struct {
	config     OAuthConfig
	store      *TokenStore
	mu         sync.RWMutex
	httpClient *http.Client
}

func NewOAuthClient(config OAuthConfig) *OAuthClient {
	return &OAuthClient{
		config:     config,
		store:      &TokenStore{},
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (o *OAuthClient) LoadTokens() error {
	data, err := os.ReadFile(o.config.TokenFile)
	if err != nil {
		return fmt.Errorf("failed to read token file: %w", err)
	}

	var store TokenStore
	if err := json.Unmarshal(data, &store); err != nil {
		return fmt.Errorf("failed to parse token file: %w", err)
	}

	if store.RefreshToken == "" {
		return fmt.Errorf("token file contains no refresh token")
	}

	o.mu.Lock()
	o.store = &store
	o.mu.Unlock()
	return nil
}

func (o *OAuthClient) SaveTokens() error {
	o.mu.RLock()
	data, err := json.MarshalIndent(o.store, "", "  ")
	o.mu.RUnlock()
	if err != nil {
		return fmt.Errorf("failed to marshal tokens: %w", err)
	}

	if err := os.WriteFile(o.config.TokenFile, data, 0600); err != nil {
		return fmt.Errorf("failed to write token file: %w", err)
	}
	return nil
}

func (o *OAuthClient) NeedsRefresh() bool {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.store.AccessToken == "" || time.Now().Add(refreshBuffer).After(o.store.ExpiresAt)
}

func (o *OAuthClient) GetValidToken(ctx context.Context) (string, error) {
	if !o.NeedsRefresh() {
		o.mu.RLock()
		defer o.mu.RUnlock()
		return o.store.AccessToken, nil
	}

	if err := o.RefreshAccessToken(ctx); err != nil {
		return "", fmt.Errorf("failed to refresh token: %w", err)
	}

	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.store.AccessToken, nil
}

func (o *OAuthClient) RefreshAccessToken(ctx context.Context) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	// Double-check under write lock
	if o.store.AccessToken != "" && time.Now().Add(refreshBuffer).Before(o.store.ExpiresAt) {
		return nil
	}

	if o.store.RefreshToken == "" {
		return fmt.Errorf("no refresh token available, re-authorization required")
	}

	payload := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": o.store.RefreshToken,
		"client_id":     o.config.ClientID,
		"client_secret": o.config.ClientSecret,
	}

	return o.exchangeToken(ctx, payload)
}

// RunAuthorizationFlow performs the interactive OAuth authorization code flow.
// It prints a URL for the user to open, starts a temporary callback server,
// waits for the authorization code, and exchanges it for tokens.
func (o *OAuthClient) RunAuthorizationFlow(ctx context.Context) error {
	state, err := randomState()
	if err != nil {
		return fmt.Errorf("failed to generate state: %w", err)
	}

	// Build authorization URL
	params := url.Values{
		"client_id":     {o.config.ClientID},
		"redirect_uri":  {o.config.RedirectURI},
		"scope":         {strings.Join(o.config.Scopes, " ")},
		"state":         {state},
		"response_type": {"code"},
	}
	authURL := authorizationURL + "?" + params.Encode()

	fmt.Println("\n=== Lucidchart OAuth Authorization ===")
	fmt.Println("Open this URL in your browser to authorize:")
	fmt.Printf("\n  %s\n\n", authURL)
	fmt.Println("Waiting for authorization callback...")

	// Parse redirect URI to get host:port for the callback server
	parsed, err := url.Parse(o.config.RedirectURI)
	if err != nil {
		return fmt.Errorf("invalid redirect URI: %w", err)
	}

	codeCh := make(chan string, 1)
	errCh := make(chan error, 1)

	mux := http.NewServeMux()
	mux.HandleFunc(parsed.Path, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			errCh <- fmt.Errorf("state mismatch: expected %s, got %s", state, r.URL.Query().Get("state"))
			http.Error(w, "State mismatch", http.StatusBadRequest)
			return
		}

		if errParam := r.URL.Query().Get("error"); errParam != "" {
			errCh <- fmt.Errorf("authorization error: %s - %s", errParam, r.URL.Query().Get("error_description"))
			http.Error(w, "Authorization failed", http.StatusBadRequest)
			return
		}

		code := r.URL.Query().Get("code")
		if code == "" {
			errCh <- fmt.Errorf("no authorization code received")
			http.Error(w, "No code received", http.StatusBadRequest)
			return
		}

		fmt.Fprintln(w, "Authorization successful! You can close this window.")
		codeCh <- code
	})

	listener, err := net.Listen("tcp", parsed.Host)
	if err != nil {
		return fmt.Errorf("failed to start callback server: %w", err)
	}

	server := &http.Server{Handler: mux}
	go func() {
		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			errCh <- fmt.Errorf("callback server error: %w", err)
		}
	}()

	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(shutdownCtx)
	}()

	// Wait for code or error
	select {
	case code := <-codeCh:
		log.Println("Authorization code received, exchanging for tokens...")
		payload := map[string]string{
			"grant_type":    "authorization_code",
			"code":          code,
			"client_id":     o.config.ClientID,
			"client_secret": o.config.ClientSecret,
			"redirect_uri":  o.config.RedirectURI,
		}

		o.mu.Lock()
		defer o.mu.Unlock()
		if err := o.exchangeToken(ctx, payload); err != nil {
			return err
		}
		// SaveTokens needs the lock released, so we save directly here
		data, err := json.MarshalIndent(o.store, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal tokens: %w", err)
		}
		return os.WriteFile(o.config.TokenFile, data, 0600)

	case err := <-errCh:
		return err

	case <-ctx.Done():
		return ctx.Err()
	}
}

// exchangeToken calls the token endpoint and updates the store.
// Caller must hold o.mu write lock.
func (o *OAuthClient) exchangeToken(ctx context.Context, payload map[string]string) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal token request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", tokenURL, strings.NewReader(string(body)))
	if err != nil {
		return fmt.Errorf("failed to create token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := o.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("token request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read token response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("token endpoint returned %d: %s", resp.StatusCode, string(respBody))
	}

	var tokenResp struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
		TokenType    string `json:"token_type"`
	}
	if err := json.Unmarshal(respBody, &tokenResp); err != nil {
		return fmt.Errorf("failed to parse token response: %w", err)
	}

	o.store.AccessToken = tokenResp.AccessToken
	o.store.RefreshToken = tokenResp.RefreshToken
	o.store.ExpiresAt = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
	o.store.TokenType = tokenResp.TokenType

	log.Println("Tokens updated successfully")
	return nil
}

func randomState() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
