package handle

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Res struct {
	HttpCode int
	Message  string
}

func ErrHandler(err error) Res {
	statusErr, _ := status.FromError(err)
	if statusErr.Code() == codes.Unknown {
		log.Println(fmt.Sprintf("%+v", err))
	}
	return Res{
		HttpCode: GrpcHttpCodeMap(statusErr.Code()),
		Message:  statusErr.Message(),
	}
}


func GrpcHttpCodeMap(code codes.Code) int {
	switch code {
	case codes.OK:
		return 200
	case codes.Canceled:
		return 499
	case codes.Unknown:
		return 500
	case codes.InvalidArgument:
		return 400
	case codes.DeadlineExceeded:
		return 504
	case codes.NotFound:
		return 404
	case codes.AlreadyExists:
		return 409
	case codes.PermissionDenied:
		return 403
	case codes.ResourceExhausted:
		return 429
	case codes.FailedPrecondition:
		return 400
	case codes.Aborted:
		return 409
	case codes.OutOfRange:
		return 400
	case codes.Unimplemented:
		return 501
	case codes.Internal:
		return 500
	case codes.Unavailable:
		return 503
	case codes.DataLoss:
		return 500
	case codes.Unauthenticated:
		return 401
	default:
		return 500
	}
}
