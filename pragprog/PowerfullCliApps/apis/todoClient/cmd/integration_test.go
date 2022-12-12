//go:build integration
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

func randomTaskName(t *testing.T) string {
	t.Helper()
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var p strings.Builder
	for i := 0; i < 32; i++ {
		p.WriteByte(chars[r.Intn(len(chars))])
	}

	return p.String()
}

func TestIntegration(t *testing.T) {
	apiRoot := "http://localhost:8080"

	if os.Getenv("TODO_API") != "" {
		apiRoot = os.Getenv("TODO_API")
	}
	today := time.Now().Format("Jan/02")
	task := randomTaskName(t)
	taskId := ""

	t.Run("AddTask", func(t *testing.T) {
		args := []string{task}
		expOut := fmt.Sprintf("Added task %q to your list.\n", task)
		var out bytes.Buffer

		if err := addAction(&out, apiRoot, args); err != nil {
			t.Fatalf("Expected no error, got %q", err)
		}
		if expOut != out.String() {
			t.Errorf("Expected output: %q, got: %q", expOut, out.String())
		}
	})
	t.Run("ListTask", func(t *testing.T) {
		var out bytes.Buffer
		if err := listAction(&out, apiRoot); err != nil {
			t.Fatalf("Expected no error, got %q", err)
		}
		outList := ""
		scanner := bufio.NewScanner(&out)
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), task) {
				outList = scanner.Text()
				break
			}
		}
		if outList == "" {
			t.Errorf("Task %q is not in the list", task)
		}
		taskCompleteStatus := strings.Fields(outList)[0]

		if taskCompleteStatus != "-" {
			t.Errorf("Expected status %q, got %q", "-", taskCompleteStatus)
		}

		taskId = strings.Fields(outList)[1]
	})
	vRes := t.Run("ViewTask", func(t *testing.T) {
		var out bytes.Buffer
		if err := viewAction(&out, apiRoot, taskId); err != nil {
			t.Fatalf("Expected no error, got %q", err)
		}
		viewOut := strings.Split(out.String(), "\n")

		if !strings.Contains(viewOut[0], task) {
			t.Errorf("Expected task %q, got %q", task, viewOut[0])
		}
		if !strings.Contains(viewOut[1], today) {
			t.Errorf("Expected creation month/day %q, got %q", today, viewOut[1])
		}
		if !strings.Contains(viewOut[2], "No") {
			t.Errorf("Expected completed status %q, got %q", "No", viewOut[2])
		}
	})
	if !vRes {
		t.Fatalf("View task failed. Stopping integration test.")
	}
	t.Run("CompleteTask", func(t *testing.T) {})
	t.Run("ListCompletedTask", func(t *testing.T) {})
	t.Run("DeleteTask", func(t *testing.T) {})
	t.Run("ListDeletedTask", func(t *testing.T) {})
}
