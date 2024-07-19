package charset

type Charset string

const (
	LowercaseAlpha Charset = "abcdefghijklmnopqrstuvwxyz"
	UppercaseAlpha Charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numeric        Charset = "0123456789"
	Special        Charset = "!@$%^&*()-_{}[]:;,.?~"
)

func (Charset *Charset) String() string {
	return string(*Charset)
}
