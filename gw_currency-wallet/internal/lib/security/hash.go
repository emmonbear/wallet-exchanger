package security

import (
	"crypto/sha1"
	"fmt"
)

const salt = "1asdqwe2@#!@sadwqe#!,;dgl"

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
