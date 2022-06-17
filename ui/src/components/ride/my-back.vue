<template>
  <my-back v-show='showVideo' />
</template>

<script>
import { promiseTimeout } from '@/helpers/utils'
import MyBack from '@/components/common/my-back.vue'

export default {
  name: 'my-back-with-toggle',
  components: { MyBack },
  data () {
    return {
      showVideo: true,
      toggleTimeout: null,
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
    receivedStream () {
      this.peak()
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
  },
  async unmounted () {
    this.touchEvents.off('tap')
    this.touchEvents.off('dbltap')
  }
}
</script>

<style scoped lang="scss">
</style>
