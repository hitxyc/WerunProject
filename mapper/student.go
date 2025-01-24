package mapper

import (
	"fmt"
	"studentManagementSystem/entity"
)

type StudentMapper struct{}

// SaveStudent 保存学生信息,对未有信息进行保存,已有信息进行覆盖
func (sm *StudentMapper) SaveStudent(student *entity.Student) {
	db := GetDatabaseUndergraduate()
	student.Deleted = false
	id := student.StudentId
	db.Length++
	db.StudentMap[id] = db.Length - 1
	db.Students = append(db.Students, *student)
}

// GetStudent 查询学生信息
func (sm *StudentMapper) GetStudent(id *string) *entity.Student {
	db := GetDatabaseUndergraduate()
	if index, ok := db.StudentMap[*id]; ok {
		return &db.Students[index]
	}
	return nil
}

// ShowStudent 显示所有学生信息
func (sm *StudentMapper) ShowStudent(page, pageSize int) (*[]entity.Student, error) {
	db := GetDatabaseUndergraduate()
	var students []entity.Student
	// 将map转化为切片类型
	for _, stu := range db.Students {
		students = append(students, stu)
	}
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
	db := GetDatabaseUndergraduate()
	var oldStudent entity.Student
	var index int
	if index_, ok := db.StudentMap[*id]; !ok {
		return fmt.Errorf("student %s not found", *id)
	} else {
		index = index_
		oldStudent = db.Students[index_]
	}
	oldStudent.Omitempty(student) // 忽略空值
	// 修改学生成绩
	if student.Score == nil {
		student.Score = oldStudent.Score // 复制学生成绩信息
	} else {
		for k, v := range student.Score {
			oldStudent.Score[k] = v
		}
		student.Score = oldStudent.Score
	}
	//判断学生学号是否发生修改
	if student.StudentId == oldStudent.StudentId {
		db.Students[index] = *student // 未发生修改
	} else {
		db.StudentMap[student.StudentId] = index // 发生修改
		delete(db.StudentMap, *id)
		db.Students[index] = *student
	}
	return nil
}

// DeleteStudent 删除学生信息, 假删除
func (sm *StudentMapper) DeleteStudent(studentId *string) error {
	db := GetDatabaseUndergraduate()
	if _, ok := db.StudentMap[*studentId]; !ok {
		return fmt.Errorf("student %s not found", *studentId)
	}
	index := db.StudentMap[*studentId]
	db.Students[index].Deleted = true
	delete(db.StudentMap, *studentId)
	return nil
}
