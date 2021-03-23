package dto

type ResumeReq struct {
	Uid      string `json:"uid"`
	ResumeId string `json:"resumeId"`
	Content  string `json:"content"`
	IsOpen   bool   `json:"isOpen"`
}

type ResumeRes struct {
	Uid      string `json:"uid"`
	ResumeId string `json:"resumeId"`
	Content  string `json:"content"`
	IsOpen   bool   `json:"isOpen"`
}
