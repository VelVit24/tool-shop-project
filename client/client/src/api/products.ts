import api from "./axios";
import type { Product, ProductResponce } from "../types/product";
export async function apiGetProducts(
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
export async function apiGetProductImage(slug: string, imageNumber: number) {
    const response = await api.get(`/products/${slug}/images/${imageNumber}`, {
        responseType: 'blob',
    });
    return response.data;
}
export async function apiDeleteProduct(productId: number) {
    const responce = await api.delete(`/products/${productId}`);
    return responce.data;
}
export async function apiUpdateProduct(product: Product) {
    const responce = await api.put(`/products/${product.id}`, product);
    return responce.data;
}
export async function apiCreateProduct(product: Product) {
    const responce = await api.post(`/products`, product);
    return responce.data;
}
export async function apiCheckSlug(name: string) {
    const responce = await api.post(`/products/get/slug`, { name });
    return responce.data;
}