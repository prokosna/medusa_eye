package main

import (
	"github.com/fgrosse/goldi"
	"github.com/prokosna/medusa_eye/app"
	"github.com/prokosna/medusa_eye/domain"
	"github.com/prokosna/medusa_eye/infra"
	"github.com/satori/go.uuid"
)

func main() {
	// Config
	conf := domain.Config{
		FrameRate: 10,
		Endpoint:  "http://www.example.com/",
		CameraId:  uuid.NewV4().String(),
		Device: "",
	}

	// DI
	registry := goldi.NewTypeRegistry()
	config := map[string]interface{}{}
	container := goldi.NewContainer(registry, config)
	container.RegisterType("Encoder", infra.NewEncoderBase64)
	container.RegisterType("Publisher", infra.NewPublisherHttp)
	container.RegisterType("Recorder", infra.NewRecorderWebcam, conf)
	container.RegisterType("Processor", app.NewProcessor,
		"@Encoder",
		"@Recorder",
		"@Publisher",
		conf)

	// Application start
	processor := container.MustGet("Processor").(*app.Processor)
	processor.Process()
}
