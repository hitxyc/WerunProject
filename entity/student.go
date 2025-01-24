package entity

type Student struct {
	StudentId string             `json:"student_id" form:"id" `
	Name      string             `json:"name" form:"name" `
	Gender    string             `json:"gender" form:"gender" `
	Class     string             `json:"class" form:"class" `
	Score     map[string]float64 `json:"score" form:"score" `
	Deleted   bool               `json:"deleted" form:"deleted" `
}

func (old *Student) Omitempty(new *Student) {
	if new.StudentId == "" {
		new.StudentId = old.StudentId
	}
	if new.Name == "" {
		new.Name = old.Name
	}
	if new.Gender == "" {
		new.Gender = old.Gender
	}
	if new.Class == "" {
		new.Class = old.Class
	}
}
