import { createRouter, createWebHistory } from 'vue-router';
import LoginView from "@/views/LoginView.vue";
import RegisterView from "@/views/RegisterView.vue";
import HomeView from "@/views/HomeView.vue";
import TransactionsView from "@/views/TransactionsView.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'welcome',
      meta: { requires_auth: false, guest_only: true }
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
    {
      path: '/home',
      name: 'home',
      component: HomeView,
      meta: { requires_auth: true, guest_only: false }
    },
    {
      path: '/transactions',
      name: 'transactions',
      component: TransactionsView,
      meta: { requires_auth: true, guest_only: false }
    }
  ]
});

export default router;
