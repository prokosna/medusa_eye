package domain

type Recorder interface {
	Initialize() error
	GetFrame() ([]byte, error)
	Close() error
}
