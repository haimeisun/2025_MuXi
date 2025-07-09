import { create } from "zustand";
import { persist } from "zustand/middleware";

const useCounterStore = create(
  persist(
    (set) => ({
      count: 0,
      // 计数器操作方法
      increment: () => set((state) => ({ count: state.count + 1 })),
      decrement: () => set((state) => ({ count: state.count - 1 })),
      // 派生状态：判断计数是否为正数
      isPositive: (state) => state.count > 0,
    }),
    {
      name: "counter-storage", // 本地存储键名
      storage: localStorage, // 持久化到localStorage
    }
  )
);

export default useCounterStore;
