package mapper

import (
	"fmt"
	"studentManagementSystem/entity"
	"testing"
)

// TestSaveStudent 测试保存和查询
func TestSaveStudent(t *testing.T) {
	createTestSaveAndGetStudentCase(t, &entity.Student{StudentId: "123", Name: "Xu", Gender: "MALE", Class: "c2", Score: map[string]float64{"math": 96}})
	createTestSaveAndGetStudentCase(t, &entity.Student{StudentId: "401", Name: "Yang", Gender: "FEMALE", Class: "c3", Score: map[string]float64{"math": 96}})
}
func createTestSaveAndGetStudentCase(t *testing.T, stu *entity.Student) {
	sm := StudentMapper{}
	sm.SaveStudent(stu)
	if st := sm.GetStudent(&stu.StudentId); st == nil {
		t.Fatalf("Student not saved or student can't get")
	} else {
		fmt.Printf("The student is: %+v\n", st)
	}
}

func setup() {
	sm := StudentMapper{}
	sm.SaveStudent(&entity.Student{StudentId: "123", Name: "Xu", Gender: "MALE", Class: "c2", Score: map[string]float64{"Math": 96}})
	sm.SaveStudent(&entity.Student{StudentId: "401", Name: "Yang", Gender: "FEMALE", Class: "c3", Score: map[string]float64{"Math": 96}})
}

// TestDeleteStudent测试删除学生信息
func TestDeleteStudent(t *testing.T) {
	sm := StudentMapper{}
	setup()
	var stuId = "123"
	sm.DeleteStudent(&stuId)
	if st := sm.GetStudent(&stuId); st != nil {
		t.Fatalf("Student delete failed")
	} else {
		fmt.Printf("Student delete passed, %+v\n", GetDB())
	}
}

// TestUpdateStudentInfo测试更新学生信息
func TestUpdateStudentInfo(t *testing.T) {
	createTestUpdateStudentInfoCase(t, "401", &entity.Student{Name: "YangMingXi", Score: map[string]float64{"Chinese": 94, "Math": 98}})
}
func createTestUpdateStudentInfoCase(t *testing.T, id string, stu *entity.Student) {
	sm := StudentMapper{}
	setup()
	err := sm.UpdateStudent(&id, stu)
	if err != nil {
		t.Fatalf("Student update failed, %v\n", err)
	}
	if st := sm.GetStudent(&id); st.Name != "YangMingXi" {
		t.Fatalf("StudentInfo update failed")
	} else if stu.Score["Math"] != 98 || stu.Score["Chinese"] != 94 {
		t.Fatalf("StudentInfo update failed")
	} else {
		fmt.Printf("StudentInfo update passed, %+v\n", GetDB())
	}
}
