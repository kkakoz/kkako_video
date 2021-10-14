package domain

import "context"

type User struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Avatar      string  `json:"avatar"`
	Brief       string  `json:"brief"`
	FollowCount int64   `json:"follow_count"`
	FansCount   int64   `json:"fans_count"`
	LikeCount   int64   `json:"like_count"`
	State       int32   `json:"state"`
	Follows     []int64 `json:"follows" gorm:"-"`
	Fans        []int64 `json:"fans" gorm:"-"`
	Likes       []int64 `json:"likes" gorm:"-"`
}

type News struct {
	Id      int64 `json:"id"`
	UserId  int64 `json:"user_id"`
	Content int64 `json:"content"`
	Type    int32 `json:"type"`
}

type IUserLogic interface {
	AddUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id int64) (*User, error)
	GetUsers(ctx context.Context, ids []int64) ([]*User, error)
}

type IUserRepo interface {
	AddUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id int64) (*User, error)
	GetUserList(ctx context.Context, ids []int64) ([]*User, error)
}
