package infra

type EncoderBase64 struct{}

func NewEncoderBase64() *EncoderBase64 {
	return &EncoderBase64{}
}

func (e *EncoderBase64) Encode(raw []byte) string {
	// TODO: Unimplemented
	return "Temporary encoded value"
}

func (e *EncoderBase64) Decode(encoded string) []byte {
	// TODO: Unimplemented
	return []byte("Temporary decoded value")
}
