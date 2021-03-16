package models

/**
	项目经历
 */
type ProjectExperience struct {
	Common
	Experience
}

func (ProjectExperience) TableName() string {
return Prefix+"project_experience"
}

