package service

import (
	"Ai/param/global"
	"Ai/param/req"
	"Ai/param/res"
	"context"
	"fmt"
	"go.uber.org/zap"
)

type NewMiddlewareServer func(service AIService) AIService

// 服务日志中间件
type logMiddlewareServer struct {
	logger *zap.Logger
	next   AIService
}

func NewLogMiddlewareServer(log *zap.Logger) NewMiddlewareServer {
	return func(service AIService) AIService {
		return logMiddlewareServer{
			logger: log,
			next:   service,
		}
	}
}

func (l logMiddlewareServer) ShowAiDistort(ctx context.Context, req *req.ShowVO) (rsp *res.ShowAiDistortRsp, err error) {
	defer func() {
		l.logger.Debug(fmt.Sprint(ctx.Value(global.ContextReqUUid)), zap.Any("调用 ShowAiDistort logMiddlewareServer", "ShowAiDistort"), zap.Any("req", req), zap.Any("res", rsp))
	}()
	return l.next.ShowAiDistort(ctx, req)
}

func (l logMiddlewareServer) ShowAiFailure(ctx context.Context, req *req.ShowVO) (rsp *res.ShowAiFailureRsp, err error) {
	defer func() {
		l.logger.Debug(fmt.Sprint(ctx.Value(global.ContextReqUUid)), zap.Any("调用 ShowAiFailure logMiddlewareServer", "ShowAiFailure"), zap.Any("req", req), zap.Any("res", rsp))
	}()
	return l.next.ShowAiFailure(ctx, req)
}

func (l logMiddlewareServer) AddAiDistort(ctx context.Context, req *req.AddAiDistortVO) (rsp *res.Ack, err error) {
	defer func() {
		l.logger.Debug(fmt.Sprint(ctx.Value(global.ContextReqUUid)), zap.Any("调用 AddAiDistort logMiddlewareServer", "AddAiDistort"), zap.Any("req", req), zap.Any("res", err))
	}()
	return l.next.AddAiDistort(ctx, req)
}

func (l logMiddlewareServer) AddAiFailure(ctx context.Context, req *req.AddAiFailureVO) (rsp *res.Ack, err error) {
	defer func() {
		l.logger.Debug(fmt.Sprint(ctx.Value(global.ContextReqUUid)), zap.Any("调用 AddAiFailure logMiddlewareServer", "AddAiFailure"), zap.Any("req", req), zap.Any("res", err))
	}()
	return l.next.AddAiFailure(ctx, req)
}

func (l logMiddlewareServer) DeleteAiDistort(ctx context.Context, req *req.DeleteAiDistortVO) (rsp *res.Ack, err error) {
	defer func() {
		l.logger.Debug(fmt.Sprint(ctx.Value(global.ContextReqUUid)), zap.Any("调用 DeleteAiDistort logMiddlewareServer", "DeleteAiDistort"), zap.Any("req", req), zap.Any("res", err))
	}()
	return l.next.DeleteAiDistort(ctx, req)
}

func (l logMiddlewareServer) DeleteAiFailure(ctx context.Context, req *req.DeleteAiFailureVO) (rsp *res.Ack, err error) {
	defer func() {
		l.logger.Debug(fmt.Sprint(ctx.Value(global.ContextReqUUid)), zap.Any("调用 DeleteAiFailure logMiddlewareServer", "DeleteAiFailure"), zap.Any("req", req), zap.Any("res", err))
	}()
	return l.next.DeleteAiFailure(ctx, req)
}
