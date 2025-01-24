package mapper

import "studentManagementSystem/entity"

type Database struct {
	StudentMap map[string]int
	Students   []entity.Student
	Length     int
}

// 运用哈希表储存学生信息
var db = Database{StudentMap: make(map[string]int), Students: []entity.Student{}, Length: 0}

func GetDatabase() *Database {
	return &db
}
