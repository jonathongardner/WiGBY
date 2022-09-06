package server

import (
	"net/http"
	"io/fs"
	"context"

	"github.com/jonathongardner/wegyb/app"
	"github.com/jonathongardner/wegyb/camera"

	log "github.com/sirupsen/logrus"
	"github.com/gorilla/websocket"
	"github.com/gorilla/mux"
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

func NewServer(host string, videoLocation string, ch *camera.Hub, ui fs.FS) (*http.Server) {
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

	apiV1Version := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewSuccessJsonResponse(map[string]string{ "version": app.Version }, 200).Write(w)
	})
	apiV1Videos := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		videoFiles(videoLocation).Write(w)
	})
	apiV1Video := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		videoFile(videoLocation, vars["filename"]).Write(w)
	})
	deleteApiV1Video := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		deleteVideoFile(videoLocation, vars["filename"]).Write(w)
	})

	allElse := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewErrorResponse("Unknown route", 404, nil).Write(w)
	})

	serverMux := mux.NewRouter()
	serverMux.Use(middleware)
	// websocket for streaming
	serverMux.Handle("/api/v1/mjpeg", apiV1Mjpeg).Methods("GET")
	// video crud
	serverMux.Handle("/api/v1/recordings", apiV1Videos).Methods("GET")
	serverMux.Handle("/api/v1/recordings/{filename}", apiV1Video).Methods("GET")
	serverMux.Handle("/api/v1/recordings/{filename}", deleteApiV1Video).Methods("DELETE")
	// settings
	serverMux.Handle("/api/v1/version", apiV1Version).Methods("GET")
	// ui
	serverMux.PathPrefix("/").Handler(http.FileServer(http.FS(ui))).Methods("GET")
	// fallback
	serverMux.PathPrefix("/").Handler(allElse)

	return &http.Server{
		Addr:    host,
		Handler: serverMux,
	}
}
