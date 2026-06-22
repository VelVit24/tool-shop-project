import type { Product } from '../types/product';
import { createContext, useState, useEffect, useContext } from 'react';
import { useAuth } from './AuthContext';
import {
  getCart,
  addCartItem,
  removeCartItem,
  changeCartItemQuantity,
  clearCartItems,
} from '../api/cart';

type CartContextType = {
  cart: CartItem[];
  addToCart: (product: Product) => void;
  removeFromCart: (id: number) => void;
  changeQuantity: (id: number, quantity: number) => void;
  clearCart: () => void;
  getCartItem: (productId: number) => CartItem | undefined;
};

export type CartItem = {
  id_product: number;
  name: string;
  price: number;
  stock: number;
  image_count: number;
  amount: number;
  is_in_stock: boolean;
};

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
        const saved = localStorage.getItem('cart');
        setCart(saved ? JSON.parse(saved) : []);
      }
    }
    fetchCart();
  }, [token]);

  async function addToCart(product: Product) {
    if (token) {
      await addCartItem(product.id, 1);

      const data = await getCart();
      setCart(data);
    } else {
      setCart((prev) => {
        const exists = prev.find((item) => item.id_product === product.id);
        if (exists) {
          return prev.map((item) =>
            item.id_product === product.id
              ? { ...item, amount: item.amount + 1 }
              : item,
          );
        } else {
          return [
            ...prev,
            {
              id_product: product.id,
              name: product.name,
              price: product.price,
              stock: product.stock,
              image_count: product.image_count,
              amount: 1,
              is_in_stock: product.stock > 0,
            },
          ];
        }
      });
    }
  }

  async function changeQuantity(id: number, quantity: number) {
    quantity = normalizeQuantity(quantity);
    if (token) {
      try {
        if (quantity < 1) {
          return;
        }
        await changeCartItemQuantity(id, quantity);
        const data = await getCart();
        setCart(data);
      } catch (error) {
        console.error(
          'Ошибка при изменении количества товара в корзине',
          error,
        );
      }
    } else {
      setCart((prevCart) =>
        prevCart.map((item) =>
          item.id_product === id ? { ...item, amount: quantity } : item,
        ),
      );
    }
  }
  function normalizeQuantity(q: number) {
    if (q < 1) return 1;
    return Math.floor(q);
  }

  async function removeFromCart(id: number) {
    if (token) {
      try {
        await removeCartItem(id);
        const data = await getCart();
        setCart(data);
      } catch (error) {
        console.error('Ошибка при удалении товара из корзины', error);
      }
    } else {
      setCart((prevCart) =>
        prevCart.filter((product) => product.id_product !== id),
      );
    }
  }

  async function clearCart() {
    if (token) {
      console.log('ПОЧЕМУ ЗДЕСЬ ТОКЕН:', token);
      try {
        await clearCartItems();
        setCart([]);
      } catch (error) {
        console.error('Ошибка при очистке корзины', error);
      }
    } else {
      localStorage.removeItem('cart');
      setCart([]);
    }
  }
  function getCartItem(productId: number): CartItem | undefined {
    return cart.find((item) => item.id_product === productId);
  }

  return (
    <CartContext.Provider
      value={{
        cart,
        addToCart,
        removeFromCart,
        changeQuantity,
        clearCart,
        getCartItem,
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
