package app

import (
	"sync"
	"time"

	"github.com/prokosna/medusa_eye/domain"
)

type Processor struct {
	encoder   domain.Encoder
	recorder  domain.Recorder
	publisher domain.Publisher
	config    domain.Config
}

func NewProcessor(
	encoder domain.Encoder,
	recorder domain.Recorder,
	publisher domain.Publisher,
	config domain.Config) *Processor {
	return &Processor{
		encoder:   encoder,
		recorder:  recorder,
		publisher: publisher,
		config:    config,
	}
}

func (p *Processor) Process() error {
	err := p.recorder.Initialize()
	if err != nil {
		return err
	}
	defer p.recorder.Close()

	t := time.NewTicker(time.Duration(1.0/float32(p.config.FrameRate)*1000000000) * time.Nanosecond)
	wg := sync.WaitGroup{}
	errCh := make(chan error, 1000)
	for {
		select {
		case <-t.C:
			wg.Add(1)
			go func(errCh chan error, wg *sync.WaitGroup) {
				defer wg.Done()

				frame, err := p.recorder.GetFrame()
				if err != nil {
					errCh <- err
					return
				}
				if frame == nil {
					return
				}

				enc := p.encoder.Encode(frame.Data)

				image := domain.Image{
					CameraId:  p.config.CameraId,
					ImageId:   frame.Id,
					Timestamp: frame.Timestamp,
					Data:      enc,
				}

				// -- Debugging --
				//err = util.WriteFile(image.ImageId, frame.Data)
				//if err != nil {
				//	errCh <- err
				//	return
				//}
				//fmt.Printf("%+v", image)
				// ------------

				err = p.publisher.Publish(p.config.Endpoint, &image)
				if err != nil {
					errCh <- err
					return
				}
			}(errCh, &wg)
		case err := <-errCh:
			wg.Wait()
			return err
		default:
		}
	}
}
