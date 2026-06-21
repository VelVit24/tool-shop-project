import type { Product } from '../types/product';
import { createContext, useState, useEffect, useContext } from 'react';

type CartContextType = {
  cart: Product[];
  addToCart: (product: Product) => void;
  //   removeFromCart: (productId: number) => void;
  //   clearCart: () => void;
};

const CartContext = createContext<CartContextType | undefined>(undefined);

export function CartProvider({ children }: { children: React.ReactNode }) {
  const [cart, setCart] = useState<Product[]>([]);
  function addToCart(product: Product) {
    setCart((prevCart) => [...prevCart, product]);
  }

  return (
    <CartContext.Provider
      value={{
        cart,
        addToCart,
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
