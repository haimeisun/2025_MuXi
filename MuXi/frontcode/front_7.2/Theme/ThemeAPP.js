import React from "react";
import { ThemeProvider, useTheme } from "./ThemeContext";
import "./ThemeAPP.css";

function ThemeCard() {
  const { isDarkTheme } = useTheme();
  return (
    <div className={"theme-card ${isDarkTheme?'dark':'light'}"}>
      <h2>当前主题</h2>
      <p>{isDarkTheme ? "深色模式" : "浅色模式"}</p>
    </div>
  );
}

function ThemeButton() {
  const { toggleTheme } = useTheme();
  return (
    <button className="theme-button" onClick={toggleTheme}>
      🎨 切换主题
    </button>
  );
}

function ThemeAPP() {
  return (
    <ThemeProvider>
      <div className="theme-app">
        <h1>主题切换应用</h1>
        <ThemeCard />
        <ThemeButton />
      </div>
    </ThemeProvider>
  );
}
