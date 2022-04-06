package peer

import (
	"log"

	"github.com/pion/mediadevices"
	"github.com/pion/webrtc/v3"
)


type connection struct {
	config webrtc.Configuration
	api    *webrtc.API
}

func NewConnection(engine *webrtc.MediaEngine) (*connection) {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{},
			},
		},
	}

	api := webrtc.NewAPI(webrtc.WithMediaEngine(engine))

	connection := connection{config: config, api: api,}

	return &connection
}

func (con *connection) AddPeer(offer webrtc.SessionDescription, cameraTrack mediadevices.MediaStream) (*webrtc.SessionDescription, error) {
	peerConnection, err := con.api.NewPeerConnection(con.config)
	if err != nil {
		return nil, err
	}

	// Set the handler for ICE connection state
	// This will notify you when the peer has connected/disconnected
	// might need to remove thie once its done
	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		// closed, failed, disconnected
		log.Printf("Connection State has changed %s \n", connectionState.String())
	})

	for _, track := range cameraTrack.GetTracks() {
		track.OnEnded(func(err error) {
			log.Printf("Track (ID: %s) ended with error: %v\n", track.ID(), err)
		})

		_, err = peerConnection.AddTransceiverFromTrack(track,
			webrtc.RtpTransceiverInit{
				Direction: webrtc.RTPTransceiverDirectionSendonly,
			},
		)
		if err != nil {
			return nil, err
		}
	}

	// Set the remote SessionDescription
	err = peerConnection.SetRemoteDescription(offer)
	if err != nil {
		return nil, err
	}

	// Create an answer
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		return nil, err
	}

	// Create channel that is blocked until ICE Gathering is complete
	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	// Sets the LocalDescription, and starts our UDP listeners
	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		return nil, err
	}

	// Block until ICE Gathering is complete, disabling trickle ICE
	// we do this because we only can exchange one signaling message
	// in a production application you should exchange ICE Candidates via OnICECandidate
	<-gatherComplete

	// Output the answer in base64 so we can paste it in browser
	return peerConnection.LocalDescription(), nil
}
