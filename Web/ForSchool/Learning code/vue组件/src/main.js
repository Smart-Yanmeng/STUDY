import {createApp} from 'vue'
import App from './App.vue'

import './assets/main.css'
import MyComponent from "./components/MyComponent.vue";


const app = createApp(App)
app.component('MyComponent', MyComponent)
app.mount('#app')