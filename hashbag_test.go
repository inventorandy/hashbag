package hashbag

import (
	"strings"
	"testing"

	"github.com/inventorandy/hashbag/charset"
)

func TestRandomString_Length(t *testing.T) {
	length := 10
	expectedLength := length

	result, err := RandomString(length)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(result) != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, len(result))
	}
}

func TestRandomString_Charsets(t *testing.T) {
	length := 10
	expectedLength := length

	result, err := RandomString(length, charset.LowercaseAlpha)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(result) != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, len(result))
	}

	if strings.ContainsAny(
		result,
		string(charset.UppercaseAlpha)+
			string(charset.Numeric)+
			string(charset.Special),
	) {
		t.Errorf("Expected only lowercase alpha characters, but got %s", result)
	}

	result, err = RandomString(length, charset.UppercaseAlpha)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(result) != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, len(result))
	}

	if strings.ContainsAny(
		result,
		string(charset.LowercaseAlpha)+
			string(charset.Numeric)+
			string(charset.Special),
	) {
		t.Errorf("Expected only uppercase alpha characters, but got %s", result)
	}

	result, err = RandomString(length, charset.Numeric)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(result) != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, len(result))
	}

	if strings.ContainsAny(
		result,
		string(charset.LowercaseAlpha)+
			string(charset.UppercaseAlpha)+
			string(charset.Special),
	) {
		t.Errorf("Expected only numeric characters, but got %s", result)
	}

	result, err = RandomString(length, charset.Special)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(result) != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, len(result))
	}

	if strings.ContainsAny(
		result,
		string(charset.LowercaseAlpha)+
			string(charset.UppercaseAlpha)+
			string(charset.Numeric),
	) {
		t.Errorf("Expected only special characters, but got %s", result)
	}
}

func TestSHA256HashString_Length(t *testing.T) {
	length := 10
	expectedLength := 64

	randStr, err := RandomString(length)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	result := SHA256HashString(randStr)

	if len(result) != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, len(result))
	}
}

func TestSHA512HashString_Length(t *testing.T) {
	length := 10
	expectedLength := 128

	randStr, err := RandomString(length)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	result := SHA512HashString(randStr)

	if len(result) != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, len(result))
	}
}

func TestHMACSHA256HashString_Deterministic(t *testing.T) {
	key := []byte("super-secret-key")
	value := "test@example.com"

	h1 := HMACSHA256HashString(key, value)
	h2 := HMACSHA256HashString(key, value)

	if h1 != h2 {
		t.Errorf("HMAC should be deterministic but values differ")
	}
}

func TestHMACSHA256HashString_DifferentKeys(t *testing.T) {
	key1 := []byte("super-secret-key")
	key2 := []byte("another-secret-key")
	value := "test@example.com"

	h1 := HMACSHA256HashString(key1, value)
	h2 := HMACSHA256HashString(key2, value)

	if h1 == h2 {
		t.Errorf("HMAC with different keys should produce different hashes")
	}
}

func TestHMACSHA256HashString_DifferentValues(t *testing.T) {
	key := []byte("super-secret-key")
	value1 := "test@example.com"
	value2 := "another@example.com"

	h1 := HMACSHA256HashString(key, value1)
	h2 := HMACSHA256HashString(key, value2)

	if h1 == h2 {
		t.Errorf("HMAC with different values should produce different hashes")
	}
}

func TestHMACSHA512HashString_Deterministic(t *testing.T) {
	key := []byte("super-secret-key")
	value := "test@example.com"

	h1 := HMACSHA512HashString(key, value)
	h2 := HMACSHA512HashString(key, value)

	if h1 != h2 {
		t.Errorf("HMAC should be deterministic but values differ")
	}
}

func TestHMACSHA512HashString_DifferentKeys(t *testing.T) {
	key1 := []byte("super-secret-key")
	key2 := []byte("another-secret-key")
	value := "test@example.com"

	h1 := HMACSHA512HashString(key1, value)
	h2 := HMACSHA512HashString(key2, value)

	if h1 == h2 {
		t.Errorf("HMAC with different keys should produce different hashes")
	}
}

func TestHMACSHA512HashString_DifferentValues(t *testing.T) {
	key := []byte("super-secret-key")
	value1 := "test@example.com"
	value2 := "another@example.com"

	h1 := HMACSHA512HashString(key, value1)
	h2 := HMACSHA512HashString(key, value2)

	if h1 == h2 {
		t.Errorf("HMAC with different values should produce different hashes")
	}
}
