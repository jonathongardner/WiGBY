package server

import (
	"net/http"
	"io/fs"

	"github.com/jonathongardner/wegyb/cameraHub"

	log "github.com/sirupsen/logrus"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// need to change for dev only
		return true
	},
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Request to %v from %v by %v\n", r.URL.Path, r.RemoteAddr, r.Method)
		next.ServeHTTP(w, r)
	})
}

func ListenAndServe(host string, device int, ui fs.FS) {
	// apiV1 := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	// })
	// http.Handle("/api/v1/offer", middleware(apiV1))
	ch := cameraHub.NewHub()
	go ch.Run()
	go ch.StartCamera(device)

	apiV1Mjpeg := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ipAddress := readUserIP(r)
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Error(err)
			return
		}
		cl := ch.NewClient(conn, ipAddress, "")
		// Allow collection of memory referenced by the caller by doing all work in
	  go cl.WritePump()
	})
	http.Handle("/api/v1/mjpeg", middleware(apiV1Mjpeg))
	http.Handle("/", middleware(http.FileServer(http.FS(ui))))

	// Block forever
	log.Infof("Listening at %v", host)
	http.ListenAndServe(host, nil)
}

func readUserIP(r *http.Request) string {
    IPAddress := r.Header.Get("X-Real-Ip")
    if IPAddress == "" {
        IPAddress = r.Header.Get("X-Forwarded-For")
    }
    if IPAddress == "" {
        IPAddress = r.RemoteAddr
    }
    return IPAddress
}
