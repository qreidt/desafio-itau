import { defineStore } from "pinia";
import axios from "axios";

function loadStateFromLocalStorage() {
    const local = localStorage.getItem('auth');

    if (! local) {
        return {
            user: null,
            token: null
        };
    }

    return JSON.parse(local);
}

function writeStateToLocalStorage(state) {
    localStorage.setItem('auth', JSON.stringify(state));
}

export const useAuthStore = defineStore('auth', {
    state: loadStateFromLocalStorage,

    getters: {
        getBankAccounts: (state) => state.user?.bank_accounts,
    },

    actions: {
        async login(document, password) {
            const { data } = await axios.post('/login', {
                document, password
            });

            this.user = data.user;
            this.token = data.token;

            writeStateToLocalStorage(data);
        },


    }
});
