<script setup>
import { vMaska } from "maska"
import GuestLayout from "@/layouts/GuestLayout.vue";
</script>

<template>
  <GuestLayout>
    <div class="sm:mx-auto sm:w-full sm:max-w-md">
      <h2 class="mt-6 text-center text-3xl font-bold tracking-tight text-gray-900">Acesse sua conta</h2>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
        <form class="space-y-6" @submit.prevent="submitLogin">
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700">Documento (CPF ou CNPJ)</label>
            <div class="mt-1">
              <input
                  v-maska v-model="form.document" data-maska="[ '###.###.###-##', '##.###.###/####-##' ]"
                  id="document" name="document" placeholder="CPF (xxx.xxx.xxx-xx) ou CNPJ (xx.xxx.xxx/xxxx-xx)"  type="text"  required
                  class="block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 placeholder-gray-400 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
              />
              <p v-if="errors['document.required']" class="mt-1.5 text-xs text-red-500">
                {{ errors['document.required'] }}
              </p>
              <p v-if="errors['document.auth']" class="mt-1.5 text-xs text-red-500">
                {{ errors['document.auth'] }}
              </p>
              <p v-if="errors['document.throttle']" class="mt-1.5 text-xs text-red-500">
                {{ errors['document.throttle'] }}
              </p>
            </div>
          </div>

          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">Senha</label>
            <div class="mt-1">
              <input
                  v-model="form.password" id="password" name="password" type="password"
                  autocomplete="current-password" placeholder="*********" required
                  class="block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 placeholder-gray-400 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"/>
              <p v-if="errors['password.required']" class="mt-1.5 text-xs text-red-500">
                {{ errors['password.required'] }}
              </p>
            </div>
          </div>

          <div>
            <button type="submit"
                    class="flex w-full justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
              Acessar
            </button>
          </div>
        </form>
      </div>
    </div>
  </GuestLayout>
</template>

<script>

import { defineComponent } from "vue";
import {useAuthStore} from "@/store/auth";

const AuthStore = useAuthStore();

export default defineComponent({

  name: "LoginView",

  data() {
    return {
      errors: {},
      form: {
        document: '',
        password: '',
      }
    };
  },

  methods: {
    async submitLogin() {
      try {
        await AuthStore.login(this.form.document, this.form.password);
      } catch (e) {
        if (e.response?.data?.errors && e.response.status === 422) {
          const errors = e.response.data.errors;

          if (Object.keys(errors).length === 0) {
            this.errors = { 'document.auth': e.response.data.message };
            return;
          }

          this.errors = e.response.data.errors;
        }

        if (e.response?.data?.message && e.response.status === 429) {
          this.errors = { 'document.throttle' : e.response.data.message };
        }
      }
    }
  },

  mounted() {
    this.submitLogin();
  }
});
</script>