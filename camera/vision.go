package camera

import (
	"image/color"

	"gocv.io/x/gocv"
	log "github.com/sirupsen/logrus"
)

// color for the rect when car detected
var red = color.RGBA{255, 0, 0, 0}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Vision struct {
	classifier gocv.CascadeClassifier
  loaded     bool
}

func NewVision(xmlFile string) (*Vision) {
  v := Vision{classifier: gocv.NewCascadeClassifier(), loaded: false}
  if xmlFile != "" {
    v.loaded = v.classifier.Load(xmlFile)
    if !v.loaded {
      log.Errorf("Not reading cascade file: %v", xmlFile)
    }
  }

  return &v
}

func (v *Vision) Close() (error) {
  return v.classifier.Close()
}

func (v *Vision) find(img *gocv.Mat) {
  if !v.loaded {
    return
  }
  // detect faces
  rects := v.classifier.DetectMultiScale(*img)

  // draw a rectangle around each car on the original image
  for _, r := range rects {
    gocv.Rectangle(img, r, red, 3)
  }
}
