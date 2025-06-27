package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BookHandler 处理GET /book请求，获取图书查询参数
func BookHandler(w http.ResponseWriter, r *http.Request) {
	// 从URL查询参数中获取title
	title := r.URL.Query().Get("title")
	if title == "" {
		// 如果没有传入title，返回错误响应
		http.Error(w, "缺少title参数", http.StatusBadRequest)
		return
	}
	// 构造响应内容
	response := fmt.Sprintf("您正在查询图书:《%s》", title)
	// 写入响应体
	w.Write([]byte(response))
}

// Comment struct 定义评论数据结构
type Comment struct {
	User    string `json:"user"`
	Comment string `json:"comment"`
}

// CommentHandler 处理POST /comment请求，接收评论数据
func CommentHandler(w http.ResponseWriter, r *http.Request) {
	// 定义Comment结构体用于解析请求体
	var comment Comment
	// 解析请求体中的JSON数据
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "无效的JSON格式", http.StatusBadRequest)
		return
	}
	
	// 检查必要字段是否存在
	if comment.User == "" || comment.Comment == "" {
		http.Error(w, "缺少user或comment字段", http.StatusBadRequest)
		return
	}
	
	// 设置响应头为JSON格式
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	// 构造响应数据
	response := map[string]interface{}{
		"message": "评论提交成功",
		"user":    comment.User,
		"comment": comment.Comment,
	}
	
	// 将响应数据编码为JSON并写入响应体
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "响应编码失败", http.StatusInternalServerError)
	}
}

func main() {
	// 注册路由处理函数
	http.HandleFunc("/book", BookHandler)
	http.HandleFunc("/comment", CommentHandler)
	
	// 启动服务，监听8080端口
	port := ":8080"
	fmt.Printf("服务启动于 http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(fmt.Sprintf("服务启动失败: %v", err))
	}
}