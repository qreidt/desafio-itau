<script setup>
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
              <input id="document" name="document" type="text" placeholder="CPF (xxx.xxx.xxx-xx) ou CNPJ (xx.xxx.xxx/xxxx-xx)" required
                     class="block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 placeholder-gray-400 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"/>
            </div>
          </div>

          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">Senha</label>
            <div class="mt-1">
              <input id="password" name="password" type="password" autocomplete="current-password" placeholder="*********" required
                     class="block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 placeholder-gray-400 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"/>
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

export default defineComponent({
  name: "RegisterView",
  data() {
    return {
      form: {
        document: '',
        password: '',
      }
    };
  },

  methods: {
    async submitLogin() {

    }
  }
});
</script>