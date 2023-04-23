package friendship

type CreateRequest struct {
	UserID int `json:"user_id"`
	FriendID int `json:"friend_id"`
}

type UpdateRequest struct {
	UserID int `json:"user_id"`
	FriendID int `json:"friend_id"`
	Status string `json:"status"`
}
