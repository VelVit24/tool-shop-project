import { useState } from 'react';
import { loginapi } from '../../api/auth';
import { useNavigate } from 'react-router-dom';
import { Eye, EyeOff } from 'lucide-react';
import { useAuth } from '../../context/AuthContext';

export default function Login() {
  const navigate = useNavigate();
  const { login } = useAuth();
  const [loginStr, setLoginStr] = useState('');
  const [type, setType] = useState<'email' | 'phone'>('email');
  const [password, setPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false);
  const [error, setError] = useState<{
    email?: string;
    phone?: string;
    password?: string;
  }>({});

  function detectType(value: string) {
    if (/[a-zA-Z@]/.test(value)) {
      return 'email';
    }

    return 'phone';
  }

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
  function isPhone(value: string) {
    let numbers = value.replace(/\D/g, '');
    if (numbers.startsWith('8') || numbers.startsWith('7')) {
      return true;
    }
    return false;
  }

  async function handleCheckLogin() {
    setError((prev) => ({ ...prev, email: undefined, phone: undefined }));
    if (type === 'email') {
      if (!validateEmail(loginStr)) {
        setError((prev) => ({
          ...prev,
          email: 'Некорректный email',
        }));
        return;
      }
    } else {
      if (!validatePhone(loginStr)) {
        setError((prev) => ({
          ...prev,
          phone: 'Некорректный телефон',
        }));
        return;
      }
    }
  }

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError({});

    if (!loginStr) {
      setError((prev) => ({
        ...prev,
        email: 'Email или телефон не может быть пустым',
      }));
      return;
    }

    if (password.length < 8) {
      setError((prev) => ({
        ...prev,
        password: 'Пароль должен быть минимум 8 символов',
      }));
      return;
    }

    let request: any = {};
    if (type === 'email') {
      request.email = loginStr;
    } else {
      request.phone = loginStr.replace(/\D/g, '');
    }
    request.password = password;

    try {
      const responce = await loginapi(request);
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

  async function handleCheckPassword() {
    setError((prev) => ({ ...prev, password: undefined }));
    if (password.length < 8) {
      setError((prev) => ({
        ...prev,
        password: 'Пароль должен быть минимум 8 символов',
      }));
    }
  }

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-3 w-80">
      <div className="flex flex-col gap-1 relative">
        Email или телефон
        <input
          type="text"
          className={
            error.email
              ? 'border border-red-400 p-2 rounded'
              : 'border border-gray-200 p-2 rounded'
          }
          value={loginStr}
          onChange={(e) => {
            let raw = e.target.value;
            const detected = detectType(raw);
            setType(detected);
            if (detected === 'phone') {
              setLoginStr(formatPhone(raw));
            } else {
              setLoginStr(raw);
            }
          }}
          onBlur={handleCheckLogin}
          placeholder="Email"
        />
        {error.email && <p className="text-red-500 text-sm">{error.email}</p>}
      </div>
      <div className="flex flex-col gap-1 relative">
        Пароль
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
        Войти
      </button>
    </form>
  );
}
