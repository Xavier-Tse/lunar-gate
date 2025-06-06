import { useAxios, type baseParams, type baseResponse, type listResponse } from "."

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
  return useAxios.get('/api/api')
}

export interface apiCreateRequest {
  name: string
  "path": string
  method: string
  group: string
}

export function apiCreateApi(data: apiCreateRequest): Promise<baseResponse<string>> {
  return useAxios.post('/api/api', data)
}
