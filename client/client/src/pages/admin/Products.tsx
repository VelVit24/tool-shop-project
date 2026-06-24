import { useState, use, useEffect } from 'react';
import { useSearchParams } from 'react-router-dom';
import type { Product } from '../../types/product';
import { ChevronLeftIcon, ChevronRightIcon } from 'lucide-react';
import {
  apiGetProducts,
  apiDeleteProduct,
  apiUpdateProduct,
  apiCheckSlug,
  apiCreateProduct,
} from '../../api/products';
import type { Category } from '../../types/category';

export function AdminProducts() {
  const [products, setProducts] = useState<Product[]>([]);
  const [searchParams, setSearchParams] = useSearchParams();
  const page = Number(searchParams.get('page') || 1);
  const [total, setTotal] = useState(0);
  const [product, setProduct] = useState<Product>();
  const [name, setName] = useState('');
  const [price, setPrice] = useState(0);
  const [slug, setSlug] = useState('');
  const [category, setCategory] = useState('');
  const [categories, setCategories] = useState<Category[]>([]);

  const [changeName, setChangeName] = useState(false);
  const [changePrice, setChangePrice] = useState(false);

  const limit = 30;
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
    async function fetchProducts() {
      try {
        const response = apiGetProducts(page, limit);
        setProducts((await response).products);
        setTotal((await response).total);
      } catch (error) {
        console.error('Error fetching products:', error);
      }
    }
    fetchProducts();
  }, [page]);

  async function handleSumbit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    const newProduct: Product = {
      id: 0,
      name: name,
      price: price,
      description: '',
      stock: 0,
      image_count: 0,
      category: category,
      slug: slug,
    };
    try {
      const response = await apiCreateProduct(newProduct);
      setProducts([...products, response]);
    } catch (error) {
      console.error('Error creating product:', error);
    }
  }

  async function checkSlug(name: string) {
    try {
      const response = await apiCheckSlug(name);
      setSlug(response.slug);
    } catch (error) {
      console.error('Error checking slug:', error);
    }
  }

  function handleDelete(productId: number) {
    async function deleteProduct() {
      try {
        await apiDeleteProduct(productId);
        setProducts(products.filter((product) => product.id !== productId));
      } catch (error) {
        console.error('Error deleting product:', error);
      }
    }
    deleteProduct();
  }
  function handleUpdate(product: Product) {
    async function updateProduct() {
      try {
        await apiUpdateProduct(product);
        setProducts(products.map((p) => (p.id === product.id ? product : p)));
      } catch (error) {
        console.error('Error updating product:', error);
      }
    }
    updateProduct();
  }

  return (
    <div className="flex-1">
      <div>
        <button></button>
      </div>
      <div className="">
        <div>
          <form onSubmit={handleSumbit}>
            <div>
              Название
              <input
                onBlur={(e) => {
                  setName(e.target.value);
                  checkSlug(e.target.value);
                }}
              />
            </div>
            <div>
              Цена
              <input
                onBlur={(e) => {
                  setPrice(Number(e.target.value));
                }}
              />
            </div>
            <div>
              Slug
              <input
                value={slug}
                onBlur={(e) => {
                  setSlug(e.target.value);
                }}
              />
            </div>
            <div>
              Категория
              <select
                onChange={(e) => {
                  const selectedCategory = categories.find(
                    (category) => category.id === Number(e.target.value),
                  );
                  if (selectedCategory) {
                    setCategory(selectedCategory.name);
                  }
                }}
              >
                {categories.map((category) => (
                  <option key={category.id} value={category.id}>
                    {category.name}
                  </option>
                ))}
              </select>
            </div>
            <button>Добавить</button>
          </form>
        </div>
      </div>
      <div className="flex flex-col gap-y-2">
        {products.map((product) => (
          <div key={product.id}>
            <div>
              {changeName ? (
                <input
                  type="text"
                  value={product.name}
                  onChange={(e) =>
                    setProducts(
                      products.map((p) =>
                        p.id === product.id
                          ? { ...p, name: e.target.value }
                          : p,
                      ),
                    )
                  }
                  onBlur={() => {
                    handleUpdate(product);
                  }}
                />
              ) : (
                <p>{product.name}</p>
              )}
              <button onClick={() => setChangeName(!changeName)}>
                {changeName ? 'Save' : 'Edit'}
              </button>
            </div>
            <div>
              {changePrice ? (
                <input
                  type="text"
                  value={product.price}
                  onChange={(e) =>
                    setProducts(
                      products.map((p) =>
                        p.id === product.id
                          ? { ...p, price: Number(e.target.value) }
                          : p,
                      ),
                    )
                  }
                  onBlur={() => {
                    handleUpdate(product);
                  }}
                />
              ) : (
                <p>{product.price}</p>
              )}
              <button onClick={() => setChangePrice(!changePrice)}>
                {changePrice ? 'Save' : 'Edit'}
              </button>
            </div>
            <div>
              <button onClick={() => handleDelete(product.id)}>Delete</button>
            </div>
          </div>
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
