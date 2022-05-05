<template>
  <div class='my-back'>
    <video ref="myBack" autoplay muted="muted" playsinline="playsinline"></video>
  </div>
</template>

<script>
import { promiseTimeout } from '@/helpers/utils'
import Video from '@/helpers/video'

export default {
  name: 'my-back',
  data () {
    return {
      videoStream: new Video(this.receivedStream, this.receivedMessage)
    }
  },
  methods: {
    receivedStream (event) {
      if (event.track.kind == 'video') {
        this.$refs.myBack.srcObject = event.streams[0]
        this.$emit('receivedMessage')
      }
    },
    receivedMessage () { // (data) {
      this.$emit('receivedMessage')
    }
  },
  async created () {
    // somehow picking up click to get here, so wait a second to add callback
    await promiseTimeout(1000)

    this.videoStream.start(this.receivedStream)
  },
  async unmounted () {
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
