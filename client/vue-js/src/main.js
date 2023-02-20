import { createApp } from 'vue';
import {createPinia} from "pinia";
import App from './App.vue';
import router from './router';

import './style.css';

const app = createApp(App)

app.use(router)
app.use(createPinia());

app.mount('#app');

import {useAuthStore} from "@/store/auth";
import axios from "axios";

const auth = useAuthStore();
axios.defaults.baseURL = import.meta.env.VITE_API_BASE_URL;
axios.defaults.headers.common.Accept = 'application/json';
if (auth.token) {
    axios.defaults.headers.common.Authorization = auth.token;
}

