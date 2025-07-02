import { createContext, useContext, useReducer, useMemo } from "react";

// 初始状态
const initialState = {
  todos: [
    { id: 1, text: "完成 React Hooks 作业", done: false },
    { id: 2, text: "学习 useContext 用法", done: true },
  ],
};

// Reducer 函数
function todoReducer(state, action) {
  switch (action.type) {
    case "ADD_TODO":
      return {
        ...state,
        todos: [
          ...state.todos,
          { id: Date.now(), text: action.text, done: false },
        ],
      };
    case "TOGGLE_TODO":
      return {
        ...state,
        todos: state.todos.map((todo) =>
          todo.id === action.id ? { ...todo, done: !todo.done } : todo
        ),
      };
    case "DELETE_TODO":
      return {
        ...state,
        todos: state.todos.filter((todo) => todo.id !== action.id),
      };
    default:
      throw new Error("未知操作类型");
  }
}

// 创建 Context
export const TodoContext = createContext();

// 自定义 Hook：获取 Context 数据
export function useTodo() {
  return useContext(TodoContext);
}

// 提供者组件
export function TodoProvider({ children }) {
  const [state, dispatch] = useReducer(todoReducer, initialState);

  // 添加待办事项
  const addTodo = (text) => {
    if (text.trim()) {
      dispatch({ type: "ADD_TODO", text });
    }
  };

  // 切换待办事项状态
  const toggleTodo = (id) => {
    dispatch({ type: "TOGGLE_TODO", id });
  };

  // 删除待办事项
  const deleteTodo = (id) => {
    dispatch({ type: "DELETE_TODO", id });
  };

  // 使用 useMemo 优化状态和方法的引用
  const value = useMemo(
    () => ({
      todos: state.todos,
      addTodo,
      toggleTodo,
      deleteTodo,
    }),
    [state.todos]
  );

  return <TodoContext.Provider value={value}>{children}</TodoContext.Provider>;
}
