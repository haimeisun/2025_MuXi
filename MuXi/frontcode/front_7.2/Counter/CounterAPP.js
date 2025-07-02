import React from "react";
import useCounter from "./useCounter";
import "./CounterApp.css"; // 样式文件
import usecounter from "./useCounter";

function CounterApp() {
  const { increment, decrement, reset, count } = usecounter(0);

  return (
    <div className="counter-container">
      <h1 className="conter-value">{count}</h1>
      <div className="button-group">
        <button className="action-button" onClick={increment}>
          ➕ 增加
        </button>
        <button className="action-button" onClick={decrement}>
          ➖ 减少
        </button>
        <button className="action-button" onClick={reset}>
          🔄 重置
        </button>
      </div>
    </div>
  );
}

export default CounterApp;
