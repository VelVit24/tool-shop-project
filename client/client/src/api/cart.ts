import api from './axios';

export async function getCart() {
  const response = await api.get('/cart');
  return response.data;
}

export async function addCartItem(id_product: number, amount: number) {
  const response = await api.post('/cart', { id_product, amount });
  return response.data;
}

export async function removeCartItem(id_product: number) {
  const response = await api.delete(`/cart/${id_product}`);
  return response.data;
}

export async function changeCartItemQuantity(id_product: number, amount: number) {
  const response = await api.put(`/cart`, { id_product, amount });
  return response.data;
}

export async function clearCartItems() {
  const response = await api.delete('/cart');
  return response.data;
}