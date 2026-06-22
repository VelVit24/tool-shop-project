import type { Product } from '../types/product';
import { createContext, useState, useEffect, useContext } from 'react';
import { useAuth } from './AuthContext';
import { getCart, addCartItem, removeCartItem } from '../api/cart';

type CartContextType = {
  cart: CartItem[];
  addToCart: (product: Product) => void;
  removeFromCart: (id: number) => void;
  changeQuantity: (id: number, quantity: number) => void;
  clearCart: () => void;
};
export type CartItem = Product & { quantity: number };

const CartContext = createContext<CartContextType | undefined>(undefined);

export function CartProvider({ children }: { children: React.ReactNode }) {
  const [cart, setCart] = useState<CartItem[]>(() => {
    const saved = localStorage.getItem('cart');
    return saved ? JSON.parse(saved) : [];
  });

  const { token } = useAuth();

  useEffect(() => {
    localStorage.setItem('cart', JSON.stringify(cart));
  }, [cart]);
  useEffect(() => {
    async function fetchCart() {
      if (token) {
        try {
          const data = await getCart();
          setCart(data);
        } catch (error) {
          console.error('Ошибка при загрузке корзины', error);
        }
      } else {
        loadCart();
      }
    }
    fetchCart();
  }, [token]);

  function loadCart() {
    const saved = localStorage.getItem('cart');
    if (saved) {
      setCart(JSON.parse(saved));
    }
  }

  function addToCart(product: Product) {
    if (token) {
      addCartItem(product.id, 1).then(() => {
        getCart().then((res) => setCart(res.data.items));
      });
    } else {
      setCart((prevCart) => {
        const exists = prevCart.find((item) => item.id === product.id);
        if (exists) {
          return prevCart.map((item) =>
            item.id === product.id
              ? { ...item, quantity: item.quantity + 1 }
              : item,
          );
        } else {
          return [...prevCart, { ...product, quantity: 1 }];
        }
      });
    }
  }

  function changeQuantity(id: number, quantity: number) {
    setCart((prevCart) =>
      prevCart.map((item) =>
        item.id === id
          ? { ...item, quantity: quantity < 1 ? 1 : quantity }
          : item,
      ),
    );
  }

  function removeFromCart(id: number) {
    setCart((prevCart) => prevCart.filter((product) => product.id !== id));
  }

  function clearCart() {
    setCart([]);
  }

  return (
    <CartContext.Provider
      value={{
        cart,
        addToCart,
        removeFromCart,
        changeQuantity,
        clearCart,
      }}
    >
      {children}
    </CartContext.Provider>
  );
}

export function useCart() {
  const context = useContext(CartContext);
  if (!context) {
    throw new Error('useCart must be used within a CartProvider');
  }
  return context;
}
