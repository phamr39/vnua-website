import api from "@/services/users.service";

const state = () => ({
    users: [],
    user: {},
    profiles: {},
    pageIndex: 1,
    totalItems: 0,
    status: false,
    isLogined: false,
});

const getters = {};

const actions = {

    async login({ commit }, payload) {
        const user = await api.login(payload);
        commit("setIsLogined", true)
        commit("setUser", user);
    },
    async logout({ commit }) {
        const status = await api.logout();
        commit("setIsLogined", false);
        commit("setStatus", status);
    },

    async getAllUsers({ state, commit }, { pageIndex }) {
        if (pageIndex) commit("setPageIndex", pageIndex);
        const response = await api.getAllUsers({
            page: state.pageIndex,
        });
        commit("setUsers", response);
    },
    async createNewUser({ commit }, payload) {
        const user = await api.createNewUser(payload);
        commit("setUser", user);
    },
    async getProfiles({ commit }) {
        const profiles = await api.getProfiles();
        commit("setProfiles", profiles);
    },
    async updateProfiles({ commit }, payload) {
        const user = await api.updateProfiles(payload);
        commit("setUser", user);
    },
    async updatePassword({ commit }, payload) {
        const status = await api.updatePassword(payload);
        commit("setStatus", status);
    },
};

const mutations = {
    setUsers(state, response) {
        state.users = response.users;
        state.totalItems = +response.totalItems;
    },
    setPageIndex(state, pageIndex) {
        state.pageIndex = pageIndex;
    },
    setUser(state, user) {
        state.user = user;
    },
    setStatus(state, status){
        state.status = status;
    },
    setIsLogined(state) {
        state.isLogined = true;
    },
    setProfiles(state, profiles){
        state.profiles = profiles;
    },
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
};