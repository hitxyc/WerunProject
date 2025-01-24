package entity

// UndergraduateStudent 本科生
type UndergraduateStudent struct {
	StudentId string             `json:"student_id" form:"id" `
	Name      string             `json:"name" form:"name" `
	Gender    string             `json:"gender" form:"gender" `
	Class     string             `json:"class" form:"class" `
	Score     map[string]float64 `json:"score" form:"score" `
	Deleted   bool               `json:"deleted" form:"deleted" `
}

func (us *UndergraduateStudent) SetStudent() interface{} {
	var stu = &UndergraduateStudent{
		StudentId: us.StudentId,
		Name:      us.Name,
		Gender:    us.Gender,
		Class:     us.Class,
		Score:     us.Score,
		Deleted:   false,
	}
	return stu
}
func (us *UndergraduateStudent) GetStudentId() string {
	return us.StudentId
}
func (us *UndergraduateStudent) SetScore(score *map[string]float64) {
	if us.Score == nil {
		us.Score = *score // 复制学生成绩信息
	} else {
		for k, v := range *score {
			us.Score[k] = v
		}
	}
}
func (us *UndergraduateStudent) GetScore() map[string]float64 {
	return us.Score
}

// Omitempty 忽略空值
func (us *UndergraduateStudent) Omitempty(new interface{}) {
	newStudent, _ := new.(*UndergraduateStudent)
	if newStudent.StudentId == "" {
		newStudent.StudentId = us.StudentId
	}
	if newStudent.Name == "" {
		newStudent.Name = us.Name
	}
	if newStudent.Gender == "" {
		newStudent.Gender = us.Gender
	}
	if newStudent.Class == "" {
		newStudent.Class = us.Class
	}
	new = newStudent
}
func (us *UndergraduateStudent) SetDeleted() {
	us.Deleted = true
}

// GraduateStudent 研究生
type GraduateStudent struct {
	StudentId string             `json:"student_id" form:"id" `
	Name      string             `json:"name" form:"name" `
	Gender    string             `json:"gender" form:"gender" `
	Tutor     string             `json:"tutor" form:"tutor" `
	Score     map[string]float64 `json:"score" form:"score" `
	Deleted   bool               `json:"deleted" form:"deleted" `
}

func (gs *GraduateStudent) SetStudent() interface{} {
	var stu = &GraduateStudent{
		StudentId: gs.StudentId,
		Name:      gs.Name,
		Gender:    gs.Gender,
		Tutor:     gs.Tutor,
		Score:     gs.Score,
	}
	return stu
}
func (gs *GraduateStudent) GetStudentId() string {
	return gs.StudentId
}
func (gs *GraduateStudent) SetScore(score *map[string]float64) {
	if gs.Score == nil {
		gs.Score = *score // 复制学生成绩信息
	} else {
		for k, v := range *score {
			gs.Score[k] = v
		}
	}
}
func (gs *GraduateStudent) GetScore() map[string]float64 {
	return gs.Score
}
func (gs *GraduateStudent) Omitempty(new interface{}) {
	newStudent, _ := new.(*GraduateStudent)
	if newStudent.StudentId == "" {
		newStudent.StudentId = gs.StudentId
	}
	if newStudent.Name == "" {
		newStudent.Name = gs.Name
	}
	if newStudent.Gender == "" {
		newStudent.Gender = gs.Gender
	}
	if newStudent.Tutor == "" {
		newStudent.Tutor = gs.Tutor
	}
	new = newStudent
}
func (gs *GraduateStudent) SetDeleted() {
	gs.Deleted = true
}

// Student 定义学生接口
type Student interface {
	Omitempty(new interface{})
	SetStudent() interface{}
	GetStudentId() string
	SetScore(*map[string]float64)
	GetScore() map[string]float64
	SetDeleted()
}
