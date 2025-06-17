import type { MenuType } from "@/components/admin/lunar-menu-variables";
import { useAxios, type baseResponse } from ".";
import type { menuType } from "./menu-api";

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

export function permissionRoleMenuTreeApi(): Promise<baseResponse<menuType[]>> {
  return useAxios.get('/api/permission/tree')
}
