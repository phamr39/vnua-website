import axios from "axios";
import { API_DOMAIN } from "@/config";

export default{
    async getProducts({ page, order, search, categoryId }){
        let filterCategory = "";
        if (categoryId) filterCategory = "&category_id=" + categoryId;

        return axios
        .get(
            `${API_DOMAIN}/api/v1/products?_page=${page}&_order=${order}&q=${search}${filterCategory}`
        )
        .then((response) => {
            const products = response.data.products.map((product) => {
                return product;
            });
            return {
                totalItems: response.data.total,
                products: products,
            };
        });
    },
    async getCategories(rootId){
        return axios.get(`${API_DOMAIN}/api/v1/categories?root=${rootId}`).then((response) => {
            return response.data;
        });
    },
    async getProductById(productId) {
        return axios.get(`${API_DOMAIN}/api/v1/products/${productId}`).then((response) => {
            return response.data;
        });
    },

    async getProductByName(productName) {
        return axios.get(`${API_DOMAIN}/api/v1/products/name/${productName}`).then((response) => {
            return response.data;
        });
    },

    async getProductByHandle(handle) {
        return axios.get(`${API_DOMAIN}/api/v1/products/handle/${handle}`).then((response) => {
            return response.data;
        });
    },

    async getRandomProducts(){
        return axios.get(`${API_DOMAIN}/api/v1/random-products`).then((response) => {
            return response.data;
        });
    },
    async getNewArrivals(){
        return axios.get(`${API_DOMAIN}/api/v1/new-products`).then((response) => {
            return response.data;
        });
    },
    async getProductsByCategory(category){
        return axios.get(`${API_DOMAIN}/api/v1/products/categories/${category}`).then((response) => {
            return response.data;
        });
    },

}