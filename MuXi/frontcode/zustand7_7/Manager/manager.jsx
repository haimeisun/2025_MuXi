import { useEffect } from "react";
import useUserStore from "./useUserStore";

const UserProfile = () => {
  const { user, fetchUser, isUserLoaded, userFullName } = useUserStore();

  // 组件卸载时清理订阅（若在组件内订阅）
  useEffect(() => {
    const unsubscribe = useUserStore.subscribe(
      (changes) => console.log("Component-level user change:", changes),
      (state) => state.user
    );
    return () => unsubscribe();
  }, []);

  return (
    <div>
      <h2>User Profile</h2>
      <button onClick={fetchUser}>Load User Info</button>
      {!isUserLoaded ? (
        <p>Loading user...</p>
      ) : (
        <div>
          <p>Full Name: {userFullName}</p>
          <p>Age: {user.age}</p>
          <p>Email: {user.email}</p>
        </div>
      )}
    </div>
  );
};

export default UserProfile;
