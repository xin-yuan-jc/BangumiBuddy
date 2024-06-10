import { http } from '@/apis'
import type { AxiosRequestHeaders } from 'axios'
import { useUserStore } from '@/stores/token'

interface LoginResponse {
  access_token: string
  refresh_token: string
}

export const login = (username: string, password: string) => {
  http.post<LoginResponse>('/token', {
    grant_type: 'password',
    username: username,
    password: password
  }, {
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    } as AxiosRequestHeaders,
    skipAuth: true,
  }).then((response:LoginResponse)=> {
    const userStore = useUserStore()
    userStore.setUserTokens(response.access_token, response.refresh_token)
  }).catch((error) => {})
}