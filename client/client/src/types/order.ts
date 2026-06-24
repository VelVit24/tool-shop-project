import type { CartItem } from "../context/CartContext"

export type Order = {
    id: number
    status: string
    total: number
    created_at: string
    phone: string
    email: string
    first_name: string
    last_name: string
}
export type OrderFull = {
    order: Order
    user:{
        id: number
        email: string
        phone: string
        first_name: string
        last_name: string
    }
    cart_items: CartItem[]
}
export type OrderResponce = {
    orders: OrderFull[]
    page: number
    limit: number
    total: number
}