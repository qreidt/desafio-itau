import { createApp } from 'vue';
import {createPinia} from "pinia";
import App from './App.vue';

import './style.css';

const app = createApp(App)
app.use(createPinia());

import router from './router';
app.use(router)

app.mount('#app');

import {useAuthStore} from "@/store/auth";
import axios from "axios";

const auth = useAuthStore();
axios.defaults.baseURL = import.meta.env.VITE_API_BASE_URL;
axios.defaults.headers.common.Accept = 'application/json';

if (auth.isLoggedIn) {
    axios.defaults.headers.common.Authorization = `Bearer ${auth.token}`;
}

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.requires_auth)) {
        if (! auth.isLoggedIn) {
            return next({ name: 'login'});
        }
    }

    if (to.matched.some(record => record.meta.guest_only)) {
      if (auth.isLoggedIn) {
        return next({ name: 'home'});
      }
    }

    return next();
});

