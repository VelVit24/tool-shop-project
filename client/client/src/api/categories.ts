import type { Category } from '../types/category';
import api from './axios';
export async function getCategories() {
    const response = await api.get<Category[]>(`/categories`);
    return response.data;
}