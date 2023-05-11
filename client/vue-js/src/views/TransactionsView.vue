<script setup>
import {vMaska} from "maska"
import AppLayout from "@/layouts/AppLayout.vue";
import Modal from "@/components/Modal.vue";
import {DialogTitle} from "@headlessui/vue";

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
          <button
            @click="modal = true" type="button"
            class="text-sm font-semibold leading-6 text-indigo-600 hover:text-indigo-500 px-3"
          >
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
              <div class="text-sm leading-6 text-gray-900">{{ currency_formatter.format(transaction.value / 100) }}</div>
              <div class="mt-1 text-xs leading-5 text-gray-500">
                {{ transaction.value > 0 ? 'Crédito' : 'Débito' }} ({{ transaction.type }})
              </div>
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

    <Modal :show="modal" @close="resetModal">
      <div class="flex flex-col space-y-5">
        <DialogTitle as="h3" class="text-lg font-semibold leading-6 text-gray-900 content-start">
          {{ form_step === 1 ? 'Nova Transação' : 'Confirmar Senha' }}
        </DialogTitle>
        <div v-if="error_message" class="px-2 py-3 bg-red-300 rounded text-red-900">
          {{ error_message }}
        </div>
        <form class="space-y-10" @submit.prevent="submit">
          <div v-if="form_step === 1" class="grid grid-cols-1 gap-x-6 gap-y-3 sm:grid-cols-6">
            <div class="sm:col-span-3">
              <label for="first-name" class="block text-sm leading-6 text-gray-900">Documento da Conta Destino</label>
              <div class="mt-2">
                <input
                  v-maska v-model="form.receiver_user_document" type="text"
                  data-maska="[ '###.###.###-##', '##.###.###/####-##' ]"
                  name="receiver_user_document" id="receiver_user_document"
                  placeholder="Ex: 000.000.000-00" required
                  class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>

            <div class="sm:col-span-3">
              <label for="last-name" class="block text-sm leading-6 text-gray-900">Número da Conta Destino</label>
              <div class="mt-2">
                <input
                  v-model="form.receiver_account_number" type="text"
                  name="receiver_account_number" id="receiver_account_number"
                  placeholder="Ex: 00000000000-0" required
                  class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>

            <div class="sm:col-span-3">
              <label for="country" class="block text-sm leading-6 text-gray-900">Tipo</label>
              <div class="mt-2">
                <select
                  v-model="form.type" id="type" name="type"
                  class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6"
                >
                  <option>PIX</option>
                  <option>TED</option>
                  <option>DOC</option>
                </select>
              </div>
            </div>

            <div class="sm:col-span-3">
              <label for="email" class="block text-sm leading-6 text-gray-900">Valor</label>
              <div class="mt-2">
                <input
                  v-model.number="form.value" id="value" name="value" type="number" min="0.01" step="0.01"
                  class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>
          </div>
          <div v-if="form_step === 2" class="grid sm:grid-cols-2 gap-5">
            <div class="sm:col-span-2 py-4 flex justify-center">
              <input
                :value="'*'.repeat(password_combinations.length)"
                type="password" disabled readonly class="bg-gray-200 border-gray-50 rounded"
              >
            </div>

            <div
              v-for="combination in numeral_combinations"
              @click="addToPasswordCombinations(combination)"
              class="flex justify-center p-4 bg-blue-600 hover:bg-blue-700 rounded text-white font-bold text-xl cursor-pointer"
            >
              {{ combination[0] }} ou {{ combination[1] }}
            </div>
            <div
              @click="password_combinations = []"
              class="flex justify-center p-4 bg-amber-600 hover:bg-amber-700 rounded text-white font-bold text-xl  cursor-pointer"
            >
              Limpar
            </div>

          </div>
          <div class="flex justify-end space-x-3">
            <button
              type="button" @click="resetModal"
              class="inline-flex justify-center rounded-md bg-gray-400 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-gray-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gray-600"
            >
              Cancelar
            </button>
            <button
              type="submit" :disabled="loading_post || (form_step === 2 && password_combinations.length < 2)"
              class="inline-flex justify-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 disabled:opacity-75"
            >
              {{ form_step === 1 ? 'Continuar' : 'Enviar' }}
            </button>
          </div>
        </form>
      </div>
    </Modal>
  </AppLayout>
</template>

<script>

import {defineComponent} from "vue";
import {useAuthStore} from "@/store/auth";
import axios from "axios";

const default_form_data = {
    receiver_user_document: null,
    receiver_account_number: null,
    value: 0,
    type: 'PIX'
};

export default defineComponent({
    name: 'TransactionsView',

    data() {
        return {
            auth: useAuthStore().getUser,

            modal: true,
            loading_get: true,
            loading_post: false,
            transactions: [],

            form: default_form_data,
            form_step: 1,
            password_combinations: [],

            errors: {},
            error_message: null,

            numeral_combinations: this.getNewNumeralCombinations(),
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
        },

        async submit() {
            if (this.form_step === 1) {
                this.form_step = 2;
                return;
            }

            try {
                this.errors = {};
                this.error_message = null;

                this.form_step = 1;
                this.loading_post = true;

                await this.store();
                await this.getTransactions();
                this.resetModal();
            } catch (e) {
                if (e.response?.data?.errors) {
                    this.errors = e.response.data.errors;
                }

                if (e.response?.data?.message) {
                    this.error_message = e.response.data.message;
                }
            } finally {
                this.numeral_combinations = this.getNewNumeralCombinations();
                this.password_combinations = [];
                this.loading_post = false;
            }

        },

        async store() {
            const form = Object.assign({}, this.form);

            form.sender_account_id = this.auth.bank_accounts[0].id;
            form.value = Math.floor(form.value * 100);

            form.password_combinations = this.password_combinations.reduce((result, combinations) => {
                if (result.length === 0) {
                    return combinations;
                }

                return result.reduce((aggregator, password) => {
                    aggregator.push(password + combinations[0]);
                    aggregator.push(password + combinations[1]);
                    return aggregator;
                }, []);
            }, []);

            await axios.post('/transfers', form);
        },

        addToPasswordCombinations(combinations) {
          this.password_combinations.push(combinations);
        },

        resetModal() {
            this.modal = false;
            this.form_step = 1;
            this.form = default_form_data;
            this.numeral_combinations = this.getNewNumeralCombinations();
        },

        getNewNumeralCombinations() {
            const numerals = Array.from({length: 10}, (_, i) => i.toString())
                .sort(() => 0.5 - Math.random());

            return new Array(5).fill(null).map(() => {
                return numerals.splice(0, 2);
            });
        },
    },

    beforeMount() {
        this.getTransactions();
    }
});

// 49.161.066/0001-52
// 37709416-2

</script>