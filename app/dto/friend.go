package dto

type Friend struct {
	FriendReq
}

type FriendReq struct {
	FUid string `json:"fUid"`
}
