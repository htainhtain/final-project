import axios from 'axios'

const baseURL = import.meta.env.VITE_BACKEND_URL || '/api'

export const axiosInstance = axios.create({ baseURL })

export const pingpong = () => axiosInstance.get("/health/ping")

export const sql = () => axiosInstance.get("/health/sql")

export const redis = () => axiosInstance.get("/health/redis")

export const keyvault = () => axiosInstance.get("/health/keyvault")