import api from "@/services/posts.service";

const state = () => ({
    videos : [],
    video : {},
    pageIndex: 1,
    totalItems : 0,
    key : "",
});

const getters = {};

const actions = {
    async getVideos({ state, commit }, {pageIndex, key}) {
        if (pageIndex) {commit("setPage", pageIndex);}
        if (key !== undefined) {commit("setSearch", key);}
        const response = await api.getVideos({
            page : state.pageIndex,
            key : state.key,
        });
        commit("setVideos", response);
    },

    async getNewestVideo({ commit }) {
        const video = await api.getNewestVideo();
        commit("setVideo", video)
    }
};

const mutations = {
    setPage(state, pageIndex) {
        state.pageIndex = pageIndex;
    },

    setKey(state, key) {
        state.key = key;
    },
    setVideo(state, video) {
        state.video = video.data;
    },
    setVideos(state, response){
        state.videos = response.videos
        state.totalItems = response.totalItems
    }
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
};