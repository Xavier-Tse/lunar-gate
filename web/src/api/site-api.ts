import { useAxios, type baseResponse } from "."

export interface siteInfoResponse {
  site: {
    title: string
    enTitle: string
    slogan: string
    logo: string
    icp: string
  }
  project: {
    title: string
    icon: string
    path: string
  }
  login: {
    captcha: {
      enable: boolean
      type: string
    }
  }
}

export function siteInfoApi(): Promise<baseResponse<siteInfoResponse>> {
  return useAxios.get('/api/site/info')
}

export function siteInfoUpdateApi(data: siteInfoResponse): Promise<baseResponse<string>> {
  return useAxios.put('/api/site/info', data)
}

export interface emailInfoType {
  enable: boolean
  domain: string
  port: number
  sendEmail: string
  authCode: string
  sendNickname: string
}

export function emailInfoApi(): Promise<baseResponse<emailInfoType>> {
  return useAxios.get('/api/site/email')
}

export function emailInfoUpdateApi(data: emailInfoType): Promise<baseResponse<string>> {
  return useAxios.put('/api/site/email', data)
}
