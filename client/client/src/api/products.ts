import api from "./axios";
import type { ProductResponce } from "../types/product";
export async function getProducts(
    page: number, 
    limit: number,
    category?: string | null
) {
    const response = await api.get<ProductResponce>("/products", {
        params: {
            page,
            limit,
            category: category || undefined, // фильтр по категории (0 - все категории)
        }
    });
    return response.data;
}
export async function getProductImage(slug: string, imageNumber: number) {
    const response = await api.get(`/products/${slug}/images/${imageNumber}`, {
        responseType: 'blob',
    });
    return response.data;
}