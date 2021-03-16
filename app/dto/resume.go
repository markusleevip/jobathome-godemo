package dto


type ResumeReq struct {
	Uid string `json:"uid"`
	ResumeId string `json:"resumeId"`
	Content string `json:"content"`
}


type ResumeRes struct {
	Uid string `json:"uid"`
	ResumeId string `json:"resumeId"`
	Content string `json:"content"`
}



