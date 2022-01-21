import api from "@/services/products.service";

const state = () => ({
    products: [],
    product: {},
    categories: [],
    category: {},
    pageIndex: 1,
    order: "desc",
    search: "",
    totalItems: 0,
    mainImage:"",
    newArrivals: [],
    randomProducts: [],
    recommentProducts: [],
    firstCate: "",
});

const getters = {};

const actions = {
    async getProducts(
        { state, commit },
        { pageIndex, order, search, category_id }
    ) {
        if (pageIndex) commit("setPageIndex", pageIndex);
        if (order) commit("setOrder", order);
        if (search !== undefined) commit("setSearch", search);
        if (category_id === '') commit("setCategory", undefined);
        if (category_id) commit("setCategory", category_id);

        const response = await api.getProducts({
            page: state.pageIndex,
            order: state.order,
            search: state.search,
            categoryId: state.category.id,
        });

        commit("setProducts", response);
    },

    async getCategories({ commit }, rootId) {
        const categories = await api.getCategories(rootId);
        commit("setCategories", categories);
    },

    async getProductById({ commit }, productId) {
        const product = await api.getProductById(productId);
        commit("setProduct", product);
        commit("setMainImage", product.images[0].image_url)
    },

    async getProductByName({ commit }, productName) {
        const product = await api.getProductByName(productName);
        commit("setProduct", product);
        commit("setMainImage", product.images[0].image_url)
        commit("setFirstCate", product.categories[0].name)
    },

    async getProductByHandle({ commit }, handle) {
        const product = await api.getProductByHandle(handle);
        commit("setProduct", product);
        commit("setMainImage", product.images[0].image_url)
        commit("setFirstCate", product.categories[0].name)
    },

    async getProductsByCategory({ commit }, category) {
        const products = await api.getProductsByCategory(category);
        commit("setProductsByCategory", products);
    },

    async getRandomProducts({ commit }){
        const products = await api.getRandomProducts();
        commit("setRandomProducts", products)
    },

    async getNewArrivals({ commit }){
        const products = await api.getNewArrivals();
        commit("setNewArrivals", products)
    },
};

const mutations = {
    setProductsByCategory(state, products) {
        state.products = products;
    },

    setProducts(state, response) {
        state.products = response.products;
        state.totalItems = +response.totalItems;
    },

    setRandomProducts(state, products){
        state.randomProducts = products;
    },

    setNewArrivals(state, products){
        state.newArrivals = products
    },

    setProduct(state, product) {
        state.product = product;
    },

    setCategories(state, categories) {
        state.categories = categories;
    },

    setCategory(state, category_id) {
        state.category.id = category_id;
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

    setMainImage(state, product){
        state.mainImage = product
    },

    setFirstCate(state, product){
        state.firstCate = product;
    }
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
};