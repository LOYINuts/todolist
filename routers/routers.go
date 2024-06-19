package routers

import (
	"mytodolist/controller"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	// 加载所有的模板(html)
	r.LoadHTMLGlob("templates/*.html")
	// 设置静态文件路径，不修改的话是以当前html的路径来找，当然找不到
	r.Static("/static", "./static")
	// 待办事项路由组
	v1Group := r.Group("/v1")
	{
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.ViewAllTodos)
		// 添加待办事项
		v1Group.POST("/todo", controller.AddTodo)
		// 删除某个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
		// 修改某个待办事项,即设置是否做完
		v1Group.PUT("/todo/:id", controller.ModifyTodo)
	}
	r.GET("/", controller.IndexPage)
	return r
}
