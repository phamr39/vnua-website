import api from "@/services/customers.service";

const state = () => ({
    customer: {},
    store: {},
});

const getters = {};

const actions = {
    async createCustomer({ commit }, payload) {
        const customer = await api.createCustomer(payload);
        commit("setCustomer", customer);
    },

    async createStore({ commit }, payload) {
        const store = await api.createStore(payload);
        commit("setStore", store);
    },
};

const mutations = {
    setCustomer(state, customer) {
        state.customer = customer;
    },
    setStore(state, store) {
        state.store = store;
    },
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
};