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

    const data = JSON.parse(local);

    axios.defaults.headers.common.Authorization = `Bearer ${data.token}`;

    return data;
}

function writeStateToLocalStorage(state) {
    localStorage.setItem('auth', JSON.stringify(state));
}

export const useAuthStore = defineStore('auth', {
    state: loadStateFromLocalStorage,

    getters: {
        isLoggedIn: (state) => state.token != null,
        getBankAccounts: (state) => state.user?.bank_accounts,
        getUser: (state) => state.user,
    },

    actions: {
        changeLoginState({user, token}) {
            this.user = user;
            this.token = token;

            axios.defaults.headers.common.Authorization = `Bearer ${token}`;

            writeStateToLocalStorage({user, token});
        },

        async refresh() {
            const { data } = await axios.get('/auth');
            this.user = data;
        },

        logout() {
            localStorage.removeItem('auth');
            axios.defaults.headers.common.Authorization = null;

            axios.delete('/logout', { headers: {
                Authorization: `Bearer ${this.token}`
            }});

            this.user = null;
            this.token = null;
        }
    }
});
