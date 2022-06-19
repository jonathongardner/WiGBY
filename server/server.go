package server

import (
	"net/http"
	"io/fs"
	"context"

	"github.com/jonathongardner/wegyb/camera"

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
    ipAddress := r.Header.Get("X-Real-Ip")
    if ipAddress == "" {
        ipAddress = r.Header.Get("X-Forwarded-For")
    }
    if ipAddress == "" {
        ipAddress = r.RemoteAddr
    }
		log.Infof("Request to %v from %v by %v\n", r.URL.Path, ipAddress, r.Method)

		ctx := context.WithValue(r.Context(), "ipAddress", ipAddress)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NewServer(host string, ch *camera.Hub, ui fs.FS) (*http.Server) {
	// apiV1 := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	// })
	// http.Handle("/api/v1/offer", middleware(apiV1))

	log.Infof("Listening at %v", host)

	apiV1Mjpeg := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ipAddress := r.Context().Value("ipAddress").(string)
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Error(err)
			return
		}
		ch.NewClient(conn, ipAddress, "")
	})

	serverMux := http.NewServeMux()
	serverMux.Handle("/api/v1/mjpeg", middleware(apiV1Mjpeg))
	serverMux.Handle("/", middleware(http.FileServer(http.FS(ui))))

	return &http.Server{
		Addr:    host,
		Handler: serverMux,
	}
}
