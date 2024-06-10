import axios, { type AxiosError, type AxiosInstance, type AxiosResponse, type InternalAxiosRequestConfig } from 'axios'
import { message } from 'ant-design-vue'
import { useUserStore } from '@/stores/token'
import router from '@/router'

interface ErrorResponse {
  message?: string
  error_description?: string
}

const instance:AxiosInstance = axios.create({
  baseURL:"/apis/v1",
  timeout: 10000,
})

export interface RequestConfig extends InternalAxiosRequestConfig {
  skipAuth?: boolean
}

const loginPath = "/login"

const authRequestInterceptor = async (config: RequestConfig)=>{
  if (config.skipAuth) {
    return config
  }
  const accessToken = await getAccessToken()
  if (!accessToken) {
    await router.push(loginPath)
    throw new axios.Cancel("Token expired, redirecting to login")
  }
  config.headers.Authorization = `Bearer ${accessToken}`
  return config
}

instance.interceptors.request.use(authRequestInterceptor, error => {
  return Promise.reject(error);
})

async function getAccessToken(): Promise<string | null> {
  const userStore = useUserStore()
  if (userStore.accessTokenExpired()) {
    if (userStore.refreshTokenExpired()) {
      return ""
    }
    try {
      const response = await axios.get("/apis/v1/token", {
        params: {refresh_token: userStore.refreshToken}
      })
      const {access_token: newAccessToken, refresh_token: newRefreshToken} = response.data.data
      userStore.setUserTokens(newAccessToken, newRefreshToken)
      return newAccessToken
    } catch {
      return ""
    }
  }
  return userStore.accessToken
}

instance.interceptors.response.use(
  (response: AxiosResponse) => response.data,
  (error: AxiosError) => {
  if (error.response) {
    if (error.response.status == 401) {
      router.push(loginPath)
    }else {
      const responseData = error.response.data as ErrorResponse
      const alertMessage = responseData.message || responseData.error_description || "请求失败"
      message.error(alertMessage)
    }
  }else {
    message.error(error.message)
  }
  return Promise.reject(error)
})


export const http = {
  get<T>(url: string, config?: RequestConfig): Promise<T> {
    return instance.get(url, config)
  },
  post<T>(url: string, data?: any, config?: RequestConfig): Promise<T> {
    return instance.post(url, data, config)
  },
  put<T>(url: string, data?: any, config?: RequestConfig): Promise<T> {
    return instance.put(url, data, config)
  },
  delete<T>(url: string, config?: RequestConfig): Promise<T> {
    return instance.delete(url, config)
  }
}
