package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseFloat(s string) (float64, error) {
	// Remove sysmbols
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "$", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, "%", "")

	return strconv.ParseFloat(s, 64)
}

func parseBool(s string) (bool, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	switch s {
	case "true", "t", "yes", "y", "1":
		return true, nil
	case "false", "f", "no", "n", "0":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean value: %s", s)
	}
}
