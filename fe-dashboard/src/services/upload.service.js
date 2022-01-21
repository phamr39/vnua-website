import axios from "axios";
import { API_DOMAIN } from "@/config";

export default{
    async uploadImage(file){
        return axios.post(`${API_DOMAIN}/api/v1/upload`, file ,
        {
            headers: {
                'Content-Type': 'multipart/form-data'
            },
            withCredentials: true,
        }).then((response) => {
            return response.data.url;
        })
    },
}