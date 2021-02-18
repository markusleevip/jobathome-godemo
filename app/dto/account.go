package dto

type LoginReq struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type LoginRes struct {
	Username string `json:"userName"`
	Token    string `json:"token"`
}

type LogonReq struct {
	Username   string `json:"userName"`
	Password   string `json:"password"`
	NickName   string `json:"nickName"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Avatar     string `json:"avatar"`
}


type LogonRes struct {
	Username   string `json:"userName"`
}
