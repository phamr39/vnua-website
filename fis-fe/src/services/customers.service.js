import axios from "axios";
import { API_DOMAIN } from "@/config";

export default{
    async createCustomer(customer) {
        return axios.post(`${API_DOMAIN}/api/v1/customers`, customer, {withCredentials: true})
        .then((response) => {
            return response.data;
        })
    },

    async createStore(store) {
        return axios.post(`${API_DOMAIN}/api/v1/stores`, store, {withCredentials: true})
        .then((response) => {
            return response.data;
        })
    }
}