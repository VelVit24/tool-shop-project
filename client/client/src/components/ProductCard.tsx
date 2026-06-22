import type { Product } from '../types/product';
import { useCart } from '../context/CartContext';
import { Minus, Plus } from 'lucide-react';
import { useState } from 'react';

interface Props {
  product: Product;
}

export default function ProductCard({ product }: Props) {
  const { addToCart, changeQuantity, getCartItem } = useCart();

  const cartItem = getCartItem(product.id);
  const [localAmount, setLocalAmount] = useState(cartItem?.amount.toString());

  const inCart = !!cartItem;

  const imageUrl = `http://localhost:8080/static/images/products/${product.slug}/small/1.webp`;

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
        <h2 className="text-lg font-bold line-clamp-2">{product.name}</h2>

        <p>{product.description}</p>

        <div className="flex flex-row mt-auto items-end">
          <div>
            <p className="text-2xl font-bold">{product.price} ₽</p>
            <p>В наличии: {product.stock}</p>
          </div>
          <div className="flex flex-row ml-auto gap-5">
            {inCart ? (
              <div className="flex gap-2 ml-auto mr-0">
                <button
                  onClick={() => {
                    const newAmount = cartItem.amount - 1;
                    setLocalAmount(newAmount.toString());
                    changeQuantity(cartItem.id_product, newAmount);
                  }}
                >
                  <Minus size={24} />
                </button>

                <input
                  value={localAmount}
                  className="
                  w-12 border border-gray-200
                  text-xl text-center rounded-md px-2
                "
                  onChange={(e) => {
                    const value = e.target.value;
                    setLocalAmount(value);
                  }}
                  onBlur={() => {
                    const value = Number(localAmount);

                    if (Number.isNaN(value) || value < 1) {
                      setLocalAmount('1');
                      changeQuantity(cartItem.id_product, 1);
                      return;
                    }

                    changeQuantity(cartItem.id_product, value);
                  }}
                />

                <button
                  onClick={() => {
                    const newAmount = cartItem.amount + 1;
                    setLocalAmount(newAmount.toString());
                    changeQuantity(cartItem.id_product, newAmount);
                  }}
                >
                  <Plus size={24} />
                </button>
              </div>
            ) : null}

            <button
              className="
              h-10 w-40 rounded p-1
              bg-gray-200 hover:bg-gray-400
              active:bg-gray-700
            "
              onClick={() => addToCart(product)}
            >
              {inCart ? 'В корзине' : 'Купить'}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
