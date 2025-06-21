package argon

import (
	"bytes"
	"crypto/rand"
	"golang.org/x/crypto/argon2"
	"os"
	"strconv"
)

var saltLength, _ = strconv.Atoi(os.Getenv("SALT_LENGTH"))

func GetSalt(length int) []byte {
	res := make([]byte, length)
	_, err := rand.Read(res)
	if err != nil {
		return nil
	}
	return res
}

func GetHashPass(plainPassword string, salt []byte) []byte {
	if salt == nil {
		salt = GetSalt(saltLength)
	}

	hashedPass := argon2.IDKey([]byte(plainPassword), salt, 1, 64*1024, 4, 32)
	return append(salt, hashedPass...)
}

func CheckPass(plainPassword string, passHash []byte) bool {
	salt := make([]byte, saltLength)
	copy(salt, passHash[:saltLength])
	userPassHash := GetHashPass(plainPassword, salt)
	return bytes.Equal(userPassHash, passHash)
}
