package dto

type LoginReq struct {
	Username    string       `json:"userName"`
	Password    string       `json:"password"`
}

type LoginRes struct {
	Username    string       `json:"userName"`
	Token		string		 `json:"token"`
}

type Logon struct {
	Username    string       `json:"userName"`
	Password    string       `json:"password"`
	rePassword    string       `json:"rePassword"`
	NickName    string       `json:"nickName"`
	Phone    string       	 `json:"phone"`
	Email    string       	 `json:"email"`
	Avatar   string       	 `json:"avatar"`
}
