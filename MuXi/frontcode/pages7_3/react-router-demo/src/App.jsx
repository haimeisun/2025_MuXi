// src/App.jsx
import React from "react";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import AuthGuard from "./utils/auth";
import Home from "./components/Home";
import ProductsList from "./components/ProductsList";
import ProductDetail from "./components/ProductDetail";
import Profile from "./components/Profile";
import Orders from "./components/Orders";
import Settings from "./components/Settings";
import Login from "./components/Login";
import NotFound from "./components/NotFound";

const App = () => {
  return (
    <Router>
      <div>
        <nav>
          <ul>
            <li>
              <Link to="/">首页</Link>
            </li>
            <li>
              <Link to="/products">商品列表</Link>
            </li>
            <li>
              <Link to="/profile">用户中心</Link>
            </li>
          </ul>
        </nav>

        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/products" element={<ProductsList />} />
          <Route path="/products/:id" element={<ProductDetail />} />

          {/* 用户中心路由，需要身份验证 */}
          <Route
            path="/profile"
            element={
              <AuthGuard>
                <Profile />
              </AuthGuard>
            }
          >
            {/* 嵌套路由 */}
            <Route index element={<Orders />} />
            <Route path="orders" element={<Orders />} />
            <Route path="settings" element={<Settings />} />
          </Route>

          <Route path="/login" element={<Login />} />

          {/* 404页面 */}
          <Route path="*" element={<NotFound />} />
        </Routes>
      </div>
    </Router>
  );
};

export default App;
