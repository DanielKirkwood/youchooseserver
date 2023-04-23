package friendship

import "time"

type Schema struct {
	ID int
	UserID int
	FriendID int
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time
}
