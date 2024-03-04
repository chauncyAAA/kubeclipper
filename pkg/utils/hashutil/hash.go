package hashutil

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func Sha1(s string) string {
	sum := sha1.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(password, encryptionPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptionPassword), []byte(password)); err != nil {
		return false
	}
	return true
}
