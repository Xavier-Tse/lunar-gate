import axios from "axios";
import { Message } from "@arco-design/web-vue";


export const useAxios = axios.create({
    timeout: 6000,
    baseURL: "",
})

useAxios.interceptors.request.use((config) => {
    config.headers.set("token", "xxx")
    return config
})

useAxios.interceptors.response.use((res) => {
    if (res.status === 200){
        return res.data
    }
    return res
}, (res)=>{
    Message.error(res.message)
})
