package domain

import "time"

type Image struct {
	CameraId  string    `json:"camera_id"`
	ImageId   uint64    `json:"image_id"`
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
}
