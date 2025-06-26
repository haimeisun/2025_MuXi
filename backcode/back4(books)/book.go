package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Book 图书结构体
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  string `json:"stock"`
}

// 全局变量存储图书数据
var books = make(map[string]Book)

// AddBook 添加图书（POST请求）
func AddBook(c *gin.Context) {
	var book Book
	// 绑定请求体JSON数据
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// 检查ID是否已存在
	if _, exists := books[book.ID]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "图书ID已存在"})
		return
	}
	
	// 添加到全局map
	books[book.ID] = book
	c.JSON(http.StatusCreated, gin.H{
		"message": "图书添加成功",
		"book":    book,
	})
}

// DeleteBook 删除图书（DELETE请求）
func DeleteBook(c *gin.Context) {
	// 获取路径参数ID
	id := c.Param("id")
	
	// 删除图书
	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "图书不存在"})
		return
	}
	delete(books, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "图书删除成功",
		"deletedID": id,
	})
}

// UpdateBook 更新图书（PUT请求）
func UpdateBook(c *gin.Context) {
	// 获取路径参数ID
	id := c.Param("id")
	
	// 检查图书是否存在
	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "图书不存在"})
		return
	}
	
	// 绑定更新的图书数据
	var updateBook Book
	if err := c.BindJSON(&updateBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// 更新图书信息（保留原ID）
	updateBook.ID = id
	books[id] = updateBook
	c.JSON(http.StatusOK, gin.H{
		"message": "图书更新成功",
		"book":    updateBook,
	})
}

// SearchAllBook 查询所有图书（GET请求）
func SearchAllBook(c *gin.Context) {
	// 将map转换为切片以便序列化
	var bookList []Book
	for _, book := range books {
		bookList = append(bookList, book)
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "查询成功",
		"books":   bookList,
	})
}

func main() {
	r := gin.Default()
	
	// 注册路由，参考段落4-38的路由注册方式
	r.POST("/books", AddBook)         // 添加图书
	r.DELETE("/books/:id", DeleteBook) // 删除图书
	r.PUT("/books/:id", UpdateBook)    // 更新图书
	r.GET("/books", SearchAllBook)     // 查询所有图书
	
	// 启动服务
	r.Run(":8080")
}