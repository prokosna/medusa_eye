package domain

import "time"

type Image struct {
	CameraId  string
	ImageId   uint64
	Timestamp time.Time
	Data      string
}
