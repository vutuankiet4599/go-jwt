package helper

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHashValue(v string) (string, error) {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(v))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func CompareHashValue(hashedValue, unhashedValue string) bool {
	newUnhasedValue, err := GenerateHashValue(unhashedValue)
	if err != nil {
		return false
	}
	return newUnhasedValue == hashedValue
}
