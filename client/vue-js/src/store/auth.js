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
        isLoggedIn: (state) => state.token != null,
        getBankAccounts: (state) => state.user?.bank_accounts,
        getUser: (state) => state.user
    },

    actions: {
        changeLoginState({user, token}) {
            this.user = user;
            this.token = token;
            writeStateToLocalStorage({user, token});
        },

        logout() {
            localStorage.removeItem('auth');
            axios.delete('/logout');
        }
    }
});
