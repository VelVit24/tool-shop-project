import type { CartItem as CartItemType } from '../context/CartContext';
import { Plus, Minus, Trash2 } from 'lucide-react';
import { useState, useEffect } from 'react';
import AmountControl from './AmountControl';

interface Props {
  cartItem: CartItemType;
  changeQuantity: (id: number, quantity: number) => void;
  removeFromCart: (id: number) => void;
}

export default function CartItem({
  cartItem,
  changeQuantity,
  removeFromCart,
}: Props) {
  const [localAmount, setLocalAmount] = useState(cartItem.amount.toString());
  useEffect(() => {
    setLocalAmount(cartItem.amount.toString());
  }, [cartItem.amount]);
  return (
    <div
      key={cartItem.id_product}
      className="flex items-center justify-between shadow-sm p-4 rounded-2xl border-gray-200 border hover:shadow-lg transition-shadow duration-300"
    >
      <div>
        <h2 className="text-xl font-bold">{cartItem.name}</h2>
        <p>{cartItem.price} р</p>
        <p>В наличии: {cartItem.stock}</p>
      </div>
      <div className="flex items-center gap-1">
        <p className="mx-2 text-xl">{cartItem.price * cartItem.amount} р.</p>
        <AmountControl
          value={cartItem.amount}
          min={1}
          max={cartItem.stock}
          onChange={(val) => changeQuantity(cartItem.id_product, val)}
        />
        <button
          className="rounded-full w-10 h-10 flex items-center justify-center hover:bg-gray-200 active:bg-gray-400"
          onClick={() => removeFromCart(cartItem.id_product)}
        >
          <Trash2 className="text-red-500 mx-2" />
        </button>
      </div>
    </div>
  );
}
