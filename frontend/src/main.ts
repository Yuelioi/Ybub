import '@wailsio/runtime'

import '@yuelioi/toast/dist/toast.css'

import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

const pinia = createPinia()

import App from './App.vue'

const app = createApp(App)

app.use(pinia)

app.mount('#app')
