package infra

import (
	"time"

	"errors"
	"sync"

	"fmt"

	"github.com/blackjack/webcam"
	"github.com/labstack/gommon/log"
	"github.com/prokosna/medusa_eye/domain"
)

type RecorderWebcam struct {
	config        domain.Config
	camera        *webcam.Webcam
	isInitialized bool
	mutex         *sync.RWMutex
	format        string
	width         uint32
	height        uint32
	imageId       uint64
	quit          chan bool
	wg            sync.WaitGroup
	latestFrame   []byte
}

func NewRecorderWebcam(config domain.Config) *RecorderWebcam {
	return &RecorderWebcam{
		config:        config,
		camera:        nil,
		isInitialized: false,
		mutex:         new(sync.RWMutex),
		imageId:       0,
		quit:          make(chan bool),
		wg:            sync.WaitGroup{},
		latestFrame:   nil,
	}
}

func (r *RecorderWebcam) Initialize() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if r.isInitialized {
		return errors.New("RecorderWebcam has been already initialized")
	}
	cam, err := webcam.Open(r.config.Device)
	if err != nil {
		return err
	}

	// Configuration
	formats := cam.GetSupportedFormats()
	var format webcam.PixelFormat
	for f, name := range formats {
		if name == r.config.FrameFormat {
			format = f
		}
	}
	if format == 0 {
		return fmt.Errorf("%s is unsupported foramt: please use the show command", r.config.FrameFormat)
	}
	f, w, h, err := cam.SetImageFormat(format, r.config.FrameWidth, r.config.FrameHeight)
	if err != nil {
		return err
	}
	r.format = formats[f]
	r.width = w
	r.height = h

	r.isInitialized = true
	err = cam.StartStreaming()
	if err != nil {
		return err
	}
	r.camera = cam
	r.wg.Add(1)
	go func() {
		timeout := uint32(5)
		for {
			select {
			case <-r.quit:
				r.wg.Done()
				return
			default:
				err := r.camera.WaitForFrame(timeout)
				if err != nil {
					log.Warn(err)
					continue
				}
				frame, err := r.camera.ReadFrame()
				if err != nil {
					log.Warn(err)
					continue
				}
				r.mutex.Lock()
				r.latestFrame = frame
				r.mutex.Unlock()
			}
		}
	}()
	return nil
}

func (r *RecorderWebcam) GetFrame() (*domain.Frame, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if r.latestFrame == nil {
		return nil, nil
	}
	defer func() {
		r.imageId += 1
	}()
	return &domain.Frame{
		Id:        r.imageId,
		Width:     r.width,
		Height:    r.height,
		Data:      r.latestFrame,
		Timestamp: time.Now(),
	}, nil
}

func (r *RecorderWebcam) Close() error {
	r.mutex.RLock()
	if !r.isInitialized {
		return errors.New("RecorderWebcam has not been initialized yet")
	}
	r.mutex.RUnlock()
	r.quit <- true
	r.wg.Wait()
	r.mutex.Lock()
	defer r.mutex.Unlock()
	err := r.camera.Close()
	if err != nil {
		return err
	}
	return nil
}
