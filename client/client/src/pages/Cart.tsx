import Container from '../components/Container';
import { useCart } from '../context/CartContext';
import { Plus, Minus, Trash2 } from 'lucide-react';
export default function Cart() {
  const { cart, removeFromCart, changeQuantity, clearCart } = useCart();
  const total = cart.reduce((acc, item) => acc + item.price * item.quantity, 0);

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
            {cart.map((product) => (
              <div
                key={product.id}
                className="flex items-center justify-between shadow-sm p-4 rounded-2xl border-gray-200 border hover:shadow-lg transition-shadow duration-300"
              >
                <div>
                  <h2 className="text-xl font-bold">{product.name}</h2>
                  <p>{product.price} р</p>
                  <p>В наличии: {product.stock}</p>
                </div>
                <div className="flex items-center gap-2">
                  <p className="mx-2 text-xl">
                    {product.price * product.quantity} р.
                  </p>
                  <button
                    className=""
                    onClick={() =>
                      changeQuantity(product.id, product.quantity - 1)
                    }
                  >
                    <Minus />
                  </button>
                  <input
                    value={product.quantity}
                    className="w-12 border border-gray-200 text-xl text-right rounded-md px-2"
                    onChange={(e) =>
                      changeQuantity(product.id, parseInt(e.target.value, 10))
                    }
                  />
                  <button
                    onClick={() =>
                      changeQuantity(product.id, product.quantity + 1)
                    }
                  >
                    <Plus />
                  </button>
                  <button onClick={() => removeFromCart(product.id)}>
                    <Trash2 className="text-red-500 mx-2" />
                  </button>
                </div>
              </div>
            ))}
            <h2 className="text-xl font-bold">Итого: {total} р</h2>
          </div>
        )}
      </div>
    </Container>
  );
}
