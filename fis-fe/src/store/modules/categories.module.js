import api from "@/services/products.service";

const state = () => ({
    rootCate: [],
    childCate: [],
    category: {}
});

const getters = {};

const actions = {
    async getCategories({ commit }, rootId) {

        const response = await api.getCategories(rootId);
        if (rootId == ""){
            commit("setRootCate", response);
        }
        else{
            commit("setChildCate", response);
        }
    },
};

const mutations = {
    setRootCate(state, response) {
        state.rootCate = response;
    },

    setChildCate(state, response) {
        state.childCate = response;
    },
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
};