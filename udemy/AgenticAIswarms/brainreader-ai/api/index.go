package handler

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
    "time"
)

type ContactForm struct {
    Name    string `json:"name"`
    Email   string `json:"email"`
    Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" || r.URL.Path == "" {
        homeHandler(w, r)
    } else if r.URL.Path == "/api/contact" {
        contactHandler(w, r)
    } else if r.URL.Path == "/founders" {
        foundersHandler(w, r)
    } else if r.URL.Path == "/advisors" {
        advisorsHandler(w, r)
    } else if r.URL.Path == "/investors" {
        investorsHandler(w, r)
    } else if r.URL.Path == "/admin" {
        adminHandler(w, r)
    } else {
        http.NotFound(w, r)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, indexHTML)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var form ContactForm
    if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Save to Vercel Blob
    if err := saveToBlob(form); err != nil {
        fmt.Printf("Blob storage error: %v\n", err)
    }

    fmt.Printf("Contact form submission: %+v\n", form)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "status": "success",
        "message": "Thank you for contacting Brainreader AI",
    })
}

func foundersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, foundersHTML)
}

func advisorsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, advisorsHTML)
}

func investorsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, investorsHTML)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    // Get submissions from blob storage
    submissions, err := getSubmissions()
    if err != nil {
        fmt.Printf("Error getting submissions: %v\n", err)
        submissions = []ContactForm{} // Empty if error
    }

    // Generate submissions HTML
    submissionsHTML := ""
    if len(submissions) == 0 {
        submissionsHTML = `<div class="no-data">
            <p>No submissions yet. Make sure:</p>
            <ul>
                <li>BLOB_READ_WRITE_TOKEN is set in Environment Variables</li>
                <li>Someone has submitted the contact form</li>
            </ul>
        </div>`
    } else {
        for i, sub := range submissions {
            submissionsHTML += fmt.Sprintf(`
            <div class="submission">
                <h3>Submission #%d</h3>
                <p><strong>Name:</strong> %s</p>
                <p><strong>Email:</strong> %s</p>
                <p><strong>Message:</strong> %s</p>
            </div>`, i+1, sub.Name, sub.Email, sub.Message)
        }
    }

    html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>Admin - Contact Submissions</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; background: #f5f5f5; }
        .container { max-width: 800px; margin: 0 auto; background: white; padding: 30px; border-radius: 10px; }
        h1 { color: #333; }
        .submission { border: 1px solid #ddd; padding: 15px; margin: 10px 0; border-radius: 5px; background: #f9f9f9; }
        .no-data { background: #fff3cd; padding: 15px; border-radius: 5px; border: 1px solid #ffeaa7; }
        .back-link { color: #1976d2; text-decoration: none; }
    </style>
</head>
<body>
    <div class="container">
        <a href="/" class="back-link">← Back to Home</a>
        <h1>Contact Submissions</h1>
        %s
    </div>
</body>
</html>`, submissionsHTML)

    fmt.Fprint(w, html)
}

func getSubmissions() ([]ContactForm, error) {
    blobToken := os.Getenv("BLOB_READ_WRITE_TOKEN")
    if blobToken == "" {
        return nil, fmt.Errorf("BLOB_READ_WRITE_TOKEN not configured")
    }

    // List files in contacts folder
    listURL := "https://blob.vercel-storage.com/list?prefix=contacts/"

    req, err := http.NewRequest("GET", listURL, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Authorization", "Bearer "+blobToken)

    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("failed to list files: %d", resp.StatusCode)
    }

    // Parse response to get file URLs
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var listResp struct {
        Blobs []struct {
            URL string `json:"url"`
        } `json:"blobs"`
    }

    if err := json.Unmarshal(body, &listResp); err != nil {
        return nil, err
    }

    // Fetch each file and parse
    var submissions []ContactForm
    for _, blob := range listResp.Blobs {
        if strings.Contains(blob.URL, "contact-") {
            submission, err := fetchSubmission(blob.URL, blobToken)
            if err == nil {
                submissions = append(submissions, submission)
            }
        }
    }

    return submissions, nil
}

func fetchSubmission(url, token string) (ContactForm, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return ContactForm{}, err
    }

    req.Header.Set("Authorization", "Bearer "+token)

    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return ContactForm{}, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return ContactForm{}, err
    }

    var submission ContactForm
    if err := json.Unmarshal(body, &submission); err != nil {
        return ContactForm{}, err
    }

    return submission, nil
}

func saveToBlob(form ContactForm) error {
    blobToken := os.Getenv("BLOB_READ_WRITE_TOKEN")

    if blobToken == "" {
        return fmt.Errorf("Blob credentials not configured")
    }

    // Create unique filename with timestamp
    filename := fmt.Sprintf("contacts/contact-%d.json", time.Now().Unix())

    // Convert form to JSON
    formData, err := json.Marshal(form)
    if err != nil {
        return err
    }

    // Prepare blob upload request
    url := fmt.Sprintf("https://blob.vercel-storage.com/%s", filename)

    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(formData))
    if err != nil {
        return err
    }

    req.Header.Set("Authorization", "Bearer "+blobToken)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        return fmt.Errorf("Blob upload failed with status: %d", resp.StatusCode)
    }

    return nil
}

const indexHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Brainreader AI - Understand and Control Neural Syntax</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: #0a0a0a url('/chalkboard.png') center center;
            background-size: cover;
            background-attachment: fixed;
            color: #fff;
            min-height: 100vh;
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            padding: 2rem;
        }

        .container {
            max-width: 800px;
            text-align: center;
            width: 100%;
        }

        h1 {
            font-size: 4rem;
            font-weight: 300;
            letter-spacing: 0.2em;
            margin-bottom: 1rem;
            background: linear-gradient(45deg, #fff, #aaa);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        .tagline {
            font-size: 1.5rem;
            color: #888;
            margin-bottom: 3rem;
        }

        .links {
            display: flex;
            gap: 1rem;
            justify-content: center;
            margin-bottom: 3rem;
            flex-wrap: wrap;
        }

        .link {
            padding: 0.75rem 2rem;
            border: 1px solid #444;
            border-radius: 2rem;
            text-decoration: none;
            color: #fff;
            transition: all 0.3s;
        }

        .link:hover {
            background: #222;
            border-color: #666;
        }

        .social {
            display: flex;
            gap: 1.5rem;
            justify-content: center;
            margin-bottom: 3rem;
        }

        .social a {
            width: 40px;
            height: 40px;
            border: 1px solid #444;
            border-radius: 50%;
            display: flex;
            align-items: center;
            justify-content: center;
            text-decoration: none;
            color: #fff;
            transition: all 0.3s;
        }

        .social a:hover {
            background: #222;
            border-color: #666;
        }

        .partners {
            display: flex;
            gap: 2rem;
            justify-content: center;
            align-items: center;
            margin-bottom: 3rem;
            flex-wrap: wrap;
        }

        .partner {
            padding: 0.5rem 1.5rem;
            background: #fff;
            border-radius: 0.5rem;
            text-decoration: none;
            color: #000;
            font-weight: 500;
            transition: transform 0.3s;
        }

        .partner:hover {
            transform: scale(1.05);
        }

        .partner.yc {
            background: #ff6600;
            color: #fff;
        }

        .partner.nvidia {
            background: #76b900;
        }

        .contact-section {
            margin-top: 4rem;
            padding-top: 2rem;
            border-top: 1px solid #333;
        }

        .contact-form {
            display: flex;
            flex-direction: column;
            gap: 1rem;
            max-width: 400px;
            margin: 0 auto;
        }

        .contact-form input,
        .contact-form textarea {
            padding: 0.75rem;
            background: rgba(255, 255, 255, 0.1);
            border: 1px solid #444;
            border-radius: 0.5rem;
            color: #fff;
            font-family: inherit;
        }

        .contact-form input::placeholder,
        .contact-form textarea::placeholder {
            color: #888;
        }

        .contact-form button {
            padding: 0.75rem 2rem;
            background: linear-gradient(45deg, #6a0dad, #4b0082);
            border: none;
            border-radius: 0.5rem;
            color: #fff;
            cursor: pointer;
            font-size: 1rem;
            transition: transform 0.3s;
        }

        .contact-form button:hover {
            transform: scale(1.05);
        }

        .message {
            margin-top: 1rem;
            padding: 0.75rem;
            border-radius: 0.5rem;
            display: none;
        }

        .message.success {
            background: rgba(76, 175, 80, 0.2);
            border: 1px solid #4caf50;
            color: #4caf50;
        }

        .message.error {
            background: rgba(244, 67, 54, 0.2);
            border: 1px solid #f44336;
            color: #f44336;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>BRAINREADER AI</h1>
        <p class="tagline">Understand and Control Neural Syntax</p>

        <div class="links">
            <a href="/founders" class="link">Founders</a>
            <a href="/advisors" class="link">Advisors</a>
            <a href="/investors" class="link">Investors</a>
        </div>

        <div class="social">
            <a href="https://twitter.com" target="_blank">
                <svg width="20" height="20" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M23 3a10.9 10.9 0 01-3.14 1.53 4.48 4.48 0 00-7.86 3v1A10.66 10.66 0 013 4s-4 9 5 13a11.64 11.64 0 01-7 2c9 5 20 0 20-11.5a4.5 4.5 0 00-.08-.83A7.72 7.72 0 0023 3z"/>
                </svg>
            </a>
            <a href="https://linkedin.com" target="_blank">
                <svg width="20" height="20" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M16 8a6 6 0 016 6v7h-4v-7a2 2 0 00-2-2 2 2 0 00-2 2v7h-4v-7a6 6 0 016-6zM2 9h4v12H2z"/>
                    <circle cx="4" cy="4" r="2"/>
                </svg>
            </a>
        </div>

        <div class="partners">
            <a href="https://www.ycombinator.com" target="_blank" class="partner yc">Y Combinator</a>
            <a href="https://www.nvidia.com" target="_blank" class="partner nvidia">NVIDIA Inception</a>
        </div>

        <div class="contact-section">
            <h2 style="margin-bottom: 1.5rem;">Get in Touch</h2>
            <form class="contact-form" id="contactForm">
                <input type="text" name="name" placeholder="Your Name" required>
                <input type="email" name="email" placeholder="Your Email" required>
                <textarea name="message" rows="4" placeholder="Your Message" required></textarea>
                <button type="submit">Send Message</button>
            </form>
            <div id="message" class="message"></div>
        </div>
    </div>

    <script>
        document.getElementById('contactForm').addEventListener('submit', async (e) => {
            e.preventDefault();

            const formData = new FormData(e.target);
            const data = Object.fromEntries(formData);
            const messageDiv = document.getElementById('message');

            try {
                const response = await fetch('/api/contact', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(data),
                });

                const result = await response.json();

                if (response.ok) {
                    messageDiv.textContent = result.message;
                    messageDiv.className = 'message success';
                    messageDiv.style.display = 'block';
                    e.target.reset();
                } else {
                    throw new Error('Failed to send message');
                }
            } catch (error) {
                messageDiv.textContent = 'Error sending message. Please try again.';
                messageDiv.className = 'message error';
                messageDiv.style.display = 'block';
            }

            setTimeout(() => {
                messageDiv.style.display = 'none';
            }, 5000);
        });
    </script>
</body>
</html>`

const foundersHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Founders - Brainreader AI</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: #0a0a0a;
            color: #fff;
            min-height: 100vh;
            padding: 2rem;
        }

        .container {
            max-width: 1000px;
            margin: 0 auto;
        }

        .header {
            text-align: center;
            margin-bottom: 3rem;
        }

        h1 {
            font-size: 3rem;
            font-weight: 300;
            letter-spacing: 0.2em;
            margin-bottom: 1rem;
            background: linear-gradient(45deg, #fff, #aaa);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        .back-link {
            display: inline-block;
            margin-bottom: 2rem;
            padding: 0.5rem 1rem;
            border: 1px solid #444;
            border-radius: 1rem;
            text-decoration: none;
            color: #fff;
            transition: all 0.3s;
        }

        .back-link:hover {
            background: #222;
            border-color: #666;
        }

        .founders-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
            gap: 2rem;
            margin-bottom: 3rem;
        }

        .founder-card {
            background: rgba(255, 255, 255, 0.1);
            border: 1px solid rgba(255, 255, 255, 0.2);
            border-radius: 1rem;
            padding: 2rem;
            text-align: center;
        }

        .founder-name {
            font-size: 1.5rem;
            font-weight: 600;
            margin-bottom: 0.5rem;
        }

        .founder-title {
            color: #aaa;
            margin-bottom: 1rem;
        }

        .founder-bio {
            line-height: 1.6;
        }

        .add-founder {
            background: rgba(76, 175, 80, 0.2);
            border: 2px dashed rgba(76, 175, 80, 0.5);
            display: flex;
            align-items: center;
            justify-content: center;
            min-height: 200px;
            cursor: pointer;
            transition: all 0.3s;
        }

        .add-founder:hover {
            background: rgba(76, 175, 80, 0.3);
        }

        .add-founder-text {
            color: #4caf50;
            font-size: 1.2rem;
        }
    </style>
</head>
<body>
    <div class="container">
        <a href="/" class="back-link">← Back to Home</a>

        <div class="header">
            <h1>FOUNDERS</h1>
        </div>

        <div class="founders-grid">
            <!-- Example founder - students can edit this -->
            <div class="founder-card">
                <div class="founder-name">Your Name Here</div>
                <div class="founder-title">Co-Founder & CEO</div>
                <div class="founder-bio">
                    Add your bio here. Describe your background, expertise, and vision for Brainreader AI.
                </div>
            </div>

            <!-- Add more founders button -->
            <div class="founder-card add-founder">
                <div class="add-founder-text">+ Add Founder</div>
            </div>
        </div>
    </div>
</body>
</html>`

const advisorsHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Advisors - Brainreader AI</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: #0a0a0a;
            color: #fff;
            min-height: 100vh;
            padding: 2rem;
        }

        .container {
            max-width: 1000px;
            margin: 0 auto;
        }

        .header {
            text-align: center;
            margin-bottom: 3rem;
        }

        h1 {
            font-size: 3rem;
            font-weight: 300;
            letter-spacing: 0.2em;
            margin-bottom: 1rem;
            background: linear-gradient(45deg, #fff, #aaa);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        .back-link {
            display: inline-block;
            margin-bottom: 2rem;
            padding: 0.5rem 1rem;
            border: 1px solid #444;
            border-radius: 1rem;
            text-decoration: none;
            color: #fff;
            transition: all 0.3s;
        }

        .back-link:hover {
            background: #222;
            border-color: #666;
        }

        .advisors-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
            gap: 2rem;
            margin-bottom: 3rem;
        }

        .advisor-card {
            background: rgba(255, 255, 255, 0.1);
            border: 1px solid rgba(255, 255, 255, 0.2);
            border-radius: 1rem;
            padding: 2rem;
            text-align: center;
        }

        .advisor-name {
            font-size: 1.5rem;
            font-weight: 600;
            margin-bottom: 0.5rem;
        }

        .advisor-title {
            color: #aaa;
            margin-bottom: 1rem;
        }

        .advisor-bio {
            line-height: 1.6;
        }

        .add-advisor {
            background: rgba(33, 150, 243, 0.2);
            border: 2px dashed rgba(33, 150, 243, 0.5);
            display: flex;
            align-items: center;
            justify-content: center;
            min-height: 200px;
            cursor: pointer;
            transition: all 0.3s;
        }

        .add-advisor:hover {
            background: rgba(33, 150, 243, 0.3);
        }

        .add-advisor-text {
            color: #2196f3;
            font-size: 1.2rem;
        }
    </style>
</head>
<body>
    <div class="container">
        <a href="/" class="back-link">← Back to Home</a>

        <div class="header">
            <h1>ADVISORS</h1>
        </div>

        <div class="advisors-grid">
            <!-- Example advisor - students can edit this -->
            <div class="advisor-card">
                <div class="advisor-name">Advisor Name</div>
                <div class="advisor-title">Former CTO at Tech Company</div>
                <div class="advisor-bio">
                    Add advisor bio here. Include their experience, achievements, and how they help guide Brainreader AI.
                </div>
            </div>

            <!-- Add more advisors button -->
            <div class="advisor-card add-advisor">
                <div class="add-advisor-text">+ Add Advisor</div>
            </div>
        </div>
    </div>
</body>
</html>`

const investorsHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Investors - Brainreader AI</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: #0a0a0a;
            color: #fff;
            min-height: 100vh;
            padding: 2rem;
        }

        .container {
            max-width: 1000px;
            margin: 0 auto;
        }

        .header {
            text-align: center;
            margin-bottom: 3rem;
        }

        h1 {
            font-size: 3rem;
            font-weight: 300;
            letter-spacing: 0.2em;
            margin-bottom: 1rem;
            background: linear-gradient(45deg, #fff, #aaa);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        .back-link {
            display: inline-block;
            margin-bottom: 2rem;
            padding: 0.5rem 1rem;
            border: 1px solid #444;
            border-radius: 1rem;
            text-decoration: none;
            color: #fff;
            transition: all 0.3s;
        }

        .back-link:hover {
            background: #222;
            border-color: #666;
        }

        .investors-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
            gap: 2rem;
            margin-bottom: 3rem;
        }

        .investor-card {
            background: rgba(255, 255, 255, 0.1);
            border: 1px solid rgba(255, 255, 255, 0.2);
            border-radius: 1rem;
            padding: 2rem;
            text-align: center;
        }

        .investor-name {
            font-size: 1.5rem;
            font-weight: 600;
            margin-bottom: 0.5rem;
        }

        .investor-title {
            color: #aaa;
            margin-bottom: 1rem;
        }

        .investor-bio {
            line-height: 1.6;
        }

        .add-investor {
            background: rgba(255, 152, 0, 0.2);
            border: 2px dashed rgba(255, 152, 0, 0.5);
            display: flex;
            align-items: center;
            justify-content: center;
            min-height: 200px;
            cursor: pointer;
            transition: all 0.3s;
        }

        .add-investor:hover {
            background: rgba(255, 152, 0, 0.3);
        }

        .add-investor-text {
            color: #ff9800;
            font-size: 1.2rem;
        }
    </style>
</head>
<body>
    <div class="container">
        <a href="/" class="back-link">← Back to Home</a>

        <div class="header">
            <h1>INVESTORS</h1>
        </div>

        <div class="investors-grid">
            <!-- Example investor - students can edit this -->
            <div class="investor-card">
                <div class="investor-name">Investor Name</div>
                <div class="investor-title">Partner at VC Fund</div>
                <div class="investor-bio">
                    Add investor bio here. Include their fund, investment focus, and belief in Brainreader AI's mission.
                </div>
            </div>

            <!-- Add more investors button -->
            <div class="investor-card add-investor">
                <div class="add-investor-text">+ Add Investor</div>
            </div>
        </div>
    </div>
</body>
</html>`