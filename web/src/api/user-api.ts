import { useAxios, type baseParams, type baseResponse, type listResponse } from ".";
import type { roleType } from "./role-api";

export interface userLoginRequest {
  username: string
  password: string
  captchaID?: string
  captchaCode?: string
}

export interface userLoginResponse {
  token: string
}

export function userLoginApi(data: userLoginRequest): Promise<baseResponse<userLoginResponse>> {
  return useAxios.post("/api/user/login", data)
}

export interface userInfoResponse {
  userID: number
  nickname: string
  avatar: string
  roleList?: number[]
}

export function userInfoApi(): Promise<baseResponse<userInfoResponse>> {
  return useAxios.get("/api/user/info")
}

export interface userRegisterRequest {
  email: string
  emailCode: string
  emailID: string
  password: string
  rePassword: string
  captcha: string
  captchaID: string
}

export function userRegisterApi(data: userRegisterRequest): Promise<baseResponse<string>> {
  return useAxios.post("/api/user/register", data)
}

export function userLogoutApi(): Promise<baseResponse<string>> {
  return useAxios.post('/api/user/logout')
}

export interface userListRequest extends baseParams {
  username?: string
}

export interface userListResponse {
  id: number
  createdAt: string
  updatedAt: string
  username: string
  nickname: string
  avatar: string
  email: string
  isAdmin: boolean
  roleList: roleType[]
}

export function userListApi(params?: userListRequest): Promise<baseResponse<listResponse<userListResponse>>> {
  return useAxios.get('/api/user', { params })
}

export function userRemoveApi(idList: number[]): Promise<baseResponse<string>> {
  return useAxios.delete('/api/user', { data: { idList } })
}
