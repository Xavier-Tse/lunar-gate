import type { baseResponse, optionsResponse } from "@/api"

export interface filterGroupType {
  label: string
  column: string
  callback?: (column: string, val: string) => void
  source: () => Promise<baseResponse<optionsResponse[]>> | optionsResponse[]
  options?: optionsResponse[]
}
