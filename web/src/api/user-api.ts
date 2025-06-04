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
