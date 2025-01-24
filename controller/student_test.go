package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"studentManagementSystem/entity"
	"studentManagementSystem/mapper"
	"studentManagementSystem/service"
	"testing"
)

func TestAllFunction(t *testing.T) {
	// 创建一个Gin引擎实例
	r := gin.Default()
	//注册mapper
	sm := &mapper.StudentMapper{}
	//注册service
	ss := &service.StudentService{StudentMapper: sm}
	//注册controller
	sc := &StudentController{StudentService: ss}
	// 测试SaveStudent
	{
		// 注册路由
		r.POST("/student/save", sc.SaveStudent)
		// 创建学生对象实例
		stu := entity.Student{StudentId: "123", Name: "Xu", Gender: "MALE", Class: "c2", Score: map[string]float64{"math": 96}}
		// 将学生数据转为JSON格式
		jsonData, err := json.Marshal(stu)
		if err != nil {
			t.Fatalf("Failed to marshal the data, %v\n", err)
		}
		// 创建一个POST请求,并传入JSON数据
		req, err := http.NewRequest("POST", "/student/save", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatalf("Failed to create response, %v\n", err)
		}
		// 设置请求头
		req.Header.Set("Content-Type", "application/json")
		// 创建一个http响应记录器
		w := httptest.NewRecorder()
		// 执行请求
		r.ServeHTTP(w, req)
		// 检查响应内容
		var result entity.ResultEntity
		err = json.Unmarshal(w.Body.Bytes(), &result)
		if err != nil {
			t.Fatalf("Failed to unmarshal result, %v\n", err)
		}
		// 检查状态码
		if w.Code != http.StatusOK {
			t.Fatalf("Failed to save the data, %v\n", result.Message)
		}
		// 检查保存数据
		var data = result.Data
		if stu_, ok := data.(*entity.Student); ok {
			// 检查id
			if stu_.StudentId != stu_.StudentId {
				t.Fatalf("Failed to save the StudentId, %v\n", stu_)
			}
			// 检查name
			if stu_.Name != stu.Name {
				t.Fatalf("Failed to save the Name, %v\n", stu_.Name)
			}
			// 检查gender
			if stu_.Gender != stu.Gender {
				t.Fatalf("Failed to save the Gender, %v\n", stu_)
			}
			// 检查class
			if stu_.Class != stu.Class {
				t.Fatalf("Failed to save the Class, %v\n", stu_)
			}
			// 检查score
			for k, _ := range stu.Score {
				if stu_.Score[k] != stu.Score[k] {
					t.Fatalf("Failed to save the Score, %v\n", stu_)
				}
			}
		}
		// 经过测试save功能passed
		t.Logf("Successfully saved the data, %v\n", result.Data)
	}
	// 测试GetStudent
	{
		// 注册路由
		r.GET("/student/get", sc.GetStudent)
		// 创建查询参数
		id := "123"
		queryParams := "?id="
		// 创建http请求
		req, err := http.NewRequest("GET", "/student/get"+queryParams+id, nil)
		if err != nil {
			t.Fatalf("Failed to create request, %v\n", err)
		}
		// 创建http响应记录器
		w := httptest.NewRecorder()
		// 启动http服务
		r.ServeHTTP(w, req)
		// 断言返回的响应体
		var result entity.ResultEntity
		err = json.Unmarshal(w.Body.Bytes(), &result)
		if err != nil {
			t.Fatalf("Failed to unmarshal result, %v\n", err)
		}
		// 断言是否成功
		assert.True(t, result.Success)
		// 断言返回student的信息
		var student = result.Data.(map[string]interface{})
		var score = student["score"].(map[string]interface{})
		assert.Equal(t, "123", student["student_id"])
		assert.Equal(t, "Xu", student["name"])
		assert.Equal(t, "MALE", student["gender"])
		assert.Equal(t, "c2", student["class"])
		assert.Equal(t, float64(96), score["math"])
		// 经过测试get功能passed
		t.Logf("Successfully got the data, %v\n", result.Data)
	}
	// 测试ShowStudent
	{
		// 注册路由
		r.GET("/student/show", sc.ShowStudent)
		// 创建http请求
		req, err := http.NewRequest("GET", "/student/show", nil)
		if err != nil {
			t.Fatalf("Failed to create request, %v\n", err)
		}
		// 创建http响应记录器
		w := httptest.NewRecorder()
		// 启动http服务
		r.ServeHTTP(w, req)
		// 断言返回的响应体
		var result entity.ResultEntity
		err = json.Unmarshal(w.Body.Bytes(), &result)
		if err != nil {
			t.Fatalf("Failed to unmarshal result, %v\n", err)
		}
		assert.True(t, result.Success)
		// 提取result.Data的信息
		var students = []entity.Student{}
		jsonData, err := json.Marshal(result.Data)
		if err != nil {
			t.Fatalf("Failed to marshal the data, %v\n", err)
		}
		err = json.Unmarshal(jsonData, &students)
		if err != nil {
			t.Fatalf("Failed to unmarshal result, %v\n", err)
		}
		// 断言students的信息
		for _, stu := range students {
			if stu.StudentId == "123" {
				assert.Equal(t, "Xu", stu.Name)
				assert.Equal(t, "MALE", stu.Gender)
				assert.Equal(t, "c2", stu.Class)
				assert.Equal(t, float64(96), stu.Score["math"])
			}
		}
		t.Logf("Successfully got the data, %v\n", result.Data)
	}
	// 测试修改学生信息
	{
		// 注册路由
		r.PUT("/student/update", sc.UpdateStudent)
		// 设定修改信息
		id := "123"
		stu := entity.Student{StudentId: "0123", Name: "XuYeCheng", Class: "c3", Score: map[string]float64{"math": 98, "Chinese": 93}}
		jsonData, err := json.Marshal(stu)
		if err != nil {
			t.Fatalf("Failed to marshal the data, %v\n", err)
		}
		// 创建一个put请求，并传入json数据
		req, err := http.NewRequest("PUT", "/student/update?id="+id, bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatalf("Failed to create response, %v\n", err)
		}
		// 设置请求头
		req.Header.Set("Content-Type", "application/json")
		// 创建一个响应记录器
		w := httptest.NewRecorder()
		// 启动服务
		r.ServeHTTP(w, req)
		// 断言返回体
		var result entity.ResultEntity
		err = json.Unmarshal(w.Body.Bytes(), &result)
		if err != nil {
			t.Fatalf("Failed to unmarshal result, %v\n", err)
		}
		assert.True(t, result.Success)
		t.Logf("Successfully updated the data\n")
	}
	// 测试delete
	{
		// 注册路由
		r.DELETE("/student/delete", sc.DeleteStudent)
		// 设定要删除的学生
		id := "0123"
		// 创建一个GET请求
		req, err := http.NewRequest("DELETE", "/student/delete?id="+id, nil)
		if err != nil {
			t.Fatalf("Failed to create request, %v\n", err)
		}
		// 创建一个响应记录器
		w := httptest.NewRecorder()
		// 启动服务
		r.ServeHTTP(w, req)
		// 断言返回体
		var result entity.ResultEntity
		err = json.Unmarshal(w.Body.Bytes(), &result)
		if err != nil {
			t.Fatalf("Failed to unmarshal result, %v\n", err)
		}
		assert.True(t, result.Success)
		t.Logf("Successfully deleted the data\n")
	}
}
