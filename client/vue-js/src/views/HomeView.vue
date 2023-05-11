<script setup>
import AppLayout from "@/layouts/AppLayout.vue";
import { useAuthStore } from "@/store/auth";
import {computed} from "vue";

useAuthStore().refresh();
const auth = useAuthStore().getUser;

const auth_type = {
    'FISICAL': 'Pessoa Física',
    'LEGAL': 'Pessoa Jurídica'
}[auth?.type];

const currency_formatter = Intl.NumberFormat('pt-BR', {
    style: 'currency',
    currency: 'BRL'
});

const bank_accounts = computed(() => useAuthStore().getBankAccounts)
</script>

<template>
  <AppLayout>
    <div class="grid md:grid-cols-2 space-x-8">
      <div>
        <h2 class="text-base font-semibold leading-7 text-gray-900">Informações da Conta</h2>
        <p class="mt-1 text-sm leading-6 text-gray-500">
          Informações básicas sobre a conta que você possui
        </p>

        <dl class="mt-6 space-y-6 divide-y divide-gray-100 border-t border-gray-200 text-sm leading-6">
          <div class="pt-6 sm:flex">
            <dt class="font-medium text-gray-900 sm:w-64 sm:flex-none sm:pr-6">
              Nome Completo
            </dt>
            <dd class="mt-1 flex justify-between gap-x-6 sm:mt-0 sm:flex-auto text-gray-900">
              {{ auth.name }}
            </dd>
          </div>
          <div class="pt-6 sm:flex">
            <dt class="font-medium text-gray-900 sm:w-64 sm:flex-none sm:pr-6">
              Tipo de Conta
            </dt>
            <dd class="mt-1 flex justify-between gap-x-6 sm:mt-0 sm:flex-auto text-gray-900">
              {{ auth_type }}
            </dd>
          </div>
          <div class="pt-6 sm:flex">
            <dt class="font-medium text-gray-900 sm:w-64 sm:flex-none sm:pr-6">
              Documento
            </dt>
            <dd class="mt-1 flex justify-between gap-x-6 sm:mt-0 sm:flex-auto text-gray-900">
              {{ auth.document }}
            </dd>
          </div>
        </dl>
      </div>

      <div>
        <h2 class="text-base font-semibold leading-7 text-gray-900">Contas Bancárias</h2>
        <p class="mt-1 text-sm leading-6 text-gray-500">
          Listagem de Contas Bancárias Cadastradas
        </p>

        <dl
          v-for="bank_account in bank_accounts" :key="bank_account.id"
          class="mt-6 space-y-6 divide-y divide-gray-100 border-t border-gray-200 text-sm leading-6"
        >
          <div class="pt-6 sm:flex">
            <dt class="font-medium text-gray-900 sm:w-64 sm:flex-none sm:pr-6">
              {{ bank_account.number }}
            </dt>
            <dd class="mt-1 flex justify-between gap-x-6 sm:mt-0 sm:flex-auto text-gray-900">
              {{ currency_formatter.format(bank_account.balance / 100) }}
            </dd>
          </div>
        </dl>
      </div>
    </div>
  </AppLayout>
</template>

<script>

import {defineComponent} from "vue";

export default defineComponent({
    name: 'HomeView',
});

</script>