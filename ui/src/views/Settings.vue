<template>
  <div class="seperated-flex">
    <div class='option'>
      <div>
        Flip Camera
      </div>
      <SwitchBox class="switch u-pull-right" v-model='settings.flipCamera' />
    </div>
    <div class='option'>
      <div>
        Stale stream
      </div>
      <div>
        <input type="number" min="0.1" step="0.1" v-model='settings.staleImage'> s
      </div>
    </div>
    <div class='reset'>
      <button class='button' @click='reset'>Reset</button>
    </div>
  </div>
</template>

<script setup>
import SwitchBox from '@/components/form/switch'
// import { storeToRefs } from 'pinia'
import { useSettingStore } from '@/stores/settings'

const settings = useSettingStore()

const reset = () => {
  if (confirm(`Are you sure you want to reset your settings?`)) {
    settings.reset()
  }
}
</script>

<style scoped lang="scss">
.seperated-flex {
  line-height: 6.5rem;
  width: 100%;
  div {
    min-height: 65px;
  }
}
.option {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.reset {
  display: flex;
  justify-content: center;
  align-items: center;
  button {
    margin: 0px;
  }
}
</style>
