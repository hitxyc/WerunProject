package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"studentManagementSystem/controller"
	"studentManagementSystem/mapper"
	"studentManagementSystem/service"
)

func Register(r *gin.Engine) {
	// 注册mapper
	sm := &mapper.StudentMapper{}
	// 注册service
	ss := &service.StudentService{StudentMapper: sm}
	// 注册controller
	sc := &controller.StudentController{StudentService: ss}

	// 注册swagger路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 注册student路由组
	student := r.Group("/student")
	{
		student.POST("/save", sc.SaveStudent)
		student.POST("/saveByFile", sc.SaveStudentByFile)
		student.GET("/get", sc.GetStudent)
		student.GET("/show", sc.ShowStudent)
		student.PUT("/update", sc.UpdateStudent)
		student.DELETE("/delete", sc.DeleteStudent)
	}
}
