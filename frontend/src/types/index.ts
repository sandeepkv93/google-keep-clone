export * from './auth'
export * from './note'

export interface ApiResponse<T = any> {
  data?: T
  error?: string
  message?: string
}

export interface PaginatedResponse<T> {
  data: T[]
  total: number
  page: number
  limit: number
  totalPages: number
}

export interface WebSocketMessage {
  type: string
  user_id?: string
  payload: any
}