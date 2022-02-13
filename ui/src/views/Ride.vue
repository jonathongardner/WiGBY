<template>
  <div class="seperated-flex">
    <stat :value='time' label='Time' />
    <stat :value='distance' label='Distance' units='M' />
    <stat :value='speed' label='Speed' units='MPH' />
    <stat :value='avgSpeed' label='Average Speed' units='MPH' />
    <div id='cover'>
      <my-back />
      <controls :riding='riding' @update:riding='toggleRide' />
    </div>
  </div>
</template>

<script>
import Stat from '@/components/ride/stat.vue'
import Controls from '@/components/ride/controls.vue'
import MyBack from '@/components/ride/my-back.vue'
import Tracker from '@/helpers/tracker'

export default {
  name: 'Home',
  components: {
    Stat, Controls, MyBack
  },
  data () {
    return {
      riding: false,
      tracker: new Tracker(this.updateRideStats),
      time: '00:00:00',
      distance: '00.00',
      speed: '00.00',
      avgSpeed: '00.00',
    }
  },
  methods: {
    updateRideStats({ time, distance, speed, avgSpeed }) {
      this.time = time
      this.distance = distance
      this.speed = speed
      this.avgSpeed = avgSpeed
    },
    toggleRide () {
      if (this.riding) {
        this.riding = false
        this.tracker.stop()
      } else {
        this.riding = true
        this.tracker.start()
      }
    }
  },
  beforeRouteLeave(to, from, next) {
    const answer = !this.riding || window.confirm("Do you want to cancel this ride?")
    if (!answer) {
      this.tracker.stop()
    }
    next(answer)
  }
}
</script>

<style scoped lang="scss">
.seperated-flex {
  min-height: 100%;
}
#cover {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  box-sizing: border-box;
  padding: 10px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
}
</style>
