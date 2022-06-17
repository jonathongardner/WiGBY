<template>
  <div class='my-back'>
    <!-- <img ref="myBack" src="https://lh5.googleusercontent.com/proxy/t4bXQKCU3IHZGoeMjvGA8yls6oMw_5xNoIvrlKEortAnvYjLUo__qLpYl1nJJW2gRQg2DD-P3hiN3kk9D_HA-AV9m0BaKxPGKb2PUtDgiPrsCL5_QFNq0hOqmL72YKYZYEMoO0ioIIRqtzIzWaynXZ4OUXBHMfg29JnxmPH_VivNoqkhSGpoE1m_LeUrmQ6ukhs5aH5IiDIL0LHFLWNYlhg0KXjyduNeO3_TeFJH0_lyDqo=s1920-w1920-h1080-fcrop64=1,00001999fffff3c7-k-no-nd-mv"> -->
    <img ref="myBack">
  </div>
</template>

<script>
import { promiseTimeout } from '@/helpers/utils'

export default {
  name: 'my-back',
  data () {
    return {
      socket: null,
    }
  },
  methods: {
    setImage ({ data }) {
      // console.log(data)
      // if (!this.$refs.myBack) {
      //   this.socket.close()
      // }
      this.$refs.myBack.src = 'data:image/jpeg;base64,' + data;
    },
  },
  async created () {
    // somehow picking up click to get here, so wait a second to add callback
    await promiseTimeout(1000)

    this.socket = new WebSocket(`ws://${location.host}/api/v1/mjpeg`)
    // this.socket.binaryType = "arraybuffer"
    this.socket.onmessage = this.setImage
  },
  async unmounted () {
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
  img {
    width: 100vw;
    height: 100vh;
    object-fit: contain;
  }
}
</style>
