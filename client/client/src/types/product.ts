export interface Product {
    id: number
    name: string
    description: string
    price: number
    stock: number
    image_count: number
    category: string
    slug: string
}
export interface ProductResponce {
    limit: number
    page: number
    total: number
    products: Product[]
}