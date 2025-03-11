package argon

import (
	"bytes"
	"crypto/rand"
	"fmt"
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

func GetHashPass(salt []byte, plainPassword string) []byte {
	if salt == nil {
		salt = GetSalt(saltLength)
	}
	fmt.Println(string(salt))
	hashedPass := argon2.IDKey([]byte(plainPassword), salt, 1, 64*1024, 4, 32)
	return append(salt, hashedPass...)
}

func CheckPass(passHash []byte, plainPassword string) bool {
	salt := make([]byte, saltLength)
	copy(salt, passHash[:saltLength])
	userPassHash := GetHashPass(salt, plainPassword)
	return bytes.Equal(userPassHash, passHash)
}
