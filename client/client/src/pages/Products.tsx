import { useEffect, useState } from 'react';
import type { Product } from '../types/product';
import { getProducts } from '../api/products';
import ProductCard from '../components/ProductCard';
import { ChevronLeftIcon, ChevronRightIcon } from 'lucide-react';
import './products.css';
import { useSearchParams, useParams } from 'react-router-dom';
import Container from '../components/Container';

export default function Products() {
  const [products, setProducts] = useState<Product[]>([]);
  const [searchParams, setSearchParams] = useSearchParams();
  const page = Number(searchParams.get('page') || 1);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const { category_slug } = useParams<{ category_slug: string }>();

  // логика страниц
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

  // загрузка товаров
  useEffect(() => {
    async function fetchProducts() {
      try {
        setLoading(true);
        setError('');
        setProducts([]);
        console.log(category_slug);
        const data = await getProducts(page, limit, category_slug);
        setProducts(data.products);
        setTotal(data.total);
      } catch (error) {
        setError('Ошибка при загрузке товаров');
      } finally {
        setLoading(false);
      }
    }
    fetchProducts();
  }, [page, category_slug]);
  if (loading) {
    return <div>Загрузка...</div>;
  }
  if (error) {
    return <div>{error}</div>;
  }
  return (
    <Container>
      <div className="flex flex-col gap-y-4">
        <h1 className="text-3xl font-bold text-center mt-5">Товары</h1>

        <div className="flex gap-x-2">
          <aside
            className="
            flex-none 
            w-60 border 
            border-gray-200 
            p-4 
            rounded-xl
            shadow-sm"
          >
            <ul>
              <li>dsafdsfsd</li>
              <li>dsafdsfsd</li>
              <li>dsafdsfsd</li>
              <li>dsafdsfsd</li>
            </ul>
          </aside>
          <div className="flex-1">
            <div className="flex flex-col gap-y-2">
              {products.map((product) => (
                <ProductCard key={product.id} product={product} />
              ))}
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
                      onClick={() =>
                        setSearchParams({ page: String(p as number) })
                      }
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
        </div>
      </div>
    </Container>
  );
}
