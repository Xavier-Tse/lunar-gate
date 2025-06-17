import { useAxios, type baseResponse } from ".";

export interface dataUserResponse {
  dateList: string[]
  loginCountList: number[]
  registerCountList: number[]
}

export function dataUserApi(): Promise<baseResponse<dataUserResponse>> {
  return useAxios.get('/api/data/user')
}

export interface dataSumResponse {
  userCount: number
  nowLoginCount: number
  nowRegisterCount: number
}

export function dataSumApi(): Promise<baseResponse<dataSumResponse>> {
  return useAxios.get('/api/data/sum')
}
