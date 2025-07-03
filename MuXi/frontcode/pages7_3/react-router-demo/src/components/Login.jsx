// src/components/Login.jsx
import { useNavigate } from "react-router-dom";

const Login = () => {
  const navigate = useNavigate();

  const handleLogin = () => {
    // 模拟登录成功
    localStorage.setItem("isLoggedIn", "true");
    navigate("/profile");
  };

  return (
    <div>
      <h1>登录</h1>
      <p>请登录以访问用户中心</p>
      <button onClick={handleLogin}>登录</button>
    </div>
  );
};

export default Login;
