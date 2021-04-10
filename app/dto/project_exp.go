package dto

type ProjectExp struct {
	ProjectExpReq
	ProjectExpRes
}

type ProjectExpReq struct {
	Uid         string `json:"uid"`
	ProjectId   string `json:"projectId"`
	ProjectName string `json:"projectName"`
	Content     string `json:"content"`
	IsOpen      bool   `json:"isOpen"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
}

type ProjectExpRes struct {
	Uid         string `json:"uid"`
	ProjectId   string `json:"projectId"`
	ProjectName string `json:"projectName"`
	Content     string `json:"content"`
	IsOpen      bool   `json:"isOpen"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
}
