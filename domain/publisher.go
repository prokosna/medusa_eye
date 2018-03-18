package domain

type Publisher interface {
	Publish(endpoint string, image *Image) error
}
