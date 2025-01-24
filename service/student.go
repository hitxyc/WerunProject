package service

import (
	"fmt"
	"strings"
	"studentManagementSystem/entity"
	"studentManagementSystem/mapper"
)

type StudentService struct {
	StudentMapper *mapper.StudentMapper
}

// dealWithOmission 处理空值
func dealWithOmission(student *entity.Student) error {
	var errorMessage strings.Builder
	if student.StudentId == "" {
		errorMessage.WriteString(fmt.Sprintf("student id is empty\n"))
	}
	if student.Name == "" {
		errorMessage.WriteString(fmt.Sprintf("student name is empty\n"))
	}
	if student.Gender == "" {
		errorMessage.WriteString(fmt.Sprintf("student gender is empty\n"))
	}
	if student.Class == "" {
		errorMessage.WriteString(fmt.Sprintf("student class is empty\n"))
	}
	if student.Score == nil {
		student.Score = make(map[string]float64)
	}
	if errorMessage.Len() > 0 {
		return fmt.Errorf(errorMessage.String())
	}
	return nil
}

// SaveStudent 保存学生信息
func (ss *StudentService) SaveStudent(student *entity.Student) entity.ResultEntity {
	err := dealWithOmission(student)
	if err != nil {
		return entity.ResultEntity{Message: err.Error(), Success: false}
	}
	ss.StudentMapper.SaveStudent(student)
	return entity.ResultEntity{Message: "Student saved", Success: true, Data: *student}
}

// GetStudent 查找学生信息
func (ss *StudentService) GetStudent(studentId *string) entity.ResultEntity {
	stu := ss.StudentMapper.GetStudent(studentId)
	if stu == nil {
		return entity.ResultEntity{Message: "Student not found", Success: false}
	}
	return entity.ResultEntity{Message: "Student found successfully", Success: true, Data: *stu}
}

// ShowStudent 显示所有学生信息
func (ss *StudentService) ShowStudent(page, pageSize int) entity.ResultEntity {
	students, err := ss.StudentMapper.ShowStudent(page, pageSize)
	if err != nil {
		return entity.ResultEntity{Message: err.Error(), Success: false}
	}
	message := fmt.Sprintf("Here is  student list, page: %d, pageSize: %d :", page, pageSize)
	return entity.ResultEntity{Message: message, Success: true, Data: students}
}

// UpdateStudent 修改学生信息
func (ss *StudentService) UpdateStudent(id *string, student *entity.Student) entity.ResultEntity {
	err := ss.StudentMapper.UpdateStudent(id, student)
	if err != nil {
		return entity.ResultEntity{Message: err.Error(), Success: false}
	}
	return entity.ResultEntity{Message: "Student updated successfully", Success: true}
}

func (ss *StudentService) DeleteStudent(id *string) entity.ResultEntity {
	err := ss.StudentMapper.DeleteStudent(id)
	if err != nil {
		return entity.ResultEntity{Message: err.Error(), Success: false}
	}
	return entity.ResultEntity{Message: "Student deleted successfully", Success: true}
}
