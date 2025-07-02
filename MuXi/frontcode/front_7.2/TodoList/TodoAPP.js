import React, { useRef, useEffect } from "react";
import { TodoProvider, useTodo } from "./TodoContext";
import TodoList from "./TodoList";
import TodoInput from "./TodoInput";
import "./TodoApp.css";

function TodoApp() {
  // 组件挂载时打印日志
  useEffect(() => {
    console.log("Todo List 已加载");
  }, []);

  return (
    <TodoProvider>
      <div className="todo-app">
        <h1>简易 TodoList</h1>
        <TodoInput />
        <TodoList />
      </div>
    </TodoProvider>
  );
}

// 待办事项输入框组件
function TodoInput() {
  const inputRef = useRef(null);
  const { addTodo } = useTodo();

  const handleSubmit = (e) => {
    e.preventDefault();
    const text = inputRef.current.value.trim();
    if (text) {
      addTodo(text);
      inputRef.current.value = "";
    }
  };

  return (
    <form onSubmit={handleSubmit} className="todo-input">
      <input
        type="text"
        ref={inputRef}
        placeholder="添加待办事项..."
        required
      />
      <button type="submit">添加</button>
    </form>
  );
}

// 待办事项列表组件
function TodoList() {
  const { todos, toggleTodo, deleteTodo } = useTodo();

  // 使用 useMemo 优化列表渲染
  const renderedTodos = useMemo(() => {
    return todos.map((todo) => (
      <li key={todo.id} className={`todo-item ${todo.done ? "done" : ""}`}>
        <input
          type="checkbox"
          checked={todo.done}
          onChange={() => toggleTodo(todo.id)}
        />
        <span>{todo.text}</span>
        <button onClick={() => deleteTodo(todo.id)}>删除</button>
      </li>
    ));
  }, [todos, toggleTodo, deleteTodo]);

  return <ul className="todo-list">{renderedTodos}</ul>;
}

export default TodoApp;
