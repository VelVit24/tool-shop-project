import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Products from './pages/Products';
import Layout from './components/Layout';
import Index from './pages/Index';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route element={<Layout />}>
          <Route path="/" element={<Index />} />
          <Route path="/products" element={<Products />} />
          <Route path="/products/:category" element={<Products />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
