import type { Product } from '../types/product';
import { useCart } from '../context/CartContext';
import { Minus, Plus } from 'lucide-react';
import { useState } from 'react';
import AmountControl from './AmountControl';

interface Props {
  product: Product;
}

export default function ProductCard({ product }: Props) {
  const { addToCart, changeQuantity, getCartItem } = useCart();

  const cartItem = getCartItem(product.id);
  const [localAmount, setLocalAmount] = useState(cartItem?.amount.toString());

  const inCart = !!cartItem;

  //const imageUrl = `http://localhost:8080/static/images/products/${product.slug}/small/1.webp`;
  const imageUrl = `http://localhost:8080/static/images/products/drill-makita-df333/small/1.webp`;
  return (
    <div
      className="
      flex gap-x-5 h-56 rounded-xl border p-3
      border-gray-200 shadow-sm hover:shadow-lg
      transition-shadow duration-300
    "
    >
      <img
        src={imageUrl}
        alt={product.name}
        className="w-50 h-50 object-cover rounded-md"
      />

      <div className="flex flex-col flex-1">
        <h2>{product.name}</h2>

        <p>{product.description}</p>

        <div className="flex flex-row mt-auto items-end">
          <div>
            <p className="text-2xl font-bold">{product.price} ₽</p>
            <p>В наличии: {product.stock}</p>
          </div>
          <div className="flex flex-row ml-auto gap-5">
            {inCart ? (
              <div className="flex gap-2 ml-auto mr-0">
                <AmountControl
                  value={cartItem.amount}
                  min={1}
                  max={product.stock}
                  onChange={(val) => changeQuantity(product.id, val)}
                />
              </div>
            ) : null}

            <button
              className={`
              h-10 w-40 rounded p-1 ring ring-gray-200
              hover:bg-blue-400 hover:border-gray-400
              hover:ring-blue-1000
              active:bg-gray-400
              inset-shadow-2xs
              transition-colors duration-200
              ${inCart ? 'bg-blue-500 text-white' : 'bg-white text-black'}

            `}
              onClick={() => {
                inCart
                  ? changeQuantity(product.id, cartItem.amount + 1)
                  : (() => {
                      addToCart(product);
                      setLocalAmount('1');
                    })();
              }}
            >
              {inCart ? 'В корзине' : 'Купить'}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
