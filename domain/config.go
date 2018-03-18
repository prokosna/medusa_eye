package domain

type Config struct {
	FrameRate uint32
	Endpoint  string
	CameraId  string
	Device    string
}

//type ConfigImpl struct {
//	FrameRate uint32
//	Endpoint  string
//	CameraId  string
//}
//
//func (c ConfigImpl) GetFrameRate() uint32 {
//	return c.FrameRate
//}
//
//func (c ConfigImpl) GetEndpoint() string {
//	return c.Endpoint
//}
//
//func (c ConfigImpl) GetCameraId() string {
//	return c.CameraId
//}
