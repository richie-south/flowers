export const API_URL = 'http://localhost:3333/api'
export const GET_FLOWER_LIST_URL = `${API_URL}/flowers`
export const GET_FLOWER_URL = id => `${API_URL}/flowers/${id}`
export const WATER_FLOWER_URL = id => `${API_URL}/flowers/${id}/water`
