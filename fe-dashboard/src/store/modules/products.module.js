import api from "@/services/products.service";

const state = () => ({
    products: [],
    product: {},
    categories: [],
    category: {},
    pageIndex: 1,
    order: "desc",
    search: "",
    isSale: 1,
    totalItems: 0,
    mainImage:"",
    newArrivals: [],
    randomProducts: [],
    recommentProducts: [],
    firstCate: "",
    productId: 0,
});

const getters = {};

const actions = {
    async getProducts(
        { state, commit },
        { pageIndex, order, search, category, isSale }
    ) {
        if (isSale) commit("setIsSale", isSale)
        if (pageIndex) commit("setPageIndex", pageIndex);
        if (order) commit("setOrder", order);
        if (search !== undefined) commit("setSearch", search);
        if (category) commit("setCategory", category);

        const response = await api.getProducts({
        page: state.pageIndex,
        order: state.order,
        search: state.search,
        categoryId: state.category.id,
        isSale : state.isSale
        });

        commit("setProducts", response);
    },

    async getCategories({ commit }) {
        const categories = await api.getCategories();
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

    async getRandomProducts({ commit }){
        const products = await api.getRandomProducts();
        commit("setRandomProducts", products)
    },

    async getNewArrivals({ commit }){
        const products = await api.getNewArrivals();
        commit("setNewArrivals", products)
    },

    
    async deleteProductById({ commit },productId) {
        await api.deleteProductById(productId);
        commit("setProductID", productId)
    }
};

const mutations = {
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
    setIsSale(state, isSale){
        state.isSale = isSale
    },

    setCategories(state, categories) {
        state.categories = categories;
    },

    setCategory(state, category) {
        state.category = category;
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
    }, 
    setProductID(state, productId){
        state.productId = productId;
    }
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
};