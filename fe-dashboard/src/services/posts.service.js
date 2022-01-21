import axios from "axios";
import { API_DOMAIN } from "@/config";

export default{
    async getAllNews({ page, order, search, topicId }){
        return axios.get(`${API_DOMAIN}/api/v1/news?_page=${page}&_order=${order}&q=${search}&topic_id=${topicId}`)
        .then((response) => {
            const news = response.data.news.map((post) => {
                return post;
            });
            return {
                totalItems: response.data.total,
                news: news,
            };
        });
    },  
    async getAllEvents({ page, order, search, topicId }){
        return axios.get(`${API_DOMAIN}/api/v1/events?_page=${page}&_order=${order}&q=${search}&topic_id=${topicId}`)
        .then((response) => {
            const events = response.data.events.map((post) => {
                return post;
            });
            return {
                totalItems: response.data.total,
                events: events,
            };
        });
    },

    async getNewestNews(){
        return axios.get(`${API_DOMAIN}/api/v1/new-posts?post_type=news`)
        .then((response) =>{
            const news = response.data
            return news
        })
    },
    async getRandomNews(){
        return axios.get(`${API_DOMAIN}/api/v1/random-posts?post_type=news`)
        .then((response) =>{
            const news = response.data
            return news
        })
    },
    async getRandomEvents(){
        return axios.get(`${API_DOMAIN}/api/v1/random-posts?post_type=event`)
        .then((response) =>{
            const news = response.data
            return news
        })
    },

    async getPostById(postId) {
        return axios.get(`${API_DOMAIN}/api/v1/posts/${postId}`).then((response) => {
            return response.data;
        });
    },
    async GetTopics() {
        return axios.get(`${API_DOMAIN}/api/v1/topics`).then((response) => {
            return response.data;
        });
    },
    async getVideos({key, page}){
        return axios.get(`${API_DOMAIN}/api/v1/videos?_page=${page}&q=${key}`)
        .then((response) => {
            const videos = response.data.videos.map((video) => {
                return video;
            });
            return {
                totalItems: response.data.total,
                videos: videos,
            };
        });
    },

    async createVideo(video) {
        return axios.post(`${API_DOMAIN}/api/v1/videos`, video , {withCredentials: true})
    },

    async createNews(news) {
        return await axios.post(`${API_DOMAIN}/api/v1/posts`, news, {withCredentials: true});
    },

    async deleteVideo(videoId) {
        return axios.delete(`${API_DOMAIN}/api/v1/videos/${videoId}`, {withCredentials: true})
        .then((response) => {
            return response.data;
        });
    },

    async deletePost(postId) {
        return axios.delete(`${API_DOMAIN}/api/v1/posts/${postId}`, {withCredentials: true})
        .then((response) => {
            return response.data;
        });
    }
}