package models

/**
	工作经历
 */
type JobExperience struct {
	Common
	Experience
}

func (JobExperience) TableName() string {
	return Prefix+"job_experience"
}
