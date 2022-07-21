<template>
  <div class='my-back'>
    <!-- <img ref="myBack" src="https://lh5.googleusercontent.com/proxy/t4bXQKCU3IHZGoeMjvGA8yls6oMw_5xNoIvrlKEortAnvYjLUo__qLpYl1nJJW2gRQg2DD-P3hiN3kk9D_HA-AV9m0BaKxPGKb2PUtDgiPrsCL5_QFNq0hOqmL72YKYZYEMoO0ioIIRqtzIzWaynXZ4OUXBHMfg29JnxmPH_VivNoqkhSGpoE1m_LeUrmQ6ukhs5aH5IiDIL0LHFLWNYlhg0KXjyduNeO3_TeFJH0_lyDqo=s1920-w1920-h1080-fcrop64=1,00001999fffff3c7-k-no-nd-mv"> -->
    <canvas ref="myBack" :class='{ flip: flipCamera }' />
  </div>
</template>

<script>
import { mapState } from 'pinia'
import { useSettingStore } from '@/stores/settings'
import { drawImage } from '@/helpers/image-draw'
import { debounce } from '@/helpers/utils'

export default {
  name: 'my-back',
  data () {
    const image = new Image()
    return {
      image,
      socket: null,
      ctx: null,
    }
  },
  computed: mapState(useSettingStore, ['flipCamera']),
  methods: {
    setImageSrc ({ data }) {
      this.image.src = 'data:image/jpeg;base64,' + data
    },
    imageLoad () {
      this.clearIfStale()
      drawImage(this.ctx, this.image)
    },
    clearImage () {
      this.ctx.clearRect(0, 0, this.ctx.canvas.width, this.ctx.canvas.height)
    },
    clearIfStale: debounce(function() {
      this.clearImage()
      const width = this.ctx.canvas.width / 4
      for(let i = 0; i < 4; i++) {
        this.ctx.beginPath()
        this.ctx.rect(width * i, 0, width, this.ctx.canvas.height)
        this.ctx.fillStyle = (i % 2 === 0 ? 'blue' : 'black')
        this.ctx.fill()
      }
    }),
    setSocket () {
      this.socket = new WebSocket(`ws://${location.host}/api/v1/mjpeg`)
      // this.socket.binaryType = "arraybuffer"
      this.socket.onmessage = this.setImageSrc
      this.socket.onclose = (e) => {
        console.log('Socket is closed. Reconnect will be attempted in 1 second.', e.reason);
        setTimeout(() => this.setSocket(), 1000);
      }

      this.socket.onerror = (err) => {
        console.error('Socket encountered error: ', err.message);
        this.socket.close();
      }
    },
  },
  mounted () {
    this.ctx = this.$refs.myBack.getContext("2d")
    // wait to set onload till ctx is set
    this.image.onload = this.imageLoad
    this.setSocket()
  },
  unmounted () {
    this.socket.onclose = () => {}
    this.socket.close()
  }
}
</script>

<style scoped lang="scss">
.my-back {
  .close {
    color: white;
    position: absolute;
    right: 15px;
    top: 0px;
    font-size: 1.5em;
    font-weight: bold;
  }

  .close:hover,
  .close:focus {
    color: black;
    text-decoration: none;
    cursor: pointer;
  }
  .close {
  }
  canvas {
    width: 100vw;
    height: 100vh;
    object-fit: contain;
  }
  canvas.flip {
    -webkit-transform: scaleX(-1);
    transform: scaleX(-1);
  }
}
</style>
