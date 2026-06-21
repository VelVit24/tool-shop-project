import { useState } from 'react';
import Login from './Login';
import Register from './Register';

export default function Auth() {
  const [auth, setAuth] = useState('login');

  return (
    <div className="flex flex-col justify-center items-center mt-10 w-80 mx-auto">
      <div className="flex mb-5 w-full gap-2">
        <button
          className={
            auth == 'login'
              ? 'w-full bg-gray-200 text-black font-bold text-2xl p-2 rounded hover:bg-gray-400 active:bg-gray-600'
              : 'w-full bg-white text-gray-500 font-normal text-2xl p-2 rounded hover:bg-gray-200 active:bg-gray-600'
          }
          onClick={() => {
            setAuth('login');
          }}
        >
          Вход
        </button>
        <button
          className={
            auth == 'register'
              ? 'w-full bg-gray-200 text-black font-bold text-2xl p-2 rounded hover:bg-gray-400 active:bg-gray-600'
              : 'w-full bg-white text-gray-500 font-normal text-2xl p-2 rounded hover:bg-gray-200 active:bg-gray-600'
          }
          onClick={() => {
            setAuth('register');
          }}
        >
          Регистрация
        </button>
      </div>
      {auth === 'login' ? <Login /> : <Register />}
    </div>
  );
}
