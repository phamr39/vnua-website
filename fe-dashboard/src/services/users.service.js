import axios from "axios";
import { API_DOMAIN } from "@/config";

export default{

    async login(user) {
        return await axios.post(`${API_DOMAIN}/api/v1/login`, user, {withCredentials: true})
        .then((response) => {
            return response.data
        });
    },
    async logout(){
        return await axios.post(`${API_DOMAIN}/api/v1/logout`, null ,{withCredentials: true })
        .then((response) => {
            return response.data
        });
    },

    async updateProfiles(user) {
        return axios.put(`${API_DOMAIN}/api/v1/update-info`, user, {withCredentials: true})
        .then((response) => {
            return response.data;
        });
    },

    async updatePassword(payload) {
        return axios.put(`${API_DOMAIN}/api/v1/profiles`, payload, {withCredentials: true})
        .then((response) => {
            return response.data;
        });
    },

    async getProfiles(){
        return axios.get(`${API_DOMAIN}/api/v1/profiles`, {withCredentials: true})
        .then((response) => {
            return response.data;
        });
    },

    async getAllUsers({page}){
        return axios.get(`${API_DOMAIN}/api/v1/users?_page=${page}`, {withCredentials: true})
        .then((response) => {
            const users = response.data.users.map((user) => {
                return user;
            });
            return {
                totalItems: response.data.total,
                users: users,
            };
        });
    },
    async createNewUser(user){
        return axios.post(`${API_DOMAIN}/api/v1/users`, user ,{withCredentials: true})
        .then((response) => {
            return response.data;
        })
    },

    async deleteUser(userId) {
        return axios.delete(`${API_DOMAIN}/api/v1/users/${userId}`, {withCredentials: true})
        .then((response) => {
            return response.data;
        });
    }

}