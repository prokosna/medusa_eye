package domain

type Encoder interface {
	Encode(raw []byte) string
	Decode(encoded string) []byte
}
