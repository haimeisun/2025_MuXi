package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 模拟发送GET请求查询图书
func sendGetRequest(title string) {
	// 构造请求URL
	url := fmt.Sprintf("http://localhost:8080/book?title=%s", title)
	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Sprintf("GET请求失败: %v", err))
	}
	defer resp.Body.Close()
	
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("读取响应失败: %v", err))
	}
	
	fmt.Printf("GET请求响应: %s\n", string(body))
}

// 模拟发送POST请求提交评论
func sendPostRequest(user, comment string) {
	// 构造请求数据
	requestData := map[string]string{
		"user":    user,
		"comment": comment,
	}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(fmt.Sprintf("JSON编码失败: %v", err))
	}
	
	// 发送POST请求
	url := "http://localhost:8080/comment"
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		panic(fmt.Sprintf("POST请求失败: %v", err))
	}
	defer resp.Body.Close()
	
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("读取响应失败: %v", err))
	}
	
	fmt.Printf("POST请求响应: %s\n", string(body))
}

func main() {
	fmt.Println("===== 开始测试客户端请求 =====")
	
	// 测试GET请求
	fmt.Println("\n--- 测试GET /book请求 ---")
	sendGetRequest("三体")
	sendGetRequest("百年孤独")
	
	// 测试POST请求
	fmt.Println("\n--- 测试POST /comment请求 ---")
	sendPostRequest("小李", "这本书真棒!")
	sendPostRequest("小王", "内容很精彩，推荐阅读!")
	
	fmt.Println("\n===== 请求测试完成 =====")
}