package camera

import (
	// "fmt"

	"github.com/pion/mediadevices"
	"github.com/pion/mediadevices/pkg/frame"
	"github.com/pion/mediadevices/pkg/prop"
	"github.com/pion/webrtc/v3"

	// If you don't like x264, you can also use vpx by importing as below
	// "github.com/pion/mediadevices/pkg/codec/vpx" // This is required to use VP8/VP9 video encoder
	// or you can also use openh264 for alternative h264 implementation
	"github.com/pion/mediadevices/pkg/codec/openh264"
	// or if you use a raspberry pi like, you can use mmal for using its hardware encoder
	// "github.com/pion/mediadevices/pkg/codec/mmal"
	// "github.com/pion/mediadevices/pkg/codec/x264" // This is required to use h264 video encoder

	// Note: If you don't have a camera or microphone or your adapters are not supported,
	//       you can always swap your adapters with our dummy adapters below.
	// _ "github.com/pion/mediadevices/pkg/driver/videotest"
	// _ "github.com/pion/mediadevices/pkg/driver/audiotest"
	_ "github.com/pion/mediadevices/pkg/driver/camera"     // This is required to register camera adapter
	// _ "github.com/pion/mediadevices/pkg/driver/microphone" // This is required to register microphone adapter
)

type camera struct {
	Track  mediadevices.MediaStream
	Engine webrtc.MediaEngine
}

func NewCamera() (*camera, error) {
	// Create a new RTCPeerConnection
	// params, err := vpx.NewParams()
	params, err := openh264.NewParams()
	// params, err := mmal.NewParams()
	// params, err := x264.NewParams()
	if err != nil {
		return nil, err
	}
	params.BitRate = 500_000 // 500kbps

	codecSelector := mediadevices.NewCodecSelector(
		mediadevices.WithVideoEncoders(&params),
	)

	Engine := webrtc.MediaEngine{}
	codecSelector.Populate(&Engine)

	Track, err := mediadevices.GetUserMedia(mediadevices.MediaStreamConstraints{
		Video: func(c *mediadevices.MediaTrackConstraints) {
			c.FrameFormat = prop.FrameFormat(frame.FormatI420)
			c.Width = prop.Int(640)
			c.Height = prop.Int(480)
		},
		Codec: codecSelector,
	})
	if err != nil {
		return nil, err
	}

  cam := camera{
    Engine: Engine,
    Track: Track,
  }

  return &cam, nil
}
