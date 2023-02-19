package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not prime, it is divisible by 2"},
		{"Zero", 0, false, "0 is not prime, by definition!"},
		{"One", 1, false, "1 is not prime, by definition!"},
		{"Negative", -1, false, "Negative numbers are not prime, by definition!"},
		{"Four", 4, false, "4 is not prime, it is divisible by 2"},
	}

	for _, e := range primeTests {
		t.Run(e.name, func(t *testing.T) {
			prime, msg := isPrime(e.testNum)
			if prime != e.expected {
				t.Errorf("Expected %t, got %t", e.expected, prime)
			}
			if msg != e.msg {
				t.Errorf("Expected %s, got %s", e.msg, msg)
			}
		})
	}
}

// Test prompt

func Test_prompt(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()

	_ = w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = oldOut

	expected := "-> "
	if string(out) != expected {
		t.Errorf("Expected %s, got %s", expected, string(out))
	}
}

// Test intro

func Test_intro(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()

	_ = w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = oldOut

	expected := "Enter a number to check if it is prime.\n"

	if !strings.Contains(string(out), expected) {
		t.Errorf("intro test expected %s, got %s", expected, string(out))
	}
}

// Test checkNumbers
func Test_checkNumbers(t *testing.T) {
	checkTests := []struct {
		name     string
		input    string
		expected string
		done     bool
	}{
		{"Quit", "q", "Goodbye!", true},
		{"Not a number", "a", "Please enter a whole number or 'q' to quit", false},
		{"Prime", "7", "7 is a prime number!", false},
		{"Not prime", "8", "8 is not prime, it is divisible by 2", false},
		{"Zero", "0", "0 is not prime, by definition!", false},
		{"One", "1", "1 is not prime, by definition!", false},
		{"Negative", "-1", "Negative numbers are not prime, by definition!", false},
		{"Four", "4", "4 is not prime, it is divisible by 2", false},
	}

	for _, e := range checkTests {
		t.Run(e.name, func(t *testing.T) {
			input := strings.NewReader(e.input)
			reader := bufio.NewScanner(input)
			res, done := checkNumbers(reader)
			if res != e.expected {
				t.Errorf("Expected %s, got %s", e.expected, res)
			}
			if done != e.done {
				t.Errorf("Expected %t, got %t", e.done, done)
			}
		})
	}
}

// Test main
