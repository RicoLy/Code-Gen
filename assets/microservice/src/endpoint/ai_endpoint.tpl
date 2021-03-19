package endpoint

import (
	"Ai/param/req"
	"Ai/param/res"
	"Ai/src/service"
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

// 定义endpoints
type AiEndpoints struct {
	ShowAiDistortEndpoint   endpoint.Endpoint
	ShowAiFailureEndpoint   endpoint.Endpoint
	AddAiDistortEndpoint    endpoint.Endpoint
	AddAiFailureEndpoint    endpoint.Endpoint
	DeleteAiDistortEndpoint endpoint.Endpoint
	DeleteAiFailureEndpoint endpoint.Endpoint
}

// 构造endpoints
func NewAiEndpoints(srv service.AIService, log *zap.Logger, limit *rate.Limiter) AiEndpoints {
	var showAiDistortEndpoint endpoint.Endpoint
	{
		showAiDistortEndpoint = MakeShowAiDistortEndPoint(srv)
		// 装饰服务限流中间件
		showAiDistortEndpoint = GolangRateAllowMiddleware(limit)(showAiDistortEndpoint)
		// 装饰日志中间件
		showAiDistortEndpoint = LoggingMiddleware(log)(showAiDistortEndpoint)
	}

	var showAiFailureEndpoint endpoint.Endpoint
	{
		showAiFailureEndpoint = MakeShowAiFailureEndPoint(srv)
		// 装饰服务限流中间件
		showAiFailureEndpoint = GolangRateAllowMiddleware(limit)(showAiFailureEndpoint)
		// 装饰日志中间件
		showAiFailureEndpoint = LoggingMiddleware(log)(showAiFailureEndpoint)
	}

	var addAiDistortEndpoint endpoint.Endpoint
	{
		addAiDistortEndpoint = MakeAddAiDistortEndPoint(srv)
		// 装饰服务限流中间件
		addAiDistortEndpoint = GolangRateAllowMiddleware(limit)(addAiDistortEndpoint)
		// 装饰日志中间件
		addAiDistortEndpoint = LoggingMiddleware(log)(addAiDistortEndpoint)
	}

	var addAiFailureEndpoint endpoint.Endpoint
	{
		addAiFailureEndpoint = MakeAddAiFailureEndPoint(srv)
		// 装饰服务限流中间件
		addAiFailureEndpoint = GolangRateAllowMiddleware(limit)(addAiFailureEndpoint)
		// 装饰日志中间件
		addAiFailureEndpoint = LoggingMiddleware(log)(addAiFailureEndpoint)
	}

	var deleteAiDistortEndpoint endpoint.Endpoint
	{
		deleteAiDistortEndpoint = MakeDeleteAiDistortEndPoint(srv)
		// 装饰服务限流中间件
		deleteAiDistortEndpoint = GolangRateAllowMiddleware(limit)(deleteAiDistortEndpoint)
		// 装饰日志中间件
		deleteAiDistortEndpoint = LoggingMiddleware(log)(deleteAiDistortEndpoint)
	}

	var deleteAiFailureEndpoint endpoint.Endpoint
	{
		deleteAiFailureEndpoint = MakeDeleteAiFailureEndPoint(srv)
		// 装饰服务限流中间件
		deleteAiFailureEndpoint = GolangRateAllowMiddleware(limit)(deleteAiFailureEndpoint)
		// 装饰日志中间件
		deleteAiFailureEndpoint = LoggingMiddleware(log)(deleteAiFailureEndpoint)
	}
	return AiEndpoints{
		ShowAiDistortEndpoint:   showAiDistortEndpoint,
		ShowAiFailureEndpoint:   showAiFailureEndpoint,
		AddAiDistortEndpoint:    addAiDistortEndpoint,
		AddAiFailureEndpoint:    addAiFailureEndpoint,
		DeleteAiDistortEndpoint: deleteAiDistortEndpoint,
		DeleteAiFailureEndpoint: deleteAiFailureEndpoint,
	}
}

//获取Ai 误报/漏报配置endpoint
func (e *AiEndpoints) ShowAiDistort(ctx context.Context, req *req.ShowVO) (rsp *res.ShowAiDistortRsp, err error) {
	resp, err := e.ShowAiDistortEndpoint(ctx, req)
	if err != nil {
		return
	}
	return resp.(*res.ShowAiDistortRsp), err
}

//构造 获取Ai 误报/漏报配置endpoint
func MakeShowAiDistortEndPoint(svc service.AIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		re := request.(*req.ShowVO)
		return svc.ShowAiDistort(ctx, re)
	}
}

//获取Ai 误报/漏报配置endpoint
func (e *AiEndpoints) ShowAiFailure(ctx context.Context, req *req.ShowVO) (rsp *res.ShowAiFailureRsp, err error) {
	resp, err := e.ShowAiFailureEndpoint(ctx, req)
	if err != nil {
		return
	}
	return resp.(*res.ShowAiFailureRsp), err
}

//构造 获取Ai 误报/漏报配置endpoint
func MakeShowAiFailureEndPoint(svc service.AIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		re := request.(*req.ShowVO)
		return svc.ShowAiFailure(ctx, re)
	}
}

//添加 Ai误报配置endpoint
func (e *AiEndpoints) AddAiDistort(ctx context.Context, req *req.AddAiDistortVO) (rsp *res.Ack, err error) {
	resp, err := e.AddAiDistortEndpoint(ctx, req)
	if err != nil {
		return
	}
	return resp.(*res.Ack), err
}

//构造 添加 Ai误报配置endpoint
func MakeAddAiDistortEndPoint(svc service.AIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		re := request.(*req.AddAiDistortVO)
		return svc.AddAiDistort(ctx, re)
	}
}

//添加 Ai漏报配置endpoint
func (e *AiEndpoints) AddAiFailure(ctx context.Context, req *req.AddAiFailureVO) (rsp *res.Ack, err error) {
	resp, err := e.AddAiFailureEndpoint(ctx, req)
	if err != nil {
		return
	}
	return resp.(*res.Ack), err
}

//构造 添加 Ai漏报配置endpoint
func MakeAddAiFailureEndPoint(svc service.AIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		re := request.(*req.AddAiFailureVO)
		return svc.AddAiFailure(ctx, re)
	}
}

//删除 Ai误报配置endpoint
func (e *AiEndpoints) DeleteAiDistort(ctx context.Context, req *req.DeleteAiDistortVO) (rsp *res.Ack, err error) {
	resp, err := e.DeleteAiDistortEndpoint(ctx, req)
	if err != nil {
		return
	}
	return resp.(*res.Ack), err
}

//构造 删除 Ai误报配置endpoint
func MakeDeleteAiDistortEndPoint(svc service.AIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		re := request.(*req.DeleteAiDistortVO)
		return svc.DeleteAiDistort(ctx, re)
	}
}

//删除 Ai漏报配置endpoint
func (e *AiEndpoints) DeleteAiFailure(ctx context.Context, req *req.DeleteAiFailureVO) (rsp *res.Ack, err error) {
	resp, err := e.DeleteAiFailureEndpoint(ctx, req)
	if err != nil {
		return
	}
	return resp.(*res.Ack), err
}

//构造 删除 Ai漏报配置endpoint
func MakeDeleteAiFailureEndPoint(svc service.AIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		re := request.(*req.DeleteAiFailureVO)
		return svc.DeleteAiFailure(ctx, re)
	}
}
