import { create } from "zustand";

const useUserStore = create((set, get) => ({
  user: {}, // 初始用户信息为空对象
  // 异步加载用户信息
  fetchUser: async () => {
    // 模拟API请求
    const mockApi = () =>
      new Promise((resolve) => {
        setTimeout(() => {
          resolve({
            firstName: "John",
            lastName: "Doe",
            age: 30,
            email: "john@example.com",
          });
        }, 1000);
      });

    const userData = await mockApi();
    set({ user: userData });
  },
  // 派生状态：判断用户信息是否加载完成
  isUserLoaded: (state) => Object.keys(state.user).length > 0,
  // 派生状态：用户全名
  userFullName: (state) =>
    `${state.user.firstName || ""} ${state.user.lastName || ""}`.trim(),
}));

// 订阅用户信息变化并打印日志
useUserStore.subscribe(
  (changes, state) => {
    console.log("User info changed:", { changes, newUser: state.user });
  },
  (state) => state.user // 只监听user状态变化
);

export default useUserStore;
