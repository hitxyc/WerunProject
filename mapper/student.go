package mapper

import (
	"fmt"
	"reflect"
	"studentManagementSystem/entity"
)

type StudentMapper struct{}

func callMethod(student *reflect.Value, methodName string, param []reflect.Value) ([]reflect.Value, error) {
	studentMethod := student.MethodByName(methodName)
	if !studentMethod.IsValid() {
		return nil, fmt.Errorf("student method %s is not valid", methodName)
	}
	result := studentMethod.Call(param)
	return result, nil
}

// SaveStudent 保存学生信息,对未有信息进行保存,已有信息进行覆盖
func (sm *StudentMapper) SaveStudent(student *entity.Student) error {
	db := GetDatabase()
	// 获得学生结构体的值
	studentValue := reflect.ValueOf(*student)
	// 获得并调用 SetStudent 方法
	setStudentResult, err := callMethod(&studentValue, "SetStudent", nil)
	if err != nil {
		return err
	}
	// 获得 SetStudent 的返回对象
	var studentResult entity.Student
	switch v := setStudentResult[0].Interface().(type) {
	case *entity.UndergraduateStudent:
		studentResult = v
	case *entity.GraduateStudent:
		studentResult = v
	default:
		return fmt.Errorf("unsupported student type")
	}
	// 将学生信息存入数据库中
	db.StudentMap[studentResult.GetStudentId()] = db.Length
	db.Students = append(db.Students, studentResult)
	db.Length++
	return nil
}

// GetStudent 查询学生信息
func (sm *StudentMapper) GetStudent(id *string) *entity.Student {
	db := GetDatabase()
	if index, ok := db.StudentMap[*id]; ok {
		return &db.Students[index]
	}
	return nil
}

// ShowStudent 显示所有学生信息
func (sm *StudentMapper) ShowStudent(page, pageSize int) (*[]entity.Student, error) {
	db := GetDatabase()
	var students []entity.Student
	students = db.Students
	// 进行分页展示
	if (page-1)*pageSize > len(students) {
		return nil, fmt.Errorf("page number is out of range")
	}
	var stu []entity.Student
	for i := (page - 1) * pageSize; i < len(students) && i < page*pageSize; i++ {
		stu = append(stu, students[i])
	}
	return &stu, nil
}

// UpdateStudent 更新学生信息
func (sm *StudentMapper) UpdateStudent(id *string, student *entity.Student) error {
	db := GetDatabase()
	var oldStudent entity.Student
	var index int
	if index_, ok := db.StudentMap[*id]; !ok {
		return fmt.Errorf("student %s not found", *id)
	} else {
		index = index_
		oldStudent = db.Students[index]
	}
	// 获得学生结构体的值
	studentValue := reflect.ValueOf(*student)
	oldStudentValue := reflect.ValueOf(oldStudent)
	// 忽略空值,修改学生信息(不含成绩), 获得并调用 Omitempty 方法
	_, err := callMethod(&oldStudentValue, "Omitempty", []reflect.Value{reflect.ValueOf(*student)})
	if err != nil {
		return err
	}
	// 获得学生需修改的成绩
	getScoreResult, err := callMethod(&studentValue, "GetScore", nil)
	if err != nil {
		return err
	}
	score, _ := getScoreResult[0].Interface().(map[string]float64)
	// 修改学生成绩
	_, err = callMethod(&oldStudentValue, "SetScore", []reflect.Value{reflect.ValueOf(&score)})
	if err != nil {
		return err
	}
	// 获得学生学号
	getStudentIdResult, err := callMethod(&studentValue, "GetStudentId", nil)
	if err != nil {
		return err
	}
	studentId, _ := getStudentIdResult[0].Interface().(string)
	// 判断学生学号是否发生修改
	if studentId == *id {
		db.Students[index] = *student // 未发生修改
	} else {
		db.StudentMap[studentId] = index // 发生修改
		delete(db.StudentMap, *id)
		db.Students[index] = *student
	}
	return nil
}

// DeleteStudent 删除学生信息, 假删除
func (sm *StudentMapper) DeleteStudent(studentId *string) error {
	db := GetDatabase()
	if _, ok := db.StudentMap[*studentId]; !ok {
		return fmt.Errorf("student %s not found", *studentId)
	}
	index := db.StudentMap[*studentId]
	student := db.Students[index]
	studentValue := reflect.ValueOf(student)
	// 获得并调用 SetDeleted 方法
	_, err := callMethod(&studentValue, "SetDeleted", nil)
	if err != nil {
		return err
	}
	delete(db.StudentMap, *studentId)
	return nil
}
