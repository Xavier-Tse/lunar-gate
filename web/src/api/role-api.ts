import { useAxios, type baseParams, type baseResponse, type listResponse, type optionsResponse } from "."

export interface roleType {
  id: number
  createdAt: string
  updatedAt: string
  title: string
  roleUserCount: number
  roleApiCount: number
  roleMenuCount: number
  menuIDList: number[]
  apiIDList: number[]
}

export function roleOptionsApi(): Promise<baseResponse<optionsResponse[]>> {
  return useAxios.get('/api/role/options')
}

export function roleListApi(params?: baseParams): Promise<baseResponse<listResponse<roleType>>> {
  return useAxios.get('/api/role', { params })
}

export interface roleCreateRequest {
  id: number
  title: string
}

export function roleCreateApi(data: roleCreateRequest): Promise<baseResponse<string>> {
  if (data.id === 0) {
    return useAxios.post('/api/role', data)
  }
  return useAxios.put('/api/role', data)
}
