<script setup>
import AppLayout from "@/layouts/AppLayout.vue";

const currency_formatter = Intl.NumberFormat('pt-BR', {
    style: 'currency',
    currency: 'BRL'
});

const date_formatter = Intl.DateTimeFormat('pt-BR', {
    year: 'numeric',
    month: 'numeric',
    day: 'numeric',
    hour: 'numeric',
    minute: 'numeric'
});

function formatDateTime(datetime) {
    return date_formatter.format(new Date(datetime));
}

</script>

<template>
  <AppLayout>
    <div class="flex flex-col space-y-5 divide-y">
      <div class="flex flex-row justify-between">
        <div class="">
          <h2 class="text-base font-semibold leading-7 text-gray-900">Transações</h2>
          <p class="mt-1 text-sm leading-6 text-gray-500">
            Extrato de transações realizadas por todas as suas contas bancárias
          </p>
        </div>

        <div class="flex">
          <button type="button" class="text-sm font-semibold leading-6 text-indigo-600 hover:text-indigo-500">
            <span aria-hidden="true">+</span>
            Realizar nova Transação
          </button>
        </div>
      </div>

      <div class="pt-5">
        <table class="w-full text-left">
          <thead class="border-b border-gray-300 text-gray-900">
          <tr>
            <th class="font-semibold text-gray-900">Data, Hora</th>
            <th class="font-semibold text-gray-900">Valor e Tipo</th>
            <th class="font-semibold text-gray-900">Conta Origem</th>
            <th class="font-semibold text-gray-900">Conta Destino</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="transaction in transactions" :key="transaction.id">
            <td class="hidden py-5 pr-6 sm:table-cell text-sm text-gray-900">
              {{ formatDateTime(transaction.created_at) }}
            </td>
            <td class="text-sm font-medium leading-6 text-gray-900">
              <div class="text-sm leading-6 text-gray-900">{{ currency_formatter.format(transaction.value) }}</div>
              <div class="mt-1 text-xs leading-5 text-gray-500">{{ transaction.value > 0 ? 'Crédito' : 'Débito' }}</div>
            </td>
            <td class="hidden py-5 pr-6 sm:table-cell">
              <div class="text-sm leading-6 text-gray-900">{{ transaction.sender_user.name }}</div>
              <div class="mt-1 text-xs leading-5 text-gray-500">{{ transaction.sender_account.number }}</div>
            </td>
            <td class="hidden py-5 pr-6 sm:table-cell">
              <div class="text-sm leading-6 text-gray-900">{{ transaction.receiver_user.name }}</div>
              <div class="mt-1 text-xs leading-5 text-gray-500">{{ transaction.receiver_account.number }}</div>
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </AppLayout>
</template>

<script>

import {defineComponent} from "vue";
import {useAuthStore} from "@/store/auth";
import axios from "axios";

export default defineComponent({
    name: 'TransactionsView',

    data() {
        return {
            auth: useAuthStore().getUser,
            loading_get: true,
            loading_post: false,
            transactions: []
        };
    },

    methods: {
        async getTransactions() {
            try {
                this.loading_get = true;

                const {data} = await axios.get('/transfers');
                this.transactions = data;

            } catch (e) {
                //
            } finally {
                this.loading_get = false;
            }
        }
    },

    beforeMount() {
        this.getTransactions();
    }
});

</script>