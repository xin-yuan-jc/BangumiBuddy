import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore('user', () => {
  const accessToken= ref('')
  const refreshToken= ref('')
  const accessTokenExpiredAt= ref(0)
  const refreshTokenExpiredAt= ref(0)

  function setUserTokens(newAccessToken: string, newRefreshToken: string) {
    accessToken.value = newAccessToken
    accessTokenExpiredAt.value = getTokenExpired(newAccessToken)
    refreshToken.value = newRefreshToken
    refreshTokenExpiredAt.value = getTokenExpired(newRefreshToken)
  }
  function accessTokenExpired(): boolean {
    return Date.now() >= accessTokenExpiredAt.value
  }
  function refreshTokenExpired(): boolean {
    return Date.now() >= refreshTokenExpiredAt.value
  }
  return { accessToken, refreshToken, accessTokenExpiredAt, refreshTokenExpiredAt, setUserTokens, accessTokenExpired, refreshTokenExpired}
}, {
  persist: true
})

function getTokenExpired(token: string): number {
  try {
    const payload = atob(token.split('.')[1])
    const jwt = JSON.parse(payload) as {
      exp: number
    }
    return jwt.exp* 1000
  } catch {
    return 0
  }
}