import {
  ShoppingCartIcon,
  TextAlignJustify,
  Search,
  LogOut,
  CircleUser,
} from 'lucide-react';
import { useEffect, useState, useRef } from 'react';
import { getCategories } from '../api/categories';
import type { Category } from '../types/category';
import { useNavigate, useSearchParams, useLocation } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import { useCart } from '../context/CartContext';
import Container from './Container';

export default function Header() {
  const location = useLocation();
  const navigate = useNavigate();
  const { token, logout } = useAuth();
  const { cart } = useCart();
  const [catalogOpen, setCatalogOpen] = useState(false);
  const [categories, setCategories] = useState<Category[]>([]);
  const catalogRef = useRef<HTMLDivElement>(null);
  const [searchParams, setSearchParams] = useSearchParams();
  const [searchInput, setSearchInput] = useState(
    searchParams.get('search') || '',
  );

  useEffect(() => {
    async function fetchCategories() {
      try {
        setCategories([]);
        const data = await getCategories();
        setCategories(data);
      } catch (error) {
        console.error('Error fetching categories:', error);
      }
    }
    fetchCategories();
  }, []);

  useEffect(() => {
    function handleClickOutside(event: MouseEvent) {
      if (
        catalogRef.current &&
        !catalogRef.current.contains(event.target as Node)
      ) {
        setCatalogOpen(false);
      }
    }

    document.addEventListener('mousedown', handleClickOutside);

    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, []);

  return (
    <header className="flex items-center justify-between p-4 border-b">
      <h1 className="text-2xl font-bold">
        <a href="/" className="hover:text-blue-500">
          My Shop
        </a>
      </h1>
      <div className="flex items-center gap-2">
        <div ref={catalogRef}>
          <button
            className="
            flex items-center gap-2 bg-blue-500 text-white font-bold
            p-2 rounded-md hover:bg-blue-600 focus:outline-none 
            focus:ring-2 focus:ring-blue-600"
            onClick={() => setCatalogOpen(!catalogOpen)}
          >
            <TextAlignJustify className="w-5 h-5" /> Каталог
          </button>
          {catalogOpen && (
            <>
              {/* затемнение + размытие фона */}
              <div
                className="
                  fixed
                  inset-0
                  bg-black/1
                  backdrop-blur-xs
                  z-40
                "
                onClick={() => setCatalogOpen(false)}
              />

              {/* меню каталога */}
              <div
                className="
                  absolute
                  top-20
                  bg-white
                  rounded-md
                  shadow-lg
                  border
                  border-gray-200
                  z-50
                "
              >
                <ul>
                  {categories.map((category) => (
                    <li
                      key={category.id}
                      className="
                        p-2
                        cursor-pointer
                        hover:bg-blue-200
                      "
                      onClick={() => {
                        navigate(`/category/${category.slug}/`);
                        setCatalogOpen(false);
                      }}
                    >
                      {category.name}
                    </li>
                  ))}
                </ul>
              </div>
            </>
          )}
        </div>
        <form
          className="flex items-center gap-2"
          onSubmit={(e) => {
            e.preventDefault();
            const params = new URLSearchParams(searchParams);
            if (searchInput) {
              params.set('search', searchInput);
            } else {
              params.delete('search');
            }
            params.set('page', '1');
            setSearchParams(params);
            if (
              location.pathname !== '/category' &&
              location.pathname !== '/products'
            ) {
              navigate(`/products?${params.toString()}`);
            } else {
              setSearchParams(params);
            }
          }}
        >
          <input
            type="text"
            placeholder="Search..."
            className="w-80 border 
            border-gray-400 rounded-md p-2 
            focus:outline-none focus:ring-2 focus:ring-gray-500"
            value={searchInput}
            onChange={(e) => setSearchInput(e.target.value)}
          />
          <button
            type="submit"
            className="
            flex items-center gap-2 bg-blue-500 text-white font-bold
            p-2 rounded-md hover:bg-blue-600 focus:outline-none 
            focus:ring-2 focus:ring-blue-600"
          >
            <Search className="w-5 h-5" /> Search
          </button>
        </form>
      </div>
      <div className="flex items-center gap-2">
        <button
          className="flex items-center gap-2 bg-blue-500 text-white font-bold p-2 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-600"
          onClick={() => (window.location.href = '/cart')}
        >
          <ShoppingCartIcon className="w-5 h-5" /> Корзина {cart.length}
        </button>
        {token ? (
          <>
            <button
              className="flex items-center gap-2 bg-blue-500 text-white font-bold p-2 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-600"
              onClick={() => navigate('/profile')}
            >
              <div className="flex flex-row gap-2">
                <CircleUser />
                ЛК
              </div>
            </button>
            <button
              className="flex items-center gap-2 bg-blue-500 text-white font-bold p-2 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-600"
              onClick={() => {
                logout();
                navigate('/');
              }}
            >
              <LogOut />
            </button>
          </>
        ) : (
          <>
            <button
              className="flex items-center gap-2 bg-blue-500 text-white font-bold p-2 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-600"
              onClick={() => navigate('/auth')}
            >
              Войти
            </button>
          </>
        )}
      </div>
    </header>
  );
}
