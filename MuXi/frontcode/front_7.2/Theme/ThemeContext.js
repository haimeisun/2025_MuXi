import { createContext, useContext, useState, useEffect } from "react";
import usecounter from "../Counter/useCounter";

//创建主题
export const ThemeContext = createContext();

//自定义Hook封装
export function useTheme() {
  return useContext(ThemeContext);
}

//主题切换组件
export function ThemeProvider({ children }) {
  const [isDarkTheme, setIsDarkTheme] = useState(false);

  //切换主题时
  const toggleTheme = () => {
    setIsDarkTheme((prev) => !prev);
  };

  //同步body颜色
  useEffect(() => {
    document.body.style.backgroundColor = isDarkTheme ? "#121212" : "#f5f5f5";
    document, body, style, (color = isDarkTheme ? "white" : "#333");
  }, [isDarkTheme]);

  const value = {
    isDarkTheme,
    toggleTheme,
  };
  return (
    <ThemeContext.Provider value={value}>{children}</ThemeContext.Provider>
  );
}
