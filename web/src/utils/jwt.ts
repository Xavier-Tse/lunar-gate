export interface IJwtInfoType {
  exp: number
  isAdmin: boolean
  iss: string
  role: number[]
  userID: number
  username: string
}

export function parseJwt(token: string) {
  const strings = token.split('.')
  return JSON.parse(decodeURIComponent(escape(window.atob(strings[1].replace(/-/g, '+').replace(/_/g, '/')))))
}
