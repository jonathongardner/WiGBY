package disk

import (
	"time"
	"path/filepath"

	"gocv.io/x/gocv"
	log "github.com/sirupsen/logrus"
)

const videoLength = 10 * 60

type Writer struct {
	writer       *gocv.VideoWriter
	previousTime int64
	name         string
	fps          float64
	width        int
	height       int
}

func VideoWriterFile(name string, fps float64, width int, height int) (w *Writer) {
	return &Writer{ name: name, fps: fps, width: width, height: height, previousTime: 0 }
}

func (w *Writer) filename(t time.Time) string {
	toReturn := filepath.Join(w.name, t.Format("2006-01-02 15:04:05") + ".avi")
	log.Infof("Saving to: %v", toReturn)
	return toReturn
}
func (w *Writer) updateIfNeeded() error {
	now := time.Now()
	timeSlice := now.Unix() / videoLength
	if timeSlice != w.previousTime {
		err := w.Close() // close the inner writer first
		if err != nil {
			return err
		}

		w.writer, err = gocv.VideoWriterFile(w.filename(now), "MJPG", w.fps, w.width, w.height, true)
		if err != nil {
			return err
		}

		w.previousTime = timeSlice
	}
	return nil
}

func (w *Writer) Write(img gocv.Mat) error {
	err := w.updateIfNeeded()
	if err != nil {
		return err
	}

	return w.writer.Write(img)
}

func (w *Writer) Close() error {
	if w.writer != nil {
		err := w.writer.Close()
		w.writer = nil
		return err
	}

	return nil
}
