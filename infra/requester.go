package infra

import (
	"github.com/prokosna/medusa_eye/domain"
)

type PublisherHttp struct {
}

func NewPublisherHttp() *PublisherHttp {
	return &PublisherHttp{}
}

func (p *PublisherHttp) Publish(endpoint string, image *domain.Image) error {
	// TODO: Unimplemented
	return nil
}
