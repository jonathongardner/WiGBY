import { defineStore } from 'pinia'
import { useStorage } from '@vueuse/core'

export const useSettingStore = defineStore('settings', {
  state: () => {
    return {
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
  },
})
