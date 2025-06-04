import { useAxios, type baseResponse } from "."

export interface emailSendRequest {
  email: string
  captchaID: string
  captchaCode: string
}

export interface emailSendResponse {
  emailID: string
}

export function emailSendApi(data: emailSendRequest): Promise<baseResponse<emailSendResponse>> {
  return useAxios.post('/api/email/send', data)
}
