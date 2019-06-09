package chatsess

import (
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func NewPassword(p string) string {

}

func CheckPassword(p, h string) bool {

}

func password(p string, s []byte) string {
	key, _ := scrypt.Key([]byte(p), s, 32768, 8, 1, 32)
	return fmt.Sprintf("%x_%x", s, key)
}
