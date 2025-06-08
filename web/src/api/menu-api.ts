import { useAxios, type baseParams, type baseResponse, type listResponse } from "."

export interface menuMetaType {
  icon: string
  title: string
}

export interface menuType {
  id: number
  createdAt: string
  updatedAt: string
  enable: boolean
  name: string
  path: string
  component: string
  meta: menuMetaType
  parentMenuID: number
  children: menuType[]
  sort: number
}

export function menuListApi(params?: baseParams): Promise<baseResponse<listResponse<menuType>>> {
  return useAxios.get('/api/menu', { params })
}

export interface menuCreateRequest {
  id: number
  icon: string
  title: string
  enable: boolean
  name: string
  path: string
  component: string
  parentMenuID?: number
  sort: number
}

export function menuCreateApi(data: menuCreateRequest): Promise<baseResponse<string>> {
  return useAxios.post('/api/menu', data)
}
