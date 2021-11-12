package errx

import (
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Err struct {
	Code    uint32
	Message string
	Detail  []interface{}
}

func (e Err) Error() string {
	return e.Message
}



func ServerErrHandler(err error) error {
	if err == nil {
		return nil
	}
	e := &Err{}
	if errors.As(err, e) {
		return status.Error(codes.Code(e.Code), e.Message)
	}
	//status.FromError()
	return status.Error(codes.Unknown, fmt.Sprintf("%v", err))
}
