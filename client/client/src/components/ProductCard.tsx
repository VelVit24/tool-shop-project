import type { Product } from '../types/product';
interface Props {
  product: Product;
}
export default function ProductCard({ product }: Props) {
  return (
    <div>
      <h2>{product.name}</h2>
      <p>Цена: {product.price}</p>
    </div>
  );
}
