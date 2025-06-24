import { useAxios, type baseParams, type baseResponse, type listResponse, type optionsResponse } from "."

export interface buttonType {
  id: number
  createdAt: string
  updatedAt: string
  title: string
  name: string
  group: string
}

export interface buttonListRequest extends baseParams {
  group?: string
}

export function buttonListApi(params?: buttonListRequest): Promise<baseResponse<listResponse<buttonType>>> {
  return useAxios.get('/api/button', { params })
}

export interface buttonCreateRequest {
  id: number
  name: string
  title: string
  group: string
}

export function buttonCreateApi(data: buttonCreateRequest): Promise<baseResponse<string>> {
  if (data.id === 0) {
    return useAxios.post('/api/button', data)
  }
  return useAxios.put('/api/button', data)
}

export function buttonGroupOptionsApi(): Promise<baseResponse<optionsResponse[]>> {
  return useAxios.get('/api/button/options')
}

export interface buttonGroupListResponse {
  groupTitle: string
  list: buttonType[]
}

export function buttonGroupListApi(): Promise<baseResponse<buttonGroupListResponse>> {
  return useAxios.get('/api/button/tree')
}
