// 控制器，路由的handler
package controller

import (
	"mytodolist/db"
	"mytodolist/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 访问主页
func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// 添加待办事项
func AddTodo(c *gin.Context) {
	var td models.Todo
	c.BindJSON(&td)
	if err := db.MyDB.Create(&td).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, td)
	}
}

// 查看全部待办事项
func ViewAllTodos(c *gin.Context) {
	var tds []models.Todo
	if err := db.MyDB.Find(&tds).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, tds)
	}
}

// 修改代办事项
func ModifyTodo(c *gin.Context) {
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
	var td models.Todo
	if err := db.MyDB.Where("id = ?", id).First(&td).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 把put请求发过来的数据绑定在td上面
	c.BindJSON(&td)
	// 存入数据库
	if err := db.MyDB.Save(&td).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, td)
	}
}

// 删除待办事项
func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	if err := db.MyDB.Delete(&models.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			id: "deleted successfully",
		})
	}
}
