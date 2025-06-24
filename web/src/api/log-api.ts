import { useAxios, type baseParams, type baseResponse } from ".";

export function logListApi(params: baseParams): Promise<baseResponse<string[]>> {
  return useAxios.get('/api/log', { params })
}
