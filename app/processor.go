package app

import "github.com/prokosna/medusa_eye/domain"

type Processor struct {
	encoder   *domain.Encoder
	recorder  *domain.Recorder
	requester *domain.Requester
	config    *domain.Config
}

func NewProcessor(
	encoder *domain.Encoder,
	recorder *domain.Recorder,
	requester *domain.Requester,
	config *domain.Config) *Processor {
	return &Processor{
		encoder:   encoder,
		recorder:  recorder,
		requester: requester,
		config:    config,
	}
}

func (p Processor) Process() error {
	p.recorder.Initialize()
	return nil
}
