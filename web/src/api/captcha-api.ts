import { useAxios, type baseResponse } from ".";

export interface captchaGenerateResponse {
  captchaID: string
  captcha: string
}

export function captchaGenerateApi(): Promise<baseResponse<captchaGenerateResponse>> {
  return useAxios.get('/api/captcha/generate')
}
