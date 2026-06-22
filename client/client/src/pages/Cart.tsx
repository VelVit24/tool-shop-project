import { useState } from 'react';
import Container from '../components/Container';
import { useCart } from '../context/CartContext';
import { Plus, Minus, Trash2 } from 'lucide-react';
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
            {cart.map((cartItem) => {
              const [localAmount, setLocalAmount] = useState(
                cartItem.amount.toString(),
              );
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
                  <div className="flex items-center gap-2">
                    <p className="mx-2 text-xl">
                      {cartItem.price * cartItem.amount} р.
                    </p>
                    <button
                      onClick={() => {
                        const newAmount = cartItem.amount - 1;
                        if (newAmount >= 1) {
                          setLocalAmount(newAmount.toString());
                          changeQuantity(cartItem.id_product, newAmount);
                        }
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
                    <button onClick={() => removeFromCart(cartItem.id_product)}>
                      <Trash2 className="text-red-500 mx-2" />
                    </button>
                  </div>
                </div>
              );
            })}
            <h2 className="text-xl font-bold">Итого: {total} р</h2>
          </div>
        )}
      </div>
    </Container>
  );
}
