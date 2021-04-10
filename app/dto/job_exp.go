package dto

type JobExpReq struct {
	Uid         string `json:"uid"`
	JobId       string `json:"jobId"`
	CompanyName string `json:"companyName"`
	Content     string `json:"content"`
	IsOpen      bool   `json:"isOpen"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
}

type JobExpRes struct {
	Uid         string `json:"uid"`
	JobId       string `json:"jobId"`
	CompanyName string `json:"companyName"`
	Content     string `json:"content"`
	IsOpen      bool   `json:"isOpen"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
}
