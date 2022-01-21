import api from "@/services/posts.service";

const state = () => ({
    topics : [],
    posts: [],
    post: {},
    events: [],
    event: {},
    listNews: [],
    news: {},
    newestNews: [],
    randomEvents: [],
    randomNews: [],
    pageIndex: 1,
    order: "desc",
    search: "",
    totalItems: 0,
    topicId : "",
});

const getters = {};

const actions = {
    async getAllEvents(
        { state, commit },
        { pageIndex, order, search, topicId}
    ) {
        if (pageIndex) commit("setPageIndex", pageIndex);
        if (order) commit("setOrder", order);
        if (search !== undefined) commit("setSearch", search);
        if (topicId !== undefined) commit("setTopic", topicId)
        const response = await api.getAllEvents({
            page: state.pageIndex,
            order: state.order,
            search: state.search,
            topicId : state.topicId
        });

        commit("setEvents", response);
    },

    async getAllNews(
        { state, commit },
        { pageIndex, order, search, topicId}
    ) {
        if (pageIndex) commit("setPageIndex", pageIndex);
        if (order) commit("setOrder", order);
        if (search !== undefined) commit("setSearch", search);
        if (topicId !== undefined) commit("setTopic", topicId)
        const response = await api.getAllNews({
            page: state.pageIndex,
            order: state.order,
            search: state.search,
            topicId : state.topicId
        });

        commit("setListNews", response);
    },

    async getNewestNews({ commit }) {
        const response = await api.getNewestNews();
        commit("setNewestNews", response)
    },

    async getRandomNews({ commit }) {
        const response = await api.getRandomNews();
        commit("setRandomNews", response)
    },

    async getRandomEvents({ commit }) {
        const response = await api.getRandomEvents();
        commit("setRandomEvents", response)
    },

    async getPostById({ commit }, postId) {
        const post = await api.getPostById(postId);
        commit("setPost", post);
    },
    async getTopics({ commit }){
        const topics = await api.GetTopics();
        commit("setTopics", topics)
    }
};

const mutations = {

    setEvents(state, response) {
        state.events = response.events;
        state.totalItems = response.totalItems;
    },

    setListNews(state, response) {
        state.listNews = response.news;
        state.totalItems = response.totalItems;
    },

    setNewestNews(state, response) {
        state.newestNews = response;
    },

    setRandomNews(state, response) {
        state.randomNews = response;
    },

    setRandomEvents(state, response) {
        state.randomEvents = response;
    },

    setPost(state, post) {
        state.post = post;
    },

    setPageIndex(state, pageIndex) {
        state.pageIndex = pageIndex;
    },

    setOrder(state, order) {
        state.order = order;
    },

    setSearch(state, search) {
        state.search = search;
    },
    setTopic(state, topic){
        state.topicId = topic
    },
    setTopics(state, topics){
        state.topics = topics
    }
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
};