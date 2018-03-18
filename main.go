package main

import (
	"fmt"
	"os"

	"errors"

	"github.com/blackjack/webcam"
	"github.com/fgrosse/goldi"
	"github.com/labstack/gommon/log"
	"github.com/prokosna/medusa_eye/app"
	"github.com/prokosna/medusa_eye/domain"
	"github.com/prokosna/medusa_eye/infra"
	"github.com/satori/go.uuid"
	"github.com/urfave/cli"
)

var (
	device   string
	format   string
	width    uint
	height   uint
	endpoint string
	fps      uint
)

func main() {
	ap := cli.NewApp()
	ap.Name = "Medusa Eye"
	ap.Commands = []cli.Command{
		{
			Name:  "execute",
			Usage: "Execute main process to stream video frames to the API server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "device, D",
					Value:       "/dev/video0",
					Usage:       "Path to a video device",
					Destination: &device,
				},
				cli.StringFlag{
					Name:        "format, F",
					Value:       "Motion-JPEG",
					Usage:       "Video format (You can see all available formats list by show command, and you must choose Motion-JPEG.)",
					Destination: &format,
				},
				cli.UintFlag{
					Name:        "width, W",
					Value:       640,
					Usage:       "Width of a video frame (You can see all available sizes list by show command)",
					Destination: &width,
				},
				cli.UintFlag{
					Name:        "height, H",
					Value:       480,
					Usage:       "Height of a video frame (You can see all available sizes list by show command)",
					Destination: &height,
				},
				cli.StringFlag{
					Name:        "url, U",
					Value:       "http://localhost:8080/api/v1/medusa/frames",
					Usage:       "URL of the endpoint",
					Destination: &endpoint,
				},
				cli.UintFlag{
					Name:        "fps, R",
					Value:       1,
					Usage:       "FPS of this recorder",
					Destination: &fps,
				},
			},
			Action: func(c *cli.Context) error {
				// Config
				conf := domain.Config{
					FrameRate:   uint32(fps),
					Endpoint:    endpoint,
					CameraId:    uuid.NewV4().String(),
					Device:      device,
					FrameFormat: format,
					FrameWidth:  uint32(width),
					FrameHeight: uint32(height),
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
				return processor.Process()
			},
		},
		{
			Name:  "show",
			Usage: "Options for formats and sizes of specified device",
			Action: func(c *cli.Context) error {
				device := c.Args().First()
				if device == "" {
					return errors.New("please specify device path (e.g. /dev/video0)")
				}
				cam, err := webcam.Open(device)
				if err != nil {
					return err
				}
				defer cam.Close()
				formatDesc := cam.GetSupportedFormats()
				var formats []webcam.PixelFormat
				for f := range formatDesc {
					formats = append(formats, f)
				}

				fmt.Printf("Available formats and frame sizes of %s\n", device)
				for i, format := range formats {
					fmt.Printf("[%d] %s\n", i+1, formatDesc[format])
					frames := []webcam.FrameSize(cam.GetSupportedFrameSizes(format))
					for _, frame := range frames {
						fmt.Printf("    %s\n", frame.GetString())
					}
				}
				return nil
			},
		},
	}

	err := ap.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
