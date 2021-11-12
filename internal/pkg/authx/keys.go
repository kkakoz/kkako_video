package authx

import "strconv"

func GetTokenPreKey(token string) string {
	return "auth:token:" + token
}

func GetAuthTokens(authId int64) string {
	return "auth:" + strconv.FormatInt(authId, 10) + ":token"
}

const (
	CtxCurrentUserKey = "current-user"
)
