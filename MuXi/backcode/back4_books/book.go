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

// response通用响应结构体
type Response struct{
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
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
	if id != updateBook.ID{
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
func main() {
	r := gin.Default()
	
	// 注册路由
	r.POST("/books", AddBook)         // 添加图书
	r.DELETE("/books/:id", DeleteBook) // 删除图书
	r.PUT("/books/:id", UpdateBook)    // 更新图书
	r.GET("/books", SearchAllBook)     // 查询所有图书
	
	// 启动服务
	r.Run(":8080")
}