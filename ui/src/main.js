import { createApp } from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import TouchEvents from '@/helpers/touch-events'

createApp(App).use(router)
  .provide('touchEvents', new TouchEvents(document.querySelector('#app')))
  .mount('#app')
