import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Products from './pages/Products';
import Layout from './components/Layout';
import Index from './pages/Index';
import Auth from './pages/auth/Auth';
import Cart from './pages/Cart';
import { OrderPage } from './pages/Order';
import OrderCreatePage from './pages/CreateOrder';
import './index.css';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route element={<Layout />}>
          <Route path="/" element={<Index />} />
          <Route path="/products" element={<Products />} />
          <Route path="/category/:category_slug/" element={<Products />} />
          <Route path="/auth" element={<Auth />} />
          <Route path="/cart" element={<Cart />} />
          <Route path="/orders" element={<OrderPage />} />
          <Route path="/orders/create" element={<OrderCreatePage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
