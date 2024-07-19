package hashbag

import (
	"strings"
	"testing"

	"github.com/inventorandy/hashbag/charset"
)

func TestRandomString_Length(t *testing.T) {
	length := 10
	expectedLength := length

	result := RandomString(length)

	if len(result) != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, len(result))
	}
}

func TestRandomString_Charsets(t *testing.T) {
	length := 10
	expectedLength := length

	result := RandomString(length, charset.LowercaseAlpha)

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

	result = RandomString(length, charset.UppercaseAlpha)

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

	result = RandomString(length, charset.Numeric)

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

	result = RandomString(length, charset.Special)

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

	result := SHA256HashString(RandomString(length))

	if len(result) != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, len(result))
	}
}

func TestSHA512HashString_Length(t *testing.T) {
	length := 10
	expectedLength := 128

	result := SHA512HashString(RandomString(length))

	if len(result) != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, len(result))
	}
}
