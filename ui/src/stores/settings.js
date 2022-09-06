import { defineStore } from 'pinia'
import { useStorage } from '@vueuse/core'

export const useSettingStore = defineStore('settings', {
  state: () => {
    return {
      version: '',
      flipCamera: useStorage('flipCamera', false),
      staleImage: useStorage('staleImage', 0.5)
    }
  },
  // could also be defined as
  // state: () => ({ count: 0 })
  actions: {
    reset() {
      console.log('Reset!')
      this.flipCamera = false
      this.staleImage = 0.5
    },
    upsertVersion () {
      if (this.version) {
        return
      }

      fetch('/api/v1/version').then(res => res.json()).then(({ version }) => this.version = version)
    }
  },
})
