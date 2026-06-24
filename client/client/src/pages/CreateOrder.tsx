import { useState, useEffect } from 'react';
import type { Order, OrderFull } from '../types/order';
import { useAuth } from '../context/AuthContext';
import { apiCreateOrder, apiCreateOrderNoAuth } from '../api/order';
import { useCart } from '../context/CartContext';
import { useNavigate } from 'react-router-dom';

export default function CreateOrderPage() {
  const navigate = useNavigate();
  const { token } = useAuth();
  const { cart, clearCart } = useCart();
  const [order, setOrder] = useState<Order | null>(null);

  async function handleSumbit() {
    if (cart.length === 0) {
      console.error('Cart is empty');
      return;
    }

    addOrder(order as Order);
    navigate('/');
  }

  async function addOrder(order: Order) {
    if (token) {
      try {
        await apiCreateOrder(order);
      } catch (error) {
        console.error('Error creating order:', error);
      }
    } else {
      try {
        await apiCreateOrderNoAuth(order, cart);
        clearCart();
      } catch (error) {
        console.error('Error creating order without auth:', error);
      }
    }
  }

  return (
    <div>
      <h1>Create Order Page</h1>
      <form>
        <div>
          <input
            type="text"
            placeholder="First Name"
            value={order?.first_name || ''}
          />
          <input
            type="text"
            placeholder="Last Name"
            value={order?.last_name || ''}
          />
          <input type="email" placeholder="Email" value={order?.email || ''} />
          <input type="tel" placeholder="Phone" value={order?.phone || ''} />
        </div>
        <div>
          {cart.map((item) => (
            <div key={item.id_product}>
              <p>{item.name}</p>
              <p>{item.amount}</p>
              <p>{item.price}</p>
            </div>
          ))}
        </div>
        <button type="submit" onClick={handleSumbit}>
          Create Order
        </button>
      </form>
    </div>
  );
}
