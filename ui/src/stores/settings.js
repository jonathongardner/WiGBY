import { defineStore } from 'pinia'
import { useStorage } from '@vueuse/core'

export const useSettingStore = defineStore('settings', {
  state: () => {
    return {
      flipCamera: useStorage('flipCamera', false)
    }
  },
  // could also be defined as
  // state: () => ({ count: 0 })
  actions: {
    reset() {
      console.log('Reset!')
      this.flipCamera = false
    },
  },
})
