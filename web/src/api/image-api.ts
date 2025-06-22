import { useAxios, type baseResponse } from ".";

export function imageUploadApi(file: File): Promise<baseResponse<string>> {
  const form = new FormData()
  form.append('file', file)
  return useAxios.post('/api/image/upload', form)
}
