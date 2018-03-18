package domain

import "time"

type Frame struct {
	Id        uint64
	Width     uint32
	Height    uint32
	Data      []byte
	Timestamp time.Time
}

type Recorder interface {
	Initialize() error
	GetFrame() (*Frame, error)
	Close() error
}
