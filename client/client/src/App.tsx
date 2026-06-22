import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Products from './pages/Products';
import Layout from './components/Layout';
import Index from './pages/Index';
import Auth from './pages/auth/Auth';
import Cart from './pages/Cart';

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
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
