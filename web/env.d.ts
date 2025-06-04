/// <reference types="vite/client" />

declare module 'nprogress'

declare module 'vue-router' {
  interface RouteMeta {
    title?: string
    isLogin?: boolean
  }
}

export interface ImportMetaEnv {
  VITE_SERVER_URL: string
}
