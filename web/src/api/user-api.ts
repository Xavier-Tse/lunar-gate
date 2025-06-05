import { useAxios, type baseResponse } from ".";

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
