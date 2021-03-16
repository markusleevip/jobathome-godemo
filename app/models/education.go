package models

/**
教育经验
*/
type Education struct {
	Common
	Experience
	SchoolName 	string `json:"schoolName" gorm:"size:128;"`
	Speciality 	string `json:"speciality" gorm:"size:128;comment:专业"`
	Certificate	string `json:"certificate" gorm:"size:128;comment:证书"`
}

func (Education) TableName() string {
	return Prefix+"education"
}

