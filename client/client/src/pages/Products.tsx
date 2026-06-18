import { useEffect, useState } from 'react';
import type { Product } from '../types/product';
import api from '../api/axios';
import ProductCard from '../components/ProductCard';

export default function Products() {
  const [products, setProducts] = useState<Product[]>([]);
  useEffect(() => {
    api.get('/products?page=1&limit=10').then((responce) => {
      setProducts(responce.data.products);
    });
  }, []);
  return (
    <div>
      <h1>Товары</h1>
      {products.map((product) => (
        <ProductCard key={product.id} product={product} />
      ))}
    </div>
  );
}
