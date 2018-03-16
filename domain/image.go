package domain

import "time"

type Image struct {
	CameraId string
	ImageId uint64
	Width uint32
	Height uint32
	Timestamp time.Time
	Data string
}