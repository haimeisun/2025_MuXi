import { useCounterStore } from "./useCounterStore";

const Counter = () => {
  // 通过selector获取状态和方法
  const count = useCounterStore((state) => state.count);
  const increment = useCounterStore((state) => state.increment);
  const decrement = useCounterStore((state) => state.decrement);
  const isPositive = useCounterStore((state) => state.isPositive);

  return (
    <div>
      <h2>Counter: {count}</h2>
      <button onClick={increment}>+</button>
      <button onClick={decrement}>-</button>
      <p>{isPositive ? "Count is positive" : "Count is not positive"}</p>
    </div>
  );
};

export default Counter;
