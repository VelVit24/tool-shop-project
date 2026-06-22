import type { Product } from '../types/product';
import { useCart } from '../context/CartContext';
interface Props {
  product: Product;
}
export default function ProductCard({ product }: Props) {
  const { addToCart } = useCart();
  return (
    <div className="w-55 flex flex-col h-55 rounded-xl border p-3 border-gray-200 shadow-sm hover:shadow-lg transition-shadow duration-300">
      <h2 className="text-lg font-bold line-clamp-2">{product.name}</h2>
      <div className="flex flex-col mt-auto">
        <p>{product.price} р</p>
        <p>В наличии: {product.stock}</p>
        <button
          className="rounded p-1 outline-1 outline-gray-200 hover:bg-gray-200 active:bg-gray-700 not-hover:bg-white"
          onClick={() => addToCart(product)}
        >
          В корзину
        </button>
      </div>
    </div>
  );
}
