// src/components/Profile.jsx
import { Outlet } from "react-router-dom";

const Profile = () => {
  return (
    <div>
      <h1>用户中心</h1>
      <nav>
        <ul>
          <li>
            <Link to="/profile/orders">我的订单</Link>
          </li>
          <li>
            <Link to="/profile/settings">账户设置</Link>
          </li>
        </ul>
      </nav>
      <Outlet />
    </div>
  );
};

export default Profile;
