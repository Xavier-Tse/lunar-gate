import type { MenuType } from "@/components/admin/lunar-menu-variables";
import { useAxios, type baseResponse } from ".";

export interface permissionRoleMenuRequest {
  roleID: number
  menuIDList: number[]
}

export function permissionRoleMenuApi(data: permissionRoleMenuRequest): Promise<baseResponse<string>> {
  return useAxios.put('/api/permission/menu', data)
}

export interface permissionRoleApiRequest {
  roleID: number
  apiIDList: number[]
}

export function permissionRoleApiApi(data: permissionRoleApiRequest): Promise<baseResponse<string>> {
  return useAxios.put('/api/permission/api', data)
}

export function permissionRoleMenuTreeApi(): Promise<baseResponse<MenuType[]>> {
  return useAxios.get('/api/permission/tree')
}
