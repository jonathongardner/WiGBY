<template>
  <div v-show='showVideo' class='my-back'>
    <video ref="myBack" autoplay muted playsinline="true"></video>
  </div>
</template>

<script>
import { promiseTimeout } from '@/helpers/utils'
import Video from '@/helpers/video'

export default {
  name: 'my-back',
  data () {
    return {
      showVideo: true,
      toggleTimeout: null,
      videoStream: new Video(this.receivedStream, this.receivedMessage)
    }
  },
  inject: ['touchEvents'],
  methods: {
    peak () {
      this.clearToggleTimeout()
      this.showVideo = true
      this.toggleTimeout = setTimeout(() => {
        this.showVideo = false
        this.clearToggleTimeout()
      }, 4000)
    },
    toggle () {
      this.clearToggleTimeout()
      this.showVideo = !this.showVideo
    },
    clearToggleTimeout () {
      if (this.toggleTimeout) {
        clearInterval(this.toggleTimeout)
        this.toggleTimeout = null
      }
    },
    receivedStream (event) {
      if (event.track.kind == 'video') {
        this.$refs.myBack.srcObject = event.streams[0]
        this.peak()
      }
    },
    receivedMessage () { // (data) {
      this.peak()
    }
  },
  async created () {
    // somehow picking up click to get here, so wait a second to add callback
    await promiseTimeout(1000)

    this.touchEvents.on('tap', this.peak)
    this.touchEvents.on('dbltap', this.toggle)
    this.videoStream.start(this.receivedStream)
  },
  async unmounted () {
    this.touchEvents.off('tap')
    this.touchEvents.off('dbltap')
    await this.videoStream.stop()
  }
}
</script>

<style scoped lang="scss">
.my-back {
  video {
    max-width: 100%;
    max-height: 100%;
  }
}
</style>
