package infra

import "encoding/base64"

type EncoderBase64 struct{}

func NewEncoderBase64() *EncoderBase64 {
	return &EncoderBase64{}
}

func (e *EncoderBase64) Encode(raw []byte) string {
	return base64.StdEncoding.EncodeToString(raw)
}

func (e *EncoderBase64) Decode(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}
