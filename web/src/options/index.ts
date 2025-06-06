import type { optionsResponse } from "@/api";

export interface optionColorType extends optionsResponse<string> {
  color?: string
}

export const methodOptions: optionColorType[] = [
  { label: 'GET', value: 'GET', color: 'blue' },
  { label: 'POST', value: 'POST', color: 'orange' },
  { label: 'PUT', value: 'PUT', color: 'purple' },
  { label: 'PATCH', value: 'PATCH', color: 'pinkpurple' },
  { label: 'DELETE', value: 'DELETE', color: 'red' },
]
