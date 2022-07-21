import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import TouchEvents from '@/helpers/touch-events'

const pinia = createPinia()

createApp(App).use(router)
  .use(pinia)
  .provide('touchEvents', new TouchEvents(document.querySelector('#app')))
  .mount('#app')
