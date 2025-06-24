import type { menuType } from '@/api/menu-api'
import { permissionRoleButtonListApi, permissionRoleMenuTreeApi } from '@/api/permission-api'
import { siteInfoApi, type siteInfoResponse } from '@/api/site-api'
import { userInfoApi, userLogoutApi } from '@/api/user-api'
import router from '@/router'
import { parseJwt } from '@/utils/jwt'
import { Message } from '@arco-design/web-vue'
import { defineStore } from 'pinia'
import type { Router, RouteRecordRaw } from 'vue-router'

interface IUserStoreType {
  userInfo: {
    token: string
    userID: number
    roleList?: number[]
    nickname: string
    avatar: string
  }
  siteInfo: siteInfoResponse
  roleMenuTree: menuType[]
  buttonList: string[]
  isLoadMenu: boolean
}

export const useStore = defineStore('useStore', {
  state(): IUserStoreType {
    return {
      userInfo: {
        token: '',
        userID: 0,
        roleList: [],
        nickname: '',
        avatar: '',
      },
      siteInfo: {
        "site": {
          title: "tit",
          enTitle: "entit",
          slogan: "slo",
          logo: "lo",
          icp: "1a",
        },
        "project": {
          title: "",
          icon: "",
          path: "",
        },
        "login": {
          captcha: {
            enable: false,
            type: "math",
          }
        }
      },
      roleMenuTree: [],
      buttonList: [],
      isLoadMenu: false,
    }
  },
  actions: {
    async saveUser(token?: string) {
      if (token) {
        this.userInfo.token = token
      }
      const res = await userInfoApi()
      if (res.code) {
        Message.error(res.message)
        localStorage.removeItem('lunar-gate-userInfo')
        this.userInfo = {
          token: '',
          userID: 0,
          roleList: [],
          nickname: '',
          avatar: '',
        }
        router.push({ name: 'login' })
        return
      }

      this.userInfo.userID = res.data!.userID
      this.userInfo.avatar = res.data!.avatar
      this.userInfo.nickname = res.data!.nickname
      this.userInfo.roleList = res.data!.roleList

      localStorage.setItem('lunar-gate-userInfo', JSON.stringify(this.userInfo))
    },
    loadUser() {
      const item = localStorage.getItem('lunar-gate-userInfo')
      if (!item) {
        return
      }
      try {
        this.userInfo = JSON.parse(item)
      } catch (err) {
        Message.error('用户信息读取失败')
        localStorage.removeItem('lunar-gate-userInfo')
      }

      const info = parseJwt(this.userInfo.token)
      const exp = info.exp * 1000
      const now = new Date().getTime()
      if (exp < now) {
        Message.warning('登录已过期，请重新登录')
        localStorage.removeItem('lunar-gate-userInfo')
        this.userInfo = {
          token: '',
          userID: 0,
          roleList: [],
          nickname: '',
          avatar: '',
        }
        router.push({ name: 'login' })
        return
      }

      this.saveUser()
    },
    async logout() {
      const res = await userLogoutApi()
      if (res.code) {
        Message.error(res.message)
      }
      localStorage.removeItem('lunar-gate-userInfo')
      this.userInfo = {
        token: '',
        userID: 0,
        roleList: [],
        nickname: '',
        avatar: '',
      }
      router.push({ name: 'login' })
      Message.success('已退出登录')
    },
    async getSiteInfo() {
      const res = await siteInfoApi()
      if (res.code) {
        Message.error(res.message)
        return
      }
      this.siteInfo = res.data as any
    },
    async getRoleMenuTree() {
      if (this.isLoadMenu) {
        return
      }
      const res = await permissionRoleMenuTreeApi()
      if (res.code) {
        this.isLoadMenu = true
        return
      }
      this.roleMenuTree = res.data as any

      const dynamicRoutes = transformRoutes(res.data as any)
      dynamicRoutes.forEach(route => {
        router.addRoute('admin', route)
      })
      this.isLoadMenu = true
    },
    async getButtonList() {
      const res = await permissionRoleButtonListApi()
      if (res.code) {
        return
      }
      this.buttonList = res.data as any
    },
  },
  getters: {
    isLogin(): boolean {
      return this.userInfo.userID !== 0
    },
    hasButton(): (button: string) => boolean {
      return (button) => {
        const item = this.buttonList.find(value => value === button)
        if (item) {
          return true
        }
        return false
      }
    },
  },
})

const viewModules = import.meta.glob('../views/**/*.vue')

function dynamicImport(dynamicViewsModules: Record<string, () => Promise<unknown>>, component: string) {
  const keys = Object.keys(dynamicViewsModules)
  const matchKeys = keys.filter((key) => {
    return key === component
  })
  const matchKey = matchKeys[0]
  return dynamicViewsModules[matchKey]
}

function transformRoutes(backendRoutes: menuType[]): RouteRecordRaw[] {
  return backendRoutes.map(route => {
    const item = {
      ...route,
      children: route.children ? transformRoutes(route.children) : []
    }
    if (route.component) {
      item.component = dynamicImport(viewModules, route.component) as any
    }
    return item as any
  })
}
