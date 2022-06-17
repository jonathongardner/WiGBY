package cameraHub

import (
  "encoding/base64"

  "gocv.io/x/gocv"
  log "github.com/sirupsen/logrus"
)
// TODO need to listen to stream for errors
func (h *Hub) StartCamera(deviceID int) {
  webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
    log.Errorf("Device closed: %v\n", deviceID)
		return
	}
  log.Infoln("Started camera")

  defer webcam.Close()

  img := gocv.NewMat()
  defer img.Close()

  for {
    if ok := webcam.Read(&img); !ok {
      log.Errorf("Device closed: %v\n", deviceID)
      return
    }
    if img.Empty() {
      continue
    }

    buf, _ := gocv.IMEncode(".jpg", img)
    bytes := buf.GetBytes()
    base64String := base64.StdEncoding.EncodeToString(bytes)
    base64Bytes := []byte(base64String) // ~ Was 74029, Is 98708
    // log.Infof("Was %v, Is %v", len(bytes), len(base64Bytes))
    h.lock.Lock()
    for _, client := range h.clients {
      // Select to skip streams which are sleeping to drop frames.
      // This might need more thought.
      select {
      case client.send <- base64Bytes:
      default:
      }
    }
    h.lock.Unlock()
    buf.Close()
  }

  return
}
