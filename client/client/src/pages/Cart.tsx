import { useState } from 'react';
import Container from '../components/Container';
import { useCart } from '../context/CartContext';
import { Plus, Minus, Trash2 } from 'lucide-react';
import CartItem from '../components/CartItem';
export default function Cart() {
  const { cart, removeFromCart, changeQuantity, clearCart } = useCart();
  const total = cart.reduce((acc, item) => acc + item.price * item.amount, 0);

  return (
    <Container>
      <div className="flex flex-col gap-4">
        <div className="flex flex-row justify-between items-center mt-3">
          <h1 className="text-2xl font-bold">Корзина</h1>
          <button
            className="rounded p-1 outline-1 outline-gray-200 hover:bg-gray-200 active:bg-gray-700 not-hover:bg-white"
            onClick={clearCart}
          >
            Очистить корзину
          </button>
        </div>

        {cart.length === 0 ? (
          <p>Корзина пуста</p>
        ) : (
          <div className="flex flex-col gap-4">
            {cart.map((cartItem) => (
              <CartItem
                key={cartItem.id_product}
                cartItem={cartItem}
                changeQuantity={changeQuantity}
                removeFromCart={removeFromCart}
              />
            ))}
            <h2 className="text-xl font-bold">Итого: {total} р</h2>
          </div>
        )}
      </div>
    </Container>
  );
}
