package infra

import (
	"bytes"
	"encoding/json"
	"net/http"

	"errors"
	"io/ioutil"

	"github.com/prokosna/medusa_eye/domain"
)

type PublisherHttp struct {
	client http.Client
}

func NewPublisherHttp() *PublisherHttp {
	return &PublisherHttp{
		client: http.Client{},
	}
}

func (p *PublisherHttp) Publish(endpoint string, image *domain.Image) error {
	body, err := json.Marshal(image)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		"POST",
		endpoint,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}
	return nil
}
