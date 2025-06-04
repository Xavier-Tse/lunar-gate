import { defineStore } from 'pinia'

export const useStore = defineStore('useStore', {
  state() {
    return {
      userInfo: {
        token: '',
        userID: '',
        roleList: [],
        nickname: '',
        avatar: '',
      }
    }
  },
  actions: {
    saveUser(token: string) {
      this.userInfo.token = token
    }
  }
})
