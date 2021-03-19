package transport

import (
	"context"
	"net/http"
)

// HTTP请求参数解析
type HTTPRequestDecoder interface {
	HTTPRequestDecode(ctx context.Context, r *http.Request) (request interface{}, err error)
}

// GRPC请求参数解析
type GRPCRequestDecoder interface {
	GRPCRequestDecode(ctx context.Context, grpcReq interface{}) (request interface{}, err error)
}