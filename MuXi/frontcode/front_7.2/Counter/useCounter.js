import { useState, useEffect, useCallback } from "react";

// 自定义 Hook：封装计数器状态和方法
function usecounter(initialvalue = 0) {
  const [count, setcount] = useState(initialvalue);

  //增加计数
  const increment = useCallback(() => {
    setcount((prev) => prev + 1);
  }, []);

  //减少计数
  const decrement = useCallback(() => {
    setcount((prev) => prev - 1);
  }, []);

  //重置计数
  const reset = useCallback(() => {
    setcount((prev = 0));
  }, []);

  //记录数值变化
  useEffect(() => {
    console.log("当前计数值为：${count}");
  }, [count]);

  return { increment, decrement, reset, count };
}
export default usecounter;
