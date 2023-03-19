<script setup>
import {vMaska} from "maska"
import GuestLayout from "@/layouts/GuestLayout.vue";
</script>

<template>
  <GuestLayout>
    <div class="sm:mx-auto sm:w-full sm:max-w-md">
      <h2 class="mt-6 text-center text-3xl font-bold tracking-tight text-gray-900">Inicie seu cadastro</h2>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
        <form class="space-y-6" @submit.prevent="submitRegistration">

          <div>
            <label for="email" class="block text-sm font-medium text-gray-700">Nome</label>
            <div class="mt-1">
              <input
                  v-model="form.name" :disabled="loading"
                  id="name" name="document" placeholder="Seu Nome" type="text" required
                  class="block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 placeholder-gray-400 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
              />
              <p v-if="errors['name.required']" class="mt-1.5 text-xs text-red-500">
                {{ errors['name.required'] }}
              </p>
              <p v-if="errors['name.max']" class="mt-1.5 text-xs text-red-500">
                {{ errors['name.max'] }}
              </p>
            </div>
          </div>

          <div>
            <label for="email" class="block text-sm font-medium text-gray-700">Documento (CPF ou CNPJ)</label>
            <div class="mt-1">
              <input
                  v-maska v-model="form.document" data-maska="[ '###.###.###-##', '##.###.###/####-##' ]"
                  id="document" name="document" placeholder="CPF (xxx.xxx.xxx-xx) ou CNPJ (xx.xxx.xxx/xxxx-xx)"
                  :disabled="loading" type="text" required
                  class="block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 placeholder-gray-400 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
              />
              <p
                  v-if="errors['document.required']"
                  v-text="errors['document.required']"
                  class="mt-1.5 text-xs text-red-500"
              />
              <p
                  v-if="errors['document.regex']"
                  v-text="errors['document.regex']"
                  class="mt-1.5 text-xs text-red-500"
              />
            </div>
          </div>

          <hr>

          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">Senha de Acesso</label>
            <div class="mt-1">
              <input
                  v-model="form.password" id="password" name="password" type="password" :disabled="loading"
                  autocomplete="current-password" placeholder="*********" required
                  class="block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 placeholder-gray-400 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
              />
              <p
                  v-if="errors['password.required']"
                  v-text="errors['password.required']"
                  class="mt-1.5 text-xs text-red-500"
              />

              <p
                  v-if="errors['password.password']"
                  v-text="errors['password.password']"
                  class="mt-1.5 text-xs text-red-500"
              />
            </div>
          </div>

          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">Confirmar Senha</label>
            <div class="mt-1">
              <input
                  v-model="form.password_confirmation" id="password-confirmation" name="password_confirmation"
                  type="password" :disabled="loading"
                  autocomplete="current-password" placeholder="*********" required
                  class="block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 placeholder-gray-400 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
              />

              <p
                  v-if="errors['password.confirmed']"
                  v-text="errors['password.confirmed']"
                  class="mt-1.5 text-xs text-red-500"
              />
            </div>
          </div>

          <hr>

          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">Senha de Transações</label>
            <div class="mt-1">
              <input
                  v-model="form.transaction_password" id="transaction-password" name="transaction_password"
                  type="password" :disabled="loading"
                  autocomplete="current-password" placeholder="*********" required
                  class="block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 placeholder-gray-400 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
              />

              <p
                  v-if="errors['transaction_password.required']"
                  v-text="errors['transaction_password.required']"
                  class="mt-1.5 text-xs text-red-500"
              />

              <p
                  v-if="errors['transaction_password.digitsbetween']"
                  v-text="errors['transaction_password.digitsbetween']"
                  class="mt-1.5 text-xs text-red-500"
              />
            </div>
          </div>

          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">Confirmar Senha</label>
            <div class="mt-1">
              <input
                  v-model="form.transaction_password_confirmation" id="transaction-password-confirmation"
                  name="transaction_password_confirmation" type="password" :disabled="loading"
                  autocomplete="current-password" placeholder="*********" required
                  class="block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 placeholder-gray-400 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
              />

              <p
                  v-if="errors['transaction_password_confirmation.confirmed']"
                  v-text="errors['transaction_password_confirmation.confirmed']"
                  class="mt-1.5 text-xs text-red-500"
              />
            </div>
          </div>

          <div>
            <button
                type="submit"
                :disabled="loading"
                class="flex w-full justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 disabled:opacity-75"
            >
              Realizar Cadastro
            </button>
          </div>
        </form>

        <div class="mt-6">
          <div class="relative">
            <div class="absolute inset-0 flex items-center">
              <div class="w-full border-t border-gray-300"/>
            </div>
            <div class="relative flex justify-center text-sm">
              <span class="bg-white px-2 text-gray-500">Já possui uma conta?</span>
            </div>
          </div>

          <div class="mt-6">
            <div>
              <RouterLink to="/login"
                          class="inline-flex w-full justify-center rounded-md border border-gray-300 bg-white py-2 px-4 text-sm font-medium text-gray-500 shadow-sm hover:bg-gray-50">
                Faça Login
              </RouterLink>
            </div>
          </div>
        </div>
      </div>
    </div>
  </GuestLayout>
</template>

<script>

import {defineComponent} from "vue";
import {useAuthStore} from "@/store/auth";
import axios from "axios";

export default defineComponent({

  name: "RegisterView",

  data() {
    return {
      loading: false,

      errors: {},

      form: {
        name: '',
        document: '',
        password: '',
        password_confirmation: '',
        transaction_password: '',
        transaction_password_confirmation: '',
      }
    };
  },

  methods: {
    async submitRegistration() {
      try {
        this.loading = true;
        this.errors = {};

        const { data } = await axios.post('/register', this.form);
        useAuthStore().changeLoginState(data);

      } catch (e) {
        if (e.response?.data?.errors && e.response.status === 422) {
          const errors = e.response.data.errors;

          if (Object.keys(errors).length === 0) {
            this.errors = { 'document.auth': e.response.data.message };
          } else {
            this.errors = e.response.data.errors;
          }

        }

        if (e.response?.data?.message && e.response.status === 429) {
          this.errors = { 'document.throttle' : e.response.data.message };
        }
      } finally {
        this.loading = false;
      }
    }
  },
});
</script>