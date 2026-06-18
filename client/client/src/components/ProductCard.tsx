import type { Product } from '../types/product';
interface Props {
  product: Product;
}
export default function ProductCard({ product }: Props) {
  return (
    <div className="w-55 flex flex-col h-55 rounded-xl border p-3">
      <h2 className="text-lg font-bold line-clamp-2">{product.name}</h2>
      <p>Цена: {product.price}</p>
      <button className="mt-auto rounded p-1 outline-1 hover:bg-gray-200 active:bg-gray-700 not-hover:bg-white">
        В корзину
      </button>
    </div>
  );
}
