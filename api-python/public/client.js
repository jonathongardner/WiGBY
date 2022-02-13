var pc = null;

function negotiate() {
    pc.addTransceiver('video', { direction: 'recvonly' })
    return pc.createOffer().then(function(offer) {
        return pc.setLocalDescription(offer);
    }).then(function() {
        // wait for ICE gathering to complete
        return new Promise(function(resolve) {
            if (pc.iceGatheringState === 'complete') {
                resolve();
            } else {
                function checkState() {
                    if (pc.iceGatheringState === 'complete') {
                        pc.removeEventListener('icegatheringstatechange', checkState);
                        resolve();
                    }
                }
                pc.addEventListener('icegatheringstatechange', checkState);
            }
        });
    }).then(function() {
        var offer = pc.localDescription;
        return fetch('/api/offer', {
            body: JSON.stringify({
                sdp: offer.sdp,
                type: offer.type,
            }),
            headers: {
                'Content-Type': 'application/json'
            },
            method: 'POST'
        });
    }).then(function(response) {
        return response.json();
    }).then(function(answer) {
        return pc.setRemoteDescription(answer);
    }).catch(function(e) {
        alert(e);
    });
}

function start() {
    console.log('Starting!')

    var config = {
        sdpSemantics: 'unified-plan'
    };

    pc = new RTCPeerConnection(config);

    // connect audio / video
    pc.addEventListener('track', function(evt) {
        if (evt.track.kind == 'video') {
          console.log('Here')
          document.getElementById('video').srcObject = evt.streams[0];
        }
    });

    negotiate();
    // toggleFullScreen()
}

function stop() {
    document.getElementById('stop').style.display = 'none';

    // close peer connection
    setTimeout(function() {
        pc.close();
    }, 500);
    // toggleFullScreen()
}


function toggleFullScreen() {
  if (!document.fullscreenElement) {
    const elem = document.getElementById('video')
    if (elem.webkitRequestFullscreen) {
      elem.webkitRequestFullscreen();
    } else {
      elem.requestFullscreen();
    }
  } else {
    if (document.exitFullscreen) {
      document.exitFullscreen();
    }
  }
}
