package hashbag

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"

	crand "crypto/rand"

	"github.com/inventorandy/hashbag/charset"
)

func getCharset(cs []charset.Charset) string {
	var s string
	for _, c := range cs {
		s += c.String()
	}
	if s == "" {
		s = string(charset.LowercaseAlpha) + string(charset.UppercaseAlpha) + string(charset.Numeric)
	}
	return s
}

func hashString(h hash.Hash, s ...string) string {
	for _, str := range s {
		h.Write([]byte(str))
	}
	bs := h.Sum(nil)

	// Return the String
	return fmt.Sprintf("%x", bs)
}

// RandomString generates a random string of the given length using the given charset(s).
// If no charset is provided, it will use the default charset (lowercase alpha, uppercase alpha, and numeric).
// The charset(s) can be any combination of the following:
// - charset.LowercaseAlpha
// - charset.UppercaseAlpha
// - charset.Numeric
// - charset.Special
// Example:
//
//	RandomString(10) // "aBcD3eFgH1"
//	RandomString(10, charset.LowercaseAlpha, charset.Numeric) // "a1b2c3d4e5"
//	RandomString(10, charset.LowercaseAlpha, charset.UppercaseAlpha) // "aBcDeFgHiJ"
func RandomString(length int, charset ...charset.Charset) (string, error) {
	chars := getCharset(charset)
	if len(chars) == 0 {
		return "", fmt.Errorf("empty charset")
	}
	b := make([]byte, length)
	buf := make([]byte, length)
	if _, err := crand.Read(buf); err != nil {
		return "", err
	}
	for i := range b {
		b[i] = chars[int(buf[i])%len(chars)]
	}
	return string(b), nil
}

// MD5HashString generates a SHA256 hash of the given string(s).
// and returns the hash as a string.
func SHA256HashString(s ...string) string {
	return hashString(sha256.New(), s...)
}

// SHA512HashString generates a SHA512 hash of the given string(s).
// and returns the hash as a string.
func SHA512HashString(s ...string) string {
	return hashString(sha512.New(), s...)
}

// MD5HashString generates a MD5 hash of the given string(s).
// and returns the hash as a string.
func MD5HashString(s ...string) string {
	return hashString(md5.New(), s...)
}

// HMACSHA256HashString generates a HMAC-SHA256 hash of the
// given string(s) using the given key and returns the hash
// as a string.
func HMACSHA256HashString(key []byte, value string) string {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(value))
	return hex.EncodeToString(mac.Sum(nil))
}

// HMACSHA512HashString generates a HMAC-SHA512 hash of the
// given string(s) using the given key and returns the hash
// as a string.
func HMACSHA512HashString(key []byte, value string) string {
	mac := hmac.New(sha512.New, key)
	mac.Write([]byte(value))
	return hex.EncodeToString(mac.Sum(nil))
}
