package wegyb

import (
	"log"
	"net/http"
	"io/fs"
	"encoding/json"

	"github.com/jonathongardner/wegyb/wegyb/camera"
	"github.com/jonathongardner/wegyb/wegyb/peer"

	"github.com/pion/webrtc/v3"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request to %v from %v by %v\n", r.URL.Path, r.RemoteAddr, r.Method)
		next.ServeHTTP(w, r)
	})
}

func ListenAndServe(host string, ui fs.FS) {
	cam, err := camera.NewCamera()
	if err != nil {
		panic(err)
	}

	connection := peer.NewConnection(&cam.Engine)

	offerFunc := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		offer := webrtc.SessionDescription{}
		if err := json.NewDecoder(req.Body).Decode(&offer); err != nil {
			http.Error(w, "Bad session description!", http.StatusBadRequest)
			return
		}

	  session, err := connection.AddPeer(offer, cam.Track)
	  if err != nil {
			log.Println(err)
	    http.Error(w, "Couldnt connect!", http.StatusBadRequest)
	    return
	  }

		response, err := json.Marshal(session)
		if err != nil {
			log.Println(err)
			http.Error(w, "Couldnt create response!", http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(response); err != nil {
			log.Println(err)
			http.Error(w, "Couldnt write response!", http.StatusBadRequest)
		}
	})

	http.Handle("/api/v1/offer", middleware(offerFunc))
	http.Handle("/", middleware(http.FileServer(http.FS(ui))))

	// Block forever
	http.ListenAndServe(host, nil)
}
