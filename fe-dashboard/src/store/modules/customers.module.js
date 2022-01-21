import api from "@/services/customers.service";

const state = () => ({
    customers: [],
    stores: [],
    pageIndex: 1,
    totalItems: 0,
});

const getters = {};

const actions = {
    async getCustomers({ state, commit },{ pageIndex }) {
        if (pageIndex) commit("setPageIndex", pageIndex);
        const response = await api.getCustomers({
            page: state.pageIndex,
        });
        commit("setCustomers", response);
    },

    async getStores({ state, commit },{ pageIndex }) {
        if (pageIndex) commit("setPageIndex", pageIndex);
        const response = await api.getStores({
            page: state.pageIndex,
        });
        commit("setStores", response);
    },

    async getCustomersCSV(){
        await api.getCustomersCSV();
    },

    async getStoresCSV(){
        await api.getStoresCSV();
    }
};

const mutations = {
    setCustomers(state, response) {
        state.customers = response.customers;
        state.totalItems = +response.totalItems;
    },
    setStores(state, response) {
        state.stores = response.stores;
        state.totalItems = +response.totalItems;
    },
    setPageIndex(state, pageIndex) {
        state.pageIndex = pageIndex;
    },
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
};