package domain

type Encoder interface {
	encode(raw []byte) string
	decode(encoded string) []byte
}
