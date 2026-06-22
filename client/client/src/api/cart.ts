import api from './axios';

export async function getCart() {
  const response = await api.get('/cart');
  return response.data;
}

export async function addCartItem(productId: number, quantity: number) {
  const response = await api.post('/cart', { productId, quantity });
  return response.data;
}

export async function removeCartItem(productId: number) {
  const response = await api.delete(`/cart/${productId}`);
  return response.data;
}

export async function changeCartItemQuantity(productId: number, quantity: number) {
  const response = await api.put(`/cart`, { productId, quantity });
  return response.data;
}

export async function clearCart() {
  const response = await api.delete('/cart');
  return response.data;
}