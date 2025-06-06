import axios from "axios";
import { Message } from "@arco-design/web-vue";
import { useStore } from "@/stores";
import { ref, type Ref } from "vue";

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
  label: string
  value: string | number
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

export function getOptions(ref: Ref<optionsResponse[]>, func: () => Promise<baseResponse<optionsResponse[]>>) {
  func().then((res) => {
    ref.value = res.data as any
  })
}

export function generateOptions(func: () => Promise<baseResponse<optionsResponse[]>>): Ref<optionsResponse[]> {
  const r = ref<optionsResponse[]>([])
  func().then((res) => {
    r.value = res.data as any
  })
  return r
}

interface cacheType {
  time: number
  res: Promise<baseResponse<optionsResponse[]>>
}
const cacheData: Record<string, cacheType> = {}

export function generateOptionsCache(func: () => Promise<baseResponse<optionsResponse[]>>): Ref<optionsResponse[]> {
  const key = func.toString()
  let val = cacheData[key]
  if (!val) {
    const res = func()
    val = {
      res: res,
      time: new Date().getTime(),
    }
    cacheData[key] = val
  }

  const nowTime = new Date().getTime()
  if (nowTime - val.time > 5000) {
    const res = func()
    val = {
      res: res,
      time: new Date().getTime(),
    }
    cacheData[key] = val
  }

  const r = ref<optionsResponse[]>([])
  val.res.then((res) => {
    r.value = res.data as any
  })
  return r
}
