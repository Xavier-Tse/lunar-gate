import { useAxios, type baseResponse, type optionsResponse } from "."

export interface roleType {
  id: number
  title: string
}

export function roleOptionsApi(): Promise<baseResponse<optionsResponse[]>> {
  return useAxios.get('/api/role/options')
}
