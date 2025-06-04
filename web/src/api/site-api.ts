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
