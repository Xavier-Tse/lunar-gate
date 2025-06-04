import { userInfoApi } from '@/api/user-api'
import { Message } from '@arco-design/web-vue'
import { defineStore } from 'pinia'

interface IUserStoreType {
  userInfo: {
    token: string
    userID: number
    roleList?: number[]
    nickname: string
    avatar: string
  }
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
      }
    }
  },
  actions: {
    async saveUser(token: string) {
      this.userInfo.token = token
      const res = await userInfoApi()
      if (res.code) {
        Message.error(res.message)
        return
      }

      this.userInfo.userID = res.data!.userID
      this.userInfo.avatar = res.data!.avatar
      this.userInfo.nickname = res.data!.nickname
      this.userInfo.roleList = res.data!.roleList

      localStorage.setItem('userInfo', JSON.stringify(this.userInfo))
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
    },
  }
})
