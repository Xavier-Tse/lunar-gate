import { useAxios, type baseResponse } from ".";

export interface permissionRoleMenuRequest {
  roleID: number
  menuIDList: number[]
}

export function permissionRoleMenuApi(data: permissionRoleMenuRequest): Promise<baseResponse<string>> {
  return useAxios.put('/api/permission/menu', data)
}
