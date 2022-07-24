package camera

import (
	"github.com/jonathongardner/wegyb/disk"

	"gocv.io/x/gocv"
	log "github.com/sirupsen/logrus"
)

type Recorder struct {
	name string
	send chan bool
	img *gocv.Mat
}

func NewRecorder(name string) (*Recorder) {
	img := gocv.NewMat()
	return &Recorder{name: name, send: make(chan bool), img: &img}
}

func (r *Recorder) writePump() error {
	ok := r.waitForImageUpdate()
	if !ok {
		return nil
	}

	writer := disk.VideoWriterFile(r.name, 25, r.img.Cols(), r.img.Rows())
	defer writer.Close()
	// img.Close()

	for {
		ok := r.waitForImageUpdate()
		// closed
		if !ok {
			break
		}
		// log.Info("Saving video")

		err := writer.Write(*r.img)
		// img.Close()
		if err != nil && !disk.IsMissingWriter(err) {
			log.Errorf("Error writing closing: %v", err)
		}
	}
	return nil
}
// Use to copy mat images, cant use normal channel b/c gocv mat is binded to
// c++ mat which uses pointers under the hood so they are still linked
// when using chan it was causing random seg faults
func (r *Recorder) waitForImageUpdate() bool {
  r.send <- false
  _, ok := <- r.send
  return ok
}

func (r *Recorder) updateImageIfWaiting(src *gocv.Mat) {
  select {
  case <- r.send:
    src.CopyTo(r.img)
    r.send <- true
  default:
  }
}

func (r *Recorder) Close() error {
	close(r.send)
	r.img.Close()
	return nil
}
