export interface Product {
    id: number
    name: string
    description: string
    price: number
    stock: number
    image_url: string
    id_category: number
}
export interface ProductResponce {
    limit: number
    page: number
    total: number
    products: Product[]
}