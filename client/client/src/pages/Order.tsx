import { use, useState, useEffect } from 'react';
import type { Order, OrderFull } from '../types/order';
import { apiGetOrders } from '../api/order';
import { useSearchParams } from 'react-router-dom';
import { ChevronLeftIcon, ChevronRightIcon } from 'lucide-react';
export function OrderPage() {
  const [orders, setOrders] = useState<OrderFull[] | null>(null);
  const [searchParams, setSearchParams] = useSearchParams();
  const page = Number(searchParams.get('page') || 1);
  const [total, setTotal] = useState(0);

  const limit = 12;
  const pagesCount = Math.ceil(total / limit);
  const canGoBack = page > 1;
  const canGoNext = page < pagesCount;
  const delta = 2; // количество страниц слева и справа от текущей страницы
  const start = Math.max(1, page - delta);
  const end = Math.min(pagesCount, page + delta);
  const pages = [];
  for (let i = start; i <= end; i++) {
    pages.push(i);
  }
  const pagination = [];
  if (start > 1) {
    pagination.push(1);
    if (start > 2) {
      pagination.push('...');
    }
  }
  for (let i = start; i <= end; i++) {
    pagination.push(i);
  }
  if (end < pagesCount) {
    if (end < pagesCount - 1) {
      pagination.push('...');
    }
    pagination.push(pagesCount);
  }

  useEffect(() => {
    async function fetchOrders() {
      try {
        const data = await apiGetOrders(page, limit);
        setOrders(data.orders);
        setTotal(data.total);
      } catch (error) {
        console.error('Error fetching orders:', error);
      }
    }
    fetchOrders();
  }, [page]);

  return (
    <div>
      <div>
        {orders ? (
          orders.map((item) => (
            <div key={item.order.id}>
              <p>{item.order.created_at}</p>
              <div>
                {item.cart_items.map((cartItem) => (
                  <div key={cartItem.id_product}>
                    <p>{cartItem.name}</p>
                    <p>{cartItem.amount}</p>
                    <p>{cartItem.price}</p>
                  </div>
                ))}
              </div>
            </div>
          ))
        ) : (
          <p>No orders found</p>
        )}
      </div>
      <div className="flex items-center justify-center mt-4">
        <button
          className="border rounded-md flex items-center justify-center w-6.5 h-6.5 mx-1 hover:bg-gray-200 active:bg-gray-400 not-hover:bg-white"
          hidden={!canGoBack}
          disabled={!canGoBack}
          onClick={() => setSearchParams({ page: String(page - 1) })}
        >
          <ChevronLeftIcon size={15} />
        </button>
        <div className="flex items-center justify-center">
          {pagination.map((p, index) => {
            if (p === '...') {
              return (
                <span key={index} className="mx-1">
                  ...
                </span>
              );
            }
            return (
              <button
                key={index}
                onClick={() => setSearchParams({ page: String(p as number) })}
                className={`
                  border rounded-md
                  w-6.5 h-6.5
                  mx-1
                  flex items-center justify-center
                  hover:bg-gray-200
                  active:bg-gray-400

                  ${page === p ? 'bg-black text-white' : ''}
                `}
              >
                {p}
              </button>
            );
          })}
        </div>
        <button
          className="border rounded-md flex items-center justify-center w-6.5 h-6.5  p-0.5 mx-1 hover:bg-gray-200 active:bg-gray-400 not-hover:bg-white"
          disabled={!canGoNext}
          hidden={!canGoNext}
          onClick={() => setSearchParams({ page: String(page + 1) })}
        >
          <ChevronRightIcon size={15} />
        </button>
      </div>
    </div>
  );
}
