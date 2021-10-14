package domain

type Count struct {
	UserId   int64     `json:"user_id"`
	Type     CountType `json:"type"`
	TargetId int64     `json:"target_id"`
}

type CountType int32

const (
	CountTypeLike CountType = iota + 1
	CountTypeFollow
	CountTypeCollect
)

