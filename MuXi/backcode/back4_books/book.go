package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Book 图书结构体
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  string `json:"stock"`
}

// response通用响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 全局变量存储图书数据
var books = make(map[string]Book)

//AddBook 添加图书
//@Summary 添加图书
//@Description 传入图书信息新增图书
//@Tags 图书
//@Accept json
//@Produce json
//@Param book body Book true "图书信息"
//@Success 200 {object} Response{data=Book} "图书信息"
//@Failuer 400 {object} Response "参数解析失败"
//@Failure 409 {object} Response "图书ID已存在"
//@Ruoter /books/add [post]
/* AddBook 添加图书（POST请求）*/
func AddBook(c *gin.Context) {
	var book Book
	// 绑定请求体JSON数据
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数解析失败", Data: err.Error()})
		return
	}

	// 检查ID是否已存在
	if _, exists := books[book.ID]; exists {
		c.JSON(http.StatusConflict, Response{Code: 409, Message: "图书ID已存在"})
		return
	}

	// 添加到全局map
	books[book.ID] = book
	c.JSON(http.StatusCreated, Response{Code: 200, Message: "图书添加成功", Data: book})
}

//DeleteBook 删除图书
//@Summary 删除图书
//@Description 根据图书ID删除图书
//@Tags 图书
//@Accept json
//@Produce json
//@Param id path string true "图书删除成功"
//@Success 200 {object} Response "图书删除成功"
//@Failure 400 {object} Response "图书不存在"
//@Router /books/delete/{id} [delete]
/* DeleteBook 删除图书（DELETE请求）*/
func DeleteBook(c *gin.Context) {
	// 获取路径参数ID
	id := c.Param("id")

	// 删除图书
	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, Response{Code: 400, Message: "图书不存在"})
		return
	}
	delete(books, id)
	c.JSON(http.StatusOK, Response{Code: 200, Message: "图书删除成功"})
}

//UpdateBook 更新图书
//@Summary 更新图书
//@Description 根据ID更新图书信息
//@Tags 图书
//@Accept json
//@Produce json
//@Param id path string true "图书ID"
//@Param book body Book true "更新后图书信息"
//@Success 200 {object} Response{data=Book} "图书更新成功"
//@Failure 400 {object} Response "参数绑定失败"
//@Failure 404 {object} Response "路径ID与请求ID不一致"
//@Failure 409 {object} Response "图书不存在"
//@Router /books/update/{id} [put]
/* UpdateBook 更新图书（PUT请求）*/
func UpdateBook(c *gin.Context) {
	// 获取路径参数ID
	id := c.Param("id")

	// 检查图书是否存在
	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, Response{Code: 409, Message: "图书不存在"})
		return
	}

	// 绑定更新的图书数据
	var updateBook Book
	if err := c.BindJSON(&updateBook); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数绑定失败", Data: err.Error()})
		return
	}

	//更新的图书ID与目标图书ID对不上
	if id != updateBook.ID {
		c.JSON(http.StatusBadGateway, Response{Code: 404, Message: "路径ID与请求ID不一致"})
		return
	}

	// 更新图书信息（保留原ID）
	updateBook.ID = id
	books[id] = updateBook
	c.JSON(http.StatusOK, Response{Code: 200, Message: "图书更新成功", Data: updateBook})
}

//SearchAllBook 查询图书
//@Summary 查询图书
//@Description 根据图书ID查询图书
//@Tags 图书
//@Produce json
//@Success 200 {object} Response{data=[]Book} "图书列表"
//@Router /books/search [get]
/* SearchAllBook 查询所有图书（GET请求）*/
func SearchAllBook(c *gin.Context) {
	// 将map转换为切片以便序列化
	var bookList []Book
	for _, book := range books {
		bookList = append(bookList, book)
	}

	c.JSON(http.StatusOK, Response{Code: 200, Message: "图书获取成功", Data: bookList})
}

//@title 图书管理系统
//@version 1.0
//@Description 实现对图书的增删查改功能的图书管理系统
//@host localhost：8080
//@BasePath /

// SecretKey 是用于签名和验证 JWT 的密钥
// 注意：在实际生产环境中，绝对不能将密钥硬编码在源代码中
// 应通过环境变量、配置文件或安全密钥管理服务来获取
const SecretKey = "1234567"

// Claims 定义了 JWT 中存储的信息结构
// 嵌入了 jwt.StandardClaims 以包含标准字段
type Claims struct {
	UserId int `json:"user_id"` // 用户ID，用于标识用户身份
	jwt.StandardClaims
}

// GenerateToken 生成 JWT Token
// 参数 id 是用户ID，用于存储在 Token 中
// 返回生成的 Token 字符串和可能的错误
func GenerateToken(id int) (string, error) {
	// 创建自定义声明，包含用户ID和标准声明
	claims := Claims{
		UserId: id, // 使用传入的用户ID
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 1, 0).Unix(), // 过期时间设置为1个月后
			Issuer:    "your-application",                 // 签发者标识
		},
	}

	// 使用 HS256 算法和自定义声明创建 Token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥对 Token 进行签名，生成最终的 Token 字符串
	return token.SignedString([]byte(SecretKey))
}

// AuthMiddleware 创建一个 Gin 中间件，用于验证 HTTP 请求中的 JWT
// 验证通过后将用户ID存入上下文，便于后续处理使用
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Authorization 字段
		tokenString := c.GetHeader("Authorization")

		// 检查 Token 是否存在
		if tokenString == "" {
			c.JSON(401, gin.H{"message": "未认证"})
			c.Abort() // 终止请求处理链
			return
		}

		// 解析和验证 Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证签名方法是否为预期的 HS256
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// 返回签名密钥以验证签名
			return []byte(SecretKey), nil
		})

		// 处理解析错误
		if err != nil {
			c.JSON(401, gin.H{"message": "无效的 Token", "error": err.Error()})
			c.Abort()
			return
		}

		// 检查 Token 是否有效
		if !token.Valid {
			c.JSON(401, gin.H{"message": "无效或已过期的 Token"})
			c.Abort()
			return
		}

		// 将 Token 声明转换为 MapClaims 类型
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(401, gin.H{"message": "未认证"})
			c.Abort()
			return
		}

		// 从声明中获取 user_id 并转换为 int 类型
		// 注意：JWT 中数字默认解析为 float64 类型
		userId, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(401, gin.H{"message": "无法获取用户ID"})
			c.Abort()
			return
		}

		// 打印用户ID（调试用）并将其存入 Gin 上下文
		fmt.Println("Authenticated user ID:", userId)
		c.Set("user_id", int(userId)) // 将用户ID存入上下文，后续处理函数可通过 c.MustGet("user_id") 获取

		// 继续处理后续的处理函数
		c.Next()
	}
}

func main() {
	// 创建默认的 Gin 引擎，包含日志和恢复中间件
	r := gin.Default()

	// 应用跨域资源共享中间件
	r.Use(AuthMiddleware())

	// 定义登录接口，用于生成 Token
	// r.POST("/login", Login)    ***由于没有登陆系统，我就注释掉了***

	// 定义需要认证的图书管理路由组
	book := r.Group("/book")
	{
		// 公开接口：搜索所有图书，无需认证
		book.GET("/search", SearchAllBook)

		// 应用认证中间件，后续的接口需要有效 Token 才能访问
		book.Use(AuthMiddleware())

		// 受保护的接口
		book.POST("/add", AddBook)             // 添加图书
		book.DELETE("/delete/:id", DeleteBook) // 删除图书
		book.PUT("/update/:id", UpdateBook)    // 更新图书信息
	}

	// 启动服务器，默认监听 8080 端口
	r.Run()
}
