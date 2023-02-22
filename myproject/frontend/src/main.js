import {createApp} from 'vue'
import App from './App.vue'

runtime.EventsOn("cpu_usage", (msg) => alert(msg))
createApp(App).mount('#app')
