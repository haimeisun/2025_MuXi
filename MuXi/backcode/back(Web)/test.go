// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// //w：http.ResponseWriter类型 ⽤来向客户端写响应
// //r：*http.Request类型 ⽤来读取客户端的请求信息
// func handler(w http.ResponseWriter, r *http.Request) {
//   fmt.Fprintf(w, "Hello,%s!", r.URL.Path[1:])//表示将请求路径进⾏处理（去掉开头的/）后和“hello，%s！”⼀同写⼊响应体

// }
//  func main() {
//   http.HandleFunc("/", handler) //注册路由处理函数
//   fmt.Println("服务启动于 http://localhost:8080")
//   http.ListenAndServe(":8080", nil)//监听本地8080端⼝，使⽤默认的路由器
// }