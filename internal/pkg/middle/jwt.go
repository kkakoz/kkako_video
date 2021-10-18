package middle

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"kkako_video/internal/pkg/jwtx"
)

func Verify(verifier *jwtx.JwtTokenVerifier) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		incomingContext, b := metadata.FromIncomingContext(ctx)
		if b {
			s := incomingContext.Get("token")
			claims, err := verifier.Verifier(s[0])
			if err != nil {
				return resp, err
			}
			ctx = jwtx.WithUser(ctx, claims)
			return handler(ctx, req)
		}
		return
	}
}
