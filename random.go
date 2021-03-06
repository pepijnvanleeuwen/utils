package utils

import (
	"crypto/rand"
	"encoding/base64"
	"math"
)

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random number generator fails to function correctly, in which case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a base64 encoded securely generated random string.
// It will return an error if the system's secure random number generator fails to function correctly, in which case the caller should not continue.
func GenerateRandomString(n int) (string, error) {
	// Base64 always increases the size of a string, so scale down the initial generated bytes.
	// This decreases the amount of entropy of the generated string, so use with caution.
	l := math.Ceil(float64(n) * 3 / 4)

	println(int(l))
	b, err := GenerateRandomBytes(int(l))

	return base64.URLEncoding.EncodeToString(b), err
}
