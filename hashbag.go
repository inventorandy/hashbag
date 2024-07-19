package hashbag

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"

	"github.com/inventorandy/hashbag/charset"

	"golang.org/x/exp/rand"
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
func RandomString(length int, charset ...charset.Charset) string {
	chars := getCharset(charset)
	b := make([]byte, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
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
