package app

import "github.com/prokosna/medusa_eye/domain"

type Processor interface {
	Process() error
}

type ProcessorImpl struct {
	encoder   *domain.Encoder
	recorder  *domain.Recorder
	requester *domain.Requester
}

func NewProcessorImpl(encoder *domain.Encoder, recorder *domain.Recorder, requester *domain.Requester) *ProcessorImpl {
	return &ProcessorImpl{
		encoder:   encoder,
		recorder:  recorder,
		requester: requester,
	}
}

func (p *ProcessorImpl) Process() error {

	return nil
}
