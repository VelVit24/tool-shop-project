import { useState } from 'react';
import { checkEmail, checkPhone, register } from '../../api/auth';
import { useNavigate } from 'react-router-dom';
import { CircleCheck, Eye, EyeOff } from 'lucide-react';
import { useAuth } from '../../context/AuthContext';

export default function Register() {
  const navigate = useNavigate();
  const { login } = useAuth();
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [email, setEmail] = useState('');
  const [phone, setPhone] = useState('');
  const [password, setPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false);
  const [error, setError] = useState<{
    email?: string;
    phone?: string;
    password?: string;
  }>({});

  function validateEmail(email: string) {
    const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return re.test(email);
  }
  function validatePhone(phone: string) {
    const re = /^\+7 \(\d{3}\) \d{3}-\d{2}-\d{2}$/.test(phone);
    return re;
  }
  function formatPhone(value: string) {
    let numbers = value.replace(/\D/g, '');
    if (numbers.startsWith('8')) {
      numbers = '7' + numbers.slice(1);
    }
    numbers = numbers.slice(0, 11);
    let result = '';
    if (numbers.length >= 1) {
      result += '+7';
    }
    if (numbers.length > 1) {
      result += ' (' + numbers.slice(1, 4);
    }
    if (numbers.length > 4) {
      result += ') ' + numbers.slice(4, 7);
    }
    if (numbers.length > 7) {
      result += '-' + numbers.slice(7, 9);
    }
    if (numbers.length > 9) {
      result += '-' + numbers.slice(9, 11);
    }
    return result;
  }

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError({});

    if (password.length < 8) {
      setError({ password: 'Пароль должен быть минимум 8 символов' });
      return;
    }

    try {
      const responce = await register(
        email,
        password,
        phone,
        firstName,
        lastName,
      );
      login(responce.data.token);
      navigate('/');
    } catch (error: any) {
      if (error.response?.data?.error === 'invalid email') {
        setError({ email: 'Некорректный email' });
      } else if (error.response?.data?.error === 'invalid phone') {
        setError({ phone: 'Некорректный телефон' });
      } else {
        setError({ email: 'Ошибка при регистрации' });
      }
    }
  }

  async function handleCheckEmail() {
    setError((prev) => ({ ...prev, email: undefined }));
    if (!email) {
      setError((prev) => ({ ...prev, email: 'Email не может быть пустым' }));
      return;
    }
    if (!validateEmail(email)) {
      setError((prev) => ({ ...prev, email: 'Некорректный email' }));
      return;
    }
    try {
      await checkEmail(email);
    } catch (error: any) {
      if (error.response?.data?.error === 'invalid email') {
        setError((prev) => ({ ...prev, email: 'Некорректный email' }));
      } else if (error.response?.data?.error === 'email already exists') {
        setError((prev) => ({
          ...prev,
          email: 'Этот email уже зарегистрирован',
        }));
      } else {
        setError((prev) => ({ ...prev, email: 'Ошибка при проверке email' }));
      }
    }
  }
  async function handleCheckPhone() {
    setError((prev) => ({ ...prev, phone: undefined }));
    if (!phone) return;
    if (!validatePhone(phone)) {
      setError((prev) => ({ ...prev, phone: 'Некорректный телефон' }));
      return;
    }
    try {
      await checkPhone(phone);
    } catch (error: any) {
      if (error.response?.data?.error === 'invalid phone') {
        setError((prev) => ({ ...prev, phone: 'Некорректный телефон' }));
      } else if (error.response?.data?.error === 'phone already exists') {
        setError((prev) => ({
          ...prev,
          phone: 'Этот телефон уже зарегистрирован',
        }));
      } else {
        setError((prev) => ({
          ...prev,
          phone: 'Ошибка при проверке телефона',
        }));
      }
    }
  }
  async function handleCheckPassword() {
    setError((prev) => ({ ...prev, password: undefined }));
    if (!password) {
      setError((prev) => ({
        ...prev,
        password: 'Пароль не может быть пустым',
      }));
      return;
    }
    if (password.length < 8) {
      setError((prev) => ({
        ...prev,
        password: 'Пароль должен быть минимум 8 символов',
      }));
    }
  }

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-3 w-80">
      <div className="flex flex-col gap-1">
        Имя
        <input
          type="text"
          className="border border-gray-200  p-2 rounded"
          value={firstName}
          onChange={(e) => setFirstName(e.target.value)}
          placeholder="Имя"
        />
      </div>
      <div className="flex flex-col gap-1">
        Фамилия
        <input
          type="text"
          className="border border-gray-200  p-2 rounded"
          value={lastName}
          onChange={(e) => setLastName(e.target.value)}
          placeholder="Фамилия"
        />
      </div>
      <div className="flex flex-col gap-1 relative">
        Email *
        <input
          type="email"
          className={
            error.email
              ? 'border border-red-400 p-2 rounded'
              : 'border border-gray-200 p-2 rounded'
          }
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          onBlur={handleCheckEmail}
          placeholder="Email"
        />
        {error.email && <p className="text-red-500 text-sm">{error.email}</p>}
        {!error.email && email && (
          <div className="absolute right-2 top-7/10 -translate-y-1/2">
            <CircleCheck className="text-green-500 size-5" />
          </div>
        )}
      </div>
      <div className="flex flex-col gap-1 relative">
        Телефон
        <input
          type="tel"
          className={
            error.phone
              ? 'border border-red-400 p-2 rounded'
              : 'border border-gray-200 p-2 rounded'
          }
          value={phone}
          onChange={(e) => {
            setPhone(formatPhone(e.target.value));
          }}
          onBlur={handleCheckPhone}
          placeholder="Телефон"
        />
        {!error.phone && phone && (
          <div className="absolute right-2 top-7/10 -translate-y-1/2">
            <CircleCheck className="text-green-500 size-5" />
          </div>
        )}
        {error.phone && <p className="text-red-500 text-sm">{error.phone}</p>}
      </div>
      <div className="flex flex-col gap-1 relative">
        Пароль *
        <input
          type={showPassword ? 'text' : 'password'}
          className={
            error.password
              ? 'border border-red-400 p-2 rounded'
              : 'border border-gray-200 p-2 rounded'
          }
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          onBlur={handleCheckPassword}
          placeholder="Пароль"
        />
        <button
          type="button"
          className="absolute right-2 top-7/10 -translate-y-1/2"
          onClick={() => setShowPassword(!showPassword)}
        >
          {showPassword ? (
            <EyeOff className="size-5" />
          ) : (
            <Eye className="size-5" />
          )}
        </button>
      </div>
      {error.password && (
        <p className="text-red-500 text-sm">{error.password}</p>
      )}
      <button className="w-full bg-gray-200 p-2 rounded hover:bg-gray-400 active:bg-gray-600">
        Зарегистрироваться
      </button>
    </form>
  );
}
