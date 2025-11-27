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
