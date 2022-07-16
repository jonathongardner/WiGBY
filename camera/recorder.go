package camera

import (
	"github.com/jonathongardner/wegyb/disk"

	"gocv.io/x/gocv"
	// log "github.com/sirupsen/logrus"
)

type Recorder struct {
	send chan gocv.Mat
	name string
}

func NewRecorder(name string) (*Recorder) {
	return &Recorder{name: name, send: make(chan gocv.Mat)}
}

func (r *Recorder) writePump() error {
	img, ok := <- r.send
	if !ok {
		return nil
	}

	writer := disk.VideoWriterFile(r.name, 25, img.Cols(), img.Rows())
	defer writer.Close()

	for {
		img, ok := <- r.send
		// closed
		if !ok {
			break
		}
		// log.Info("Saving video")

		err := writer.Write(img)

		if err != nil {
			return err
		}
	}
	return nil
}
