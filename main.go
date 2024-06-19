package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MyDB *gorm.DB //数据库变量
)

const (
	dbDsn string = "root:gh20030629@tcp(localhost:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s" //数据库连接地址
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// 访问主页
func indexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// 添加待办事项
func addTodo(c *gin.Context) {
	var td Todo
	c.BindJSON(&td)
	if err := MyDB.Create(&td).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, td)
	}
}

// 查看全部待办事项
func viewAllTodos(c *gin.Context) {
	var tds []Todo
	if err := MyDB.Find(&tds).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, tds)
	}
}

// 修改代办事项
func modifyTodo(c *gin.Context) {
	// 修改代办事项的url后面有id参数
	id, ok := c.Params.Get("id")
	// 取不到参数
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	// 取出相应id的数据
	var td Todo
	if err := MyDB.Where("id = ?", id).First(&td).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 把put请求发过来的数据绑定在td上面
	c.BindJSON(&td)
	// 存入数据库
	if err := MyDB.Save(&td).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, td)
	}
}

func deleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	if err := MyDB.Delete(&Todo{}, id).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			id: "deleted successfully",
		})
	}
}

// 数据库初始化
func initMySql() (err error) {
	MyDB, err = gorm.Open(mysql.Open(dbDsn), &gorm.Config{})
	if err != nil {
		return
	}
	db, _ := MyDB.DB()
	// 测试数据库连通性
	return db.Ping()
}

func main() {
	// 连接数据库
	err := initMySql()
	if err != nil {
		panic(err)
	}

	MyDB.AutoMigrate(&Todo{})
	r := gin.Default()
	// 加载所有的模板(html)
	r.LoadHTMLGlob("templates/*.html")
	// 设置静态文件路径，不修改的话是以当前html的路径来找，当然找不到
	r.Static("/static", "./static")
	// 待办事项路由组
	v1Group := r.Group("/v1")
	{
		// 查看所有的待办事项
		v1Group.GET("/todo", viewAllTodos)
		// 添加待办事项
		v1Group.POST("/todo", addTodo)
		// 删除某个待办事项
		v1Group.DELETE("/todo/:id", deleteTodo)
		// 修改某个待办事项,即设置是否做完
		v1Group.PUT("/todo/:id", modifyTodo)
	}
	r.GET("/", indexPage)
	r.Run(":3000")
}
