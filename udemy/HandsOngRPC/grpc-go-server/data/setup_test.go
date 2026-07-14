//go:build unit

package data

import (
	"os"
	"testing"
)

var models Models

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
