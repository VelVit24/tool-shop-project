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
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import { useCart } from '../context/CartContext';

export default function Header() {
  const navigate = useNavigate();
  const { token, logout } = useAuth();
  const { cart } = useCart();
  const [catalogOpen, setCatalogOpen] = useState(false);
  const [categories, setCategories] = useState<Category[]>([]);
  const catalogRef = useRef<HTMLDivElement>(null);

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
            className="flex items-center gap-2 bg-gray-200  p-2 rounded-md hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-600"
            onClick={() => setCatalogOpen(!catalogOpen)}
          >
            <TextAlignJustify className="w-5 h-5" /> Каталог
          </button>
          {catalogOpen && (
            <div className="absolute top-16 left-0 w-full bg-white border-t shadow-md z-10">
              <ul>
                {categories.map((category) => (
                  <li
                    key={category.id}
                    className="p-2 hover:bg-gray-100 cursor-pointer"
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
          )}
        </div>
        <form className="flex items-center gap-2">
          <input
            type="text"
            placeholder="Search..."
            className="w-80 border border-gray-400 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-gray-500"
          />
          <button
            type="submit"
            className="flex items-center gap-2 bg-gray-200 p-2 rounded-md hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-600"
          >
            <Search className="w-5 h-5" /> Search
          </button>
        </form>
      </div>
      <div className="flex items-center gap-2">
        <button
          className="flex items-center gap-2 bg-gray-200  p-2 rounded-md hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-600"
          onClick={() => (window.location.href = '/cart')}
        >
          <ShoppingCartIcon className="w-5 h-5" /> Корзина {cart.length}
        </button>
        {token ? (
          <>
            <button
              className="flex items-center gap-2 bg-gray-200 p-2 rounded-md hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-600"
              onClick={() => navigate('/profile')}
            >
              <div className="flex flex-row gap-2">
                <CircleUser />
                ЛК
              </div>
            </button>
            <button
              className="flex items-center gap-2 bg-gray-200 p-2 rounded-md hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-600"
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
              className="flex items-center gap-2 bg-gray-200 p-2 rounded-md hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-600"
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
