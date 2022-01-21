import axios from "axios";
import { API_DOMAIN } from "@/config";

export default{
    async getProducts({ page, order, search, categoryId, isSale }){
        let filterCategory = "";
        if (categoryId) filterCategory = "&category_id=" + categoryId;

        return axios
        .get(
            `${API_DOMAIN}/api/v1/products?_page=${page}&_order=${order}&_sale=${isSale}&q=${search}${filterCategory}`
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
    async getCategories(){
        return axios.get(`${API_DOMAIN}/api/v1/categories`).then((response) => {
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

    async createProduct(product) {
        return await axios.post(`${API_DOMAIN}/api/v1/products`, product, {withCredentials: true});
    },

    async deleteProductById(productId) {
        return axios.delete(`${API_DOMAIN}/api/v1/products/${productId}`, {withCredentials: true})
        .then((response) => {
            return response.data;
        });
    },
    async updateProductById(productId, product) {
        return axios.put(`${API_DOMAIN}/api/v1/products/${productId}`, product, {withCredentials: true})
        .then((response) => {
            return response.data;
        });
    }

}