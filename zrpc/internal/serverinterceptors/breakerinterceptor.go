package serverinterceptors

import (
	"context"
	"errors"

	"github.com/userzhangjinlong/go-zero/core/breaker"
	"github.com/userzhangjinlong/go-zero/zrpc/internal/codes"
	"google.golang.org/grpc"
	gcodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StreamBreakerInterceptor is an interceptor that acts as a circuit breaker.
func StreamBreakerInterceptor(svr any, stream grpc.ServerStream, info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) (err error) {
	breakerName := info.FullMethod
	return breaker.DoWithAcceptable(breakerName, func() error {
		return handler(svr, stream)
	}, codes.Acceptable)
}

// UnaryBreakerInterceptor is an interceptor that acts as a circuit breaker.
func UnaryBreakerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp any, err error) {
	breakerName := info.FullMethod
	err = breaker.DoWithAcceptable(breakerName, func() error {
		var err error
		resp, err = handler(ctx, req)
		return err
	}, codes.Acceptable)
	if errors.Is(err, breaker.ErrServiceUnavailable) {
		err = status.Error(gcodes.Unavailable, err.Error())
	}

	return resp, err
}
