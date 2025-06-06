import axios from "axios";
import { Message } from "@arco-design/web-vue";
import { useStore } from "@/stores";

export const useAxios = axios.create({
  timeout: 6000,
  baseURL: "",
})

useAxios.interceptors.request.use((config) => {
  const store = useStore()
  config.headers.set("Authorization", "Bearer " + store.userInfo.token)
  return config
})

useAxios.interceptors.response.use(
  (res) => {
    if (res.status === 200) {
      return res.data
    }
    return res
  },
  (res) => {
    Message.error(res.message)
  }
)

export interface baseResponse<T> {
  code: number
  message: string
  data?: T
}

export interface listResponse<T> {
  count: number
  list: T[]
}

export interface optionsResponse {
  lable: string
  value: number
}

export interface baseParams {
  page?: number
  limit?: number
  sort?: string
  key?: string
}

export function defaultDeleteApi(url: string, idList: number[]): Promise<baseResponse<string>> {
  return useAxios.delete(url, { data: { idList } })
}
