import type { CartItem } from '../context/CartContext';
import type { Order } from '../types/order';
import api from './axios';

export async function apiCreateOrderNoAuth(order: Order, cart_items: CartItem[]) {
    return api.post('/orders/noauth', { order, cart_items });
}

export async function apiCreateOrder(order: Order) {
    return api.post('/orders', order);
}
export async function apiGetOrders(page: number, limit: number) {
    const response = await api.get(`/orders`, {
        params: {
            page,
            limit,
        }
    });
    return response.data;
}