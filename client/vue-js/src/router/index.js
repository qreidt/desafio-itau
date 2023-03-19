import { createRouter, createWebHistory } from 'vue-router';
import LoginView from "@/views/LoginView.vue";
import RegisterView from "@/views/RegisterView.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      meta: { requires_auth: true, guest_only: false }
    },
    {
      path: '/login',
      name: 'login',
      component: () => LoginView,
      meta: { requires_auth: false, guest_only: true }
    },

    {
      path: '/register',
      name: 'register',
      component: RegisterView,
      meta: { requires_auth: false, guest_only: true }
    },
  ]
});

export default router;
