package asteroids

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"strings"
)

func getHighScore() (int, error) {
	// Gte the user name
	u, err := user.Current()
	if err != nil {
		return 0, err
	}

	path := ""
	switch runtime.GOOS {
	case "darwin":
		path = fmt.Sprintf("/Users/%s/Library/Application Support/Asteroids", u.Username)
	case "Windows":
		path = fmt.Sprintf("C:\\Users\\%s\\AppData\\Roaming\\Asteroids", u.Username)
	case "linux":
		path = fmt.Sprintf("/home/%s/.config/Asteroids", u.Username)
	default:
		return 0, fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	if _, err := os.Stat(path); err != nil {
		if err := os.Mkdir(path, 0750); err != nil {
			return 0, fmt.Errorf("failed to create directory: %v", err)
		}
	}

	if _,err := os.Stat(path + "/high-score.txt"); err != nil {
		err := os.WriteFile(path + "/high-score.txt", []byte("0"), 0750)
		if err != nil {
			return 0, fmt.Errorf("failed to create high-score file: %v", err)
		}
	}

	contents, err := os.ReadFile(path + "/high-score.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read high-score file: %v", err)
	}
	score := string(contents)
	score = strings.TrimSpace(score)

	s, err := strconv.Atoi(string(score))
	if err != nil {
		return 0, fmt.Errorf("failed to convert high-score to int: %v", err)
	}

	return s, nil
}

func updateHighScore(score int) error {
	// Gte the user name
	u, err := user.Current()
	if err != nil {
		return err
	}

	path := ""
	switch runtime.GOOS {
	case "darwin":
		path = fmt.Sprintf("/Users/%s/Library/Application Support/Asteroids/high-score.txt", u.Username)
	case "Windows":
		path = fmt.Sprintf("C:\\Users\\%s\\AppData\\Roaming\\Asteroids\\high-score.txt", u.Username)
	case "linux":
		path = fmt.Sprintf("/home/%s/.config/Asteroids/high-score.txt", u.Username)
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	s := fmt.Sprintf("%d", score)
	if err := os.WriteFile(path, []byte(s), 0750); err != nil {
		return fmt.Errorf("failed to write high-score file: %v", err)
	}
	
	return nil
}	
