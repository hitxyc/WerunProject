package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"studentManagementSystem/entity"
	"studentManagementSystem/service"
	"studentManagementSystem/utils"
)

type StudentController struct {
	StudentService *service.StudentService
}

// 返回类处理函数
func dealWithResult(c *gin.Context, result *entity.ResultEntity, failed_code int) {
	if result.Success {
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(failed_code, result)
	}
}

// SaveStudent 保存学生信息
// @Summary 保存学生信息
// @Description 保存学生的学号,姓名,性别,班级,成绩
// @Tags student
// @Accept json
// @Produce json
// @Param student body entity.Student true "学生信息"
// @Param is_graduate query string true "是否为研究生"
// @Success 200 {object} entity.ResultEntity
// @Failure 400 {object} entity.ResultEntity
// @Router /student/save [POST]
func (sc *StudentController) SaveStudent(c *gin.Context) {
	var student entity.Student
	isGraduate := c.Query("is_graduate")
	// 判断是否为研究生
	if isGraduate == "true" {
		var gs entity.GraduateStudent
		err := c.ShouldBind(&gs)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		student = &gs
	} else {
		var us entity.UndergraduateStudent
		err := c.ShouldBind(&us)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		student = &us
	}
	result := sc.StudentService.SaveStudent(&student)
	dealWithResult(c, &result, http.StatusBadRequest)
}

// SaveStudentByFile 通过上传文件保存学生信息
// @Summary 通过上传文件保存学生信息
// @Description 通过上传文件保存学生信息
// @Tags student
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "上传文件"
// @Param is_graduate query string true "是否为研究生"
// @Success 200 {object} entity.ResultEntity
// @Failure 400 {object} entity.ResultEntity
// @Router /student/saveByFile [POST]
func (sc *StudentController) SaveStudentByFile(c *gin.Context) {
	// 获取上传文件,假设表单字段为file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ResultEntity{Message: err.Error(), Success: false})
		return
	}
	// 保存文件到本地
	dst := "./uploads/" + file.Filename
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ResultEntity{Message: err.Error(), Success: false})
		return
	}
	// 文件上传完毕, 读取文件信息
	isGraduate := c.Query("is_graduate") // 判断是否为研究生
	is_graduate, err := strconv.ParseBool(isGraduate)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ResultEntity{Message: err.Error(), Success: false})
	}
	err = utils.DealWithCSV("./uploads/"+file.Filename, is_graduate)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ResultEntity{Message: err.Error(), Success: false})
		return
	}
	c.JSON(http.StatusOK, entity.ResultEntity{Message: "Saved successfully", Success: true})
}

// GetStudent 查询学生信息
// @Summary 查询学生信息
// @Description 根据学生的学号查询学生信息
// @Tags student
// @Accept json
// @Produce json
// @Param id query string true "学生ID"
// @Success 200 {object} entity.ResultEntity
// @Failure 404 {object} entity.ResultEntity
// @Router /student/get [GET]
func (sc *StudentController) GetStudent(c *gin.Context) {
	var student entity.Student
	id := c.Query("id")
	err := c.ShouldBind(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ResultEntity{Message: err.Error(), Success: false})
		return
	}
	result := sc.StudentService.GetStudent(&id)
	dealWithResult(c, &result, http.StatusNotFound)
}

// ShowStudent 显示学生信息
// @Summary 显示学生信息
// @Description 分页显示学生信息
// @Tags student
// @Accept json
// @Produce json
// @Param page query int false "当前页数"
// @Param pageSize query int false "每页显示的记录数"
// @Success 200 {object} entity.ResultEntity
// @Router /student/show [GET]
func (sc *StudentController) ShowStudent(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	result := sc.StudentService.ShowStudent(page, pageSize)
	dealWithResult(c, &result, http.StatusNotFound)
}

// UpdateStudent 修改学生信息
// @Summary 更新学生信息
// @Description 通过学生的 `id` 更新学生的详细信息
// @Tags student
// @Accept json
// @Produce json
// @Param id query string true "学生ID"
// @Param is_graduate query string true "是否为研究生"
// @Param student body entity.Student true "学生信息"
// @Success 200 {object} entity.ResultEntity
// @Failure 404 {object} entity.ResultEntity
// @Router /student/update [PUT]
func (sc *StudentController) UpdateStudent(c *gin.Context) {
	var student entity.Student
	id := c.Query("id")
	isGraduate := c.Query("is_graduate")
	if isGraduate == "true" {
		var gs entity.GraduateStudent
		err := c.ShouldBind(&gs)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		student = &gs
	} else {
		var us entity.UndergraduateStudent
		err := c.ShouldBind(&us)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		student = &us
	}
	result := sc.StudentService.UpdateStudent(&id, &student)
	dealWithResult(c, &result, http.StatusNotFound)
}

// DeleteStudent 删除学生信息
// @Summary 删除学生信息
// @Description 通过学生的 `id` 删除学生的详细信息
// @Tags student
// @Accept json
// @Produce json
// @Param id query string true "学生ID"
// @Success 200 {object} entity.ResultEntity
// @Failure 404 {object} entity.ResultEntity
// @Router /student/delete [DELETE]
func (sc *StudentController) DeleteStudent(c *gin.Context) {
	id := c.Query("id")
	result := sc.StudentService.DeleteStudent(&id)
	dealWithResult(c, &result, http.StatusNotFound)
}
