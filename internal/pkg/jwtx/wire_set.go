package jwtx

import (
	"github.com/google/wire"
)

var AuthSet = wire.NewSet(NewJwtTokenVerifier, NewJwtTokenGen)