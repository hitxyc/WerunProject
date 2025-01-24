package mapper

import (
	"fmt"
	"reflect"
	"studentManagementSystem/entity"
	"testing"
)

// TestSaveStudent 测试保存和查询
func TestSaveStudent(t *testing.T) {
	stu1 := &entity.UndergraduateStudent{StudentId: "123", Name: "Xu", Gender: "MALE", Class: "c2", Score: map[string]float64{"math": 96}}
	stu2 := &entity.GraduateStudent{StudentId: "401", Name: "Yang", Gender: "FEMALE", Tutor: "Xu", Score: map[string]float64{"math": 96}}
	createTestSaveAndGetStudentCase(t, stu1)
	createTestSaveAndGetStudentCase(t, stu2)
}
func createTestSaveAndGetStudentCase(t *testing.T, stu entity.Student) {
	sm := StudentMapper{}
	sm.SaveStudent(&stu)
	var graduateStudent *entity.GraduateStudent
	var undergraduateStudent *entity.UndergraduateStudent
	switch studentType := stu.(type) {
	case *entity.UndergraduateStudent:
		undergraduateStudent = studentType
		if st := sm.GetStudent(&undergraduateStudent.StudentId); st == nil {
			t.Fatalf("Student not saved or student can't get")
		} else {
			fmt.Printf("The student is: %+v\n", *st)
		}
	case *entity.GraduateStudent:
		graduateStudent = studentType
		if st := sm.GetStudent(&graduateStudent.StudentId); st == nil {
			t.Fatalf("Student not saved or student can't get")
		} else {
			fmt.Printf("The student is: %+v\n", *st)
		}
	}
}

func setup() error {
	sm := StudentMapper{}
	var stu1 entity.Student = &entity.UndergraduateStudent{StudentId: "123", Name: "Xu", Gender: "MALE", Class: "c2", Score: map[string]float64{"Math": 96}}
	var stu2 entity.Student = &entity.GraduateStudent{StudentId: "401", Name: "Yang", Gender: "FEMALE", Tutor: "Xu", Score: map[string]float64{"math": 96}}
	err := sm.SaveStudent(&stu1)
	if err != nil {
		return err
	}
	err = sm.SaveStudent(&stu2)
	if err != nil {
		return err
	}
	return nil
}

// TestDeleteStudent测试删除学生信息
func TestDeleteStudent(t *testing.T) {
	sm := StudentMapper{}
	err := setup()
	if err != nil {
		t.Fatal(err)
	}
	var stuId = "123"
	err = sm.DeleteStudent(&stuId)
	if err != nil {
		t.Fatal(err)
	}
	if st := sm.GetStudent(&stuId); st != nil {
		t.Fatalf("Student delete failed")
	} else {
		fmt.Printf("Student delete passed, %+v\n", GetDatabase())
	}
}

// TestUpdateStudentInfo测试更新学生信息
func TestUpdateStudentInfo(t *testing.T) {
	var stu entity.Student = &entity.GraduateStudent{Name: "YangMingXi", Score: map[string]float64{"Chinese": 94, "math": 98}}
	createTestUpdateStudentInfoCase(t, "401", &stu)
}
func createTestUpdateStudentInfoCase(t *testing.T, id string, stu *entity.Student) {
	sm := StudentMapper{}
	err := setup()
	if err != nil {
		t.Fatal(err)
	}
	err = sm.UpdateStudent(&id, stu)
	if err != nil {
		t.Fatalf("Student update failed, %v\n", err)
	}
	db := GetDatabase()
	index := db.StudentMap[id]
	student := GetDatabase().Students[index]
	studentValue := reflect.ValueOf(student)
	var gs *entity.GraduateStudent
	var us *entity.UndergraduateStudent
	var ok bool
	if gs, ok = studentValue.Interface().(*entity.GraduateStudent); ok {
		fmt.Printf("The student is: %+v\n", *gs)
	}
	if us, ok = studentValue.Interface().(*entity.UndergraduateStudent); ok {
		fmt.Printf("The student is: %+v\n", *us)
	}
}
