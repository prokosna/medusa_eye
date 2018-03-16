package domain

type Requester interface {
	post(endpoint string, image *Image) error
}
