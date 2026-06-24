import { useEffect, useState } from 'react';
import type { Product } from '../types/product';
import { apiGetProducts } from '../api/products';
import ProductCard from '../components/ProductCard';
import { ChevronLeftIcon, ChevronRightIcon } from 'lucide-react';
import { useSearchParams, useParams } from 'react-router-dom';
import Container from '../components/Container';

export default function Products() {
  const [products, setProducts] = useState<Product[]>([]);
  const [searchParams, setSearchParams] = useSearchParams();
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const { category_slug } = useParams<{ category_slug: string }>();
  const filters = {
    page: Number(searchParams.get('page')) || 1,
    limit: 10,
    category: category_slug,
    priceFrom: Number(searchParams.get('priceFrom')) || undefined,
    priceTo: Number(searchParams.get('priceTo')) || undefined,
    inStock: searchParams.get('inStock') === 'true',
    sort: searchParams.get('sort') || 'name_asc',
    search: searchParams.get('search') || undefined,
  };
  const [priceFromInput, setPriceFromInput] = useState('');
  const [priceToInput, setPriceToInput] = useState('');

  // логика страниц
  const limit = 20;
  const pagesCount = Math.ceil(total / limit);
  const canGoBack = filters.page > 1;
  const canGoNext = filters.page < pagesCount;
  const delta = 2; // количество страниц слева и справа от текущей страницы
  const start = Math.max(1, filters.page - delta);
  const end = Math.min(pagesCount, filters.page + delta);
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
  async function fetchProducts() {
    try {
      setLoading(true);
      setError('');
      console.log(category_slug);
      const data = await apiGetProducts(filters);
      setProducts(data.products);
      setTotal(data.total);
    } catch (error) {
      setError('Ошибка при загрузке товаров');
    } finally {
      setLoading(false);
    }
  }
  // загрузка товаров
  useEffect(() => {
    fetchProducts();
  }, [searchParams, category_slug]);
  useEffect(() => {
    if (!searchParams.get('sort')) {
      const params = new URLSearchParams(searchParams);
      params.set('sort', 'name_asc');
      setSearchParams(params, { replace: true });
    }
  }, []);

  function updateFilter(key: string, value: string) {
    const params = new URLSearchParams(searchParams);
    if (value) {
      params.set(key, value);
    } else {
      params.delete(key);
    }
    if (key !== 'page') {
      params.set('page', '1');
    }
    setSearchParams(params);
  }

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
            <div>
              <h2 className="text-lg font-bold mb-2">Сортировать по</h2>
              <select
                className="w-full border border-gray-200 rounded p-1"
                value={filters.sort}
                onChange={(e) => {
                  updateFilter('sort', e.target.value);
                }}
              >
                <option value="price_asc">Цена: по возрастанию</option>
                <option value="price_desc">Цена: по убыванию</option>
                <option value="name_asc">Название: А-Я</option>
                <option value="name_desc">Название: Я-А</option>
              </select>
            </div>
            <div className="flex flex-col mt-2">
              <h2 className="text-lg font-bold mb-2">Фильтры</h2>
              <label className="flex flex-row gap-x-2 p-2 hover:bg-blue-100 rounded-md transition-colors duration-200">
                <input
                  className="w-5 rounded-2xl"
                  type="checkbox"
                  checked={filters.inStock}
                  onChange={(e) =>
                    updateFilter('inStock', e.target.checked ? 'true' : '')
                  }
                />
                <span>Только в наличии</span>
              </label>
              <label className="border-t border-gray-200 p-2">
                <h2 className="text-lg font-bold mb-2">Цена</h2>
                <div className="flex flex-row gap-2">
                  <input
                    className=" 
                    w-10 flex-1 border 
                    border-gray-200 rounded
                    p-1
                    "
                    type="number"
                    value={priceFromInput}
                    placeholder="От"
                    onChange={(e) => setPriceFromInput(e.target.value)}
                    onBlur={(e) => updateFilter('priceFrom', priceFromInput)}
                    onKeyDown={(e) => {
                      if (e.key === 'Enter')
                        updateFilter('priceFrom', priceFromInput);
                    }}
                  />
                  <input
                    className="
                    w-10 flex-1 border 
                    border-gray-200 rounded
                    p-1
                    "
                    type="number"
                    value={priceToInput}
                    placeholder="До"
                    onChange={(e) => setPriceToInput(e.target.value)}
                    onBlur={(e) => updateFilter('priceTo', priceToInput)}
                    onKeyDown={(e) => {
                      if (e.key === 'Enter')
                        updateFilter('priceTo', priceToInput);
                    }}
                  />
                </div>
              </label>
            </div>
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
                onClick={() => updateFilter('page', String(filters.page - 1))}
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
                      onClick={() => updateFilter('page', String(p as number))}
                      className={`
                  border rounded-md
                  w-6.5 h-6.5
                  mx-1
                  flex items-center justify-center
                  hover:bg-gray-200
                  active:bg-gray-400

                  ${filters.page === p ? 'bg-black text-white' : ''}
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
                onClick={() => updateFilter('page', String(filters.page + 1))}
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
