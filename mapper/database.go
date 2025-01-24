package mapper

import "studentManagementSystem/entity"

// Database 本科生数据库
type UndergraduateDatabase struct {
	StudentMap map[string]int
	Students   []entity.Student
	Length     int
}

// 运用哈希表储存学生信息
var dbUndergraduate = UndergraduateDatabase{StudentMap: make(map[string]int), Students: []entity.Student{}, Length: 0}

func GetDatabaseUndergraduate() *UndergraduateDatabase {
	return &dbUndergraduate
}
