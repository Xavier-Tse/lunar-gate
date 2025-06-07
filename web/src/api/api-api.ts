import { useAxios, type baseParams, type baseResponse, type listResponse, type optionsResponse } from "."

export interface apiType {
  id: number
  createdAt: string
  updatedAt: string
  name: string
  path: string
  method: string
  group: string
}

export interface apiListRequest extends baseParams {
  method?: string
  group?: string
}

export function apiListApi(params?: apiListRequest): Promise<baseResponse<listResponse<apiType>>> {
  return useAxios.get('/api/api', { params })
}

export interface apiCreateRequest {
  id: number
  name: string
  "path": string
  method: string
  group: string
}

export function apiCreateApi(data: apiCreateRequest): Promise<baseResponse<string>> {
  if (data.id === 0) {
    return useAxios.post('/api/api', data)
  }
  return useAxios.put('/api/api', data)
}

export function apiGroupOptionsApi(): Promise<baseResponse<optionsResponse[]>> {
  return useAxios.get('/api/api/options')
}

export interface systemRouterType {
  method: string
  path: string
  group: string
  name: string
  enable: boolean
}

export function systemRouterApi(): Promise<baseResponse<systemRouterType[]>> {
  return useAxios.get('/api/api/system')
}

export function apiBatchCreateApi(data: systemRouterType[]): Promise<baseResponse<string>> {
  return useAxios.post('/api/api/create/batch', { list: data })
}
