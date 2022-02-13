import { promiseTimeout } from './utils'

export default class Video {
  constructor(trackCallback, carCallback) {
    const pc = new RTCPeerConnection({ sdpSemantics: 'unified-plan' })

    // connect audio / video
    pc.addEventListener('track', trackCallback)
    pc.addTransceiver('video', { direction: 'recvonly' })
    pc.addEventListener('icegatheringstatechange', this.iceStateChange)

    const cc = pc.createDataChannel('car')
    cc.addEventListener('message', (event) => {
      carCallback(JSON.parse(event.data))
    })

    this.pc = pc
    this.cc = cc
  }

  start = async () => {
    // if (!['new', 'complete'].includes(this.pc.iceGatheringState)) {
    //   console.log(this.pc.iceGatheringState)
    //   return
    // }

    const offer = await this.pc.createOffer()
    // after offer is set it will call iceStateChange
    await this.pc.setLocalDescription(offer)
  }
  iceStateChange = () => {
    const { iceGatheringState } = this.pc
    console.log(iceGatheringState)

    if (iceGatheringState === 'complete') {
      this.getRemoteAnswer()
    }
  }
  getRemoteAnswer = async () => {
    const offer = this.pc.localDescription;
    const response = await fetch('/api/offer', {
        body: JSON.stringify({ sdp: offer.sdp, type: offer.type }),
        headers: { 'Content-Type': 'application/json' },
        method: 'POST'
    })
    const answer = await response.json()
    await this.pc.setRemoteDescription(answer)
  }
  stop = async () => {
    await promiseTimeout(500)
    this.pc.close()
  }
}
