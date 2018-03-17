package domain

type Requester interface {
	Post(endpoint string, image *Image) error
}
