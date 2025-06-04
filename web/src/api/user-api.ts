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

export function userLogin(data: userLoginRequest): Promise<baseResponse<userLoginResponse>> {
  return useAxios.post("/api/user/login", data)
}
