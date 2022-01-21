import axios from "axios";
import { API_DOMAIN } from "@/config";

export default{
    async getCustomers({page}){
        return axios.get(
            `${API_DOMAIN}/api/v1/customers?_page=${page}`,
            {withCredentials: true}
        )
        .then((response) => {
            const customers = response.data.customers.map((customer) => {
                return customer;
            });
            return {
                totalItems: response.data.total,
                customers: customers,
            };
        });
    },

    async getStores({page}){
        return axios.get(
            `${API_DOMAIN}/api/v1/stores?_page=${page}`, 
            {withCredentials: true}
        )
        .then((response) => {
            const stores = response.data.stores.map((store) => {
                return store;
            });
            return {
                totalItems: response.data.total,
                stores: stores,
            };
        });
    },

    async getCustomersCSV() {
        return axios.get(`${API_DOMAIN}/api/v1/customers/export`, {withCredentials: true}).then(() => {
            let url = `${API_DOMAIN}/api/v1/customers/export`
            window.open(url)
        });
    },

    async getStoresCSV() {
        return axios.get(`${API_DOMAIN}/api/v1/stores/export`, {withCredentials: true}).then(() => {
            let url2 = `${API_DOMAIN}/api/v1/stores/export`
            window.open(url2)
        });
    }
}