package util

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
	"unsafe"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyz123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ" //The runes for the string
	letterIdxBits = 6                                                               // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1                                            // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits                                              // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// NewToken - Creates a new token
func NewToken(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// GenerateSHA256Hash -  generate a sha256 has and return the hash and the salt used
func GenerateSHA256Hash(input string) (*string, *string) {
	h := sha256.New()
	salt := NewToken(7)
	h.Write([]byte(input + salt))
	sha256hash := hex.EncodeToString(h.Sum(nil))
	return &sha256hash, &salt
}

// GeneratePasswordHash - Generate password hash given the salt and the plain password
func GeneratePasswordHash(pass string, salt string) *string {
	h := sha256.New()
	h.Write([]byte(pass + salt))
	sha256hash := hex.EncodeToString(h.Sum(nil))
	return &sha256hash
}
