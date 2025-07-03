// src/utils/auth.jsx
import { useEffect } from "react";
import { useNavigate, useLocation } from "react-router-dom";

const AuthGuard = ({ children }) => {
  const navigate = useNavigate();
  const location = useLocation();

  // 检查用户是否已登录
  const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";

  useEffect(() => {
    if (!isLoggedIn) {
      navigate("/login", {
        state: {
          from: location,
          message: "请先登录以访问此页面",
        },
      });
    }
  }, [navigate, location, isLoggedIn]);

  return isLoggedIn ? children : null;
};

export default AuthGuard;
