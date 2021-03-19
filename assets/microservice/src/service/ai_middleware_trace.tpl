package service

import (
	"Ai/param/req"
	"Ai/param/res"
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"io"
)

type tracerMiddlewareServer struct {
	tracer opentracing.Tracer
	next AIService
}

func NewTracerMiddlewareServer(tracer opentracing.Tracer) NewMiddlewareServer {
	return func(service AIService) AIService {
		return tracerMiddlewareServer{
			tracer: tracer,
			next:   service,
		}
	}
}

func NewJaegerTracer(serviceName string) (tracer opentracing.Tracer, closer io.Closer, err error) {
	cfg := &jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const", //固定采样
			Param: 1,       //1=全采样、0=不采样
		},

		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},

		ServiceName: serviceName,
	}

	tracer, closer, err = cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		return
	}
	opentracing.SetGlobalTracer(tracer)
	return
}

func (tm tracerMiddlewareServer) ShowAiDistort(ctx context.Context, req *req.ShowVO) (rsp *res.ShowAiDistortRsp, err error) {
	span, ctxContext := opentracing.StartSpanFromContextWithTracer(ctx, tm.tracer, "ShowAiDistort", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "NewTracerServerMiddleware",
	})
	defer func() {
		span.LogKV("req", req)
		span.Finish()
	}()
	return tm.next.ShowAiDistort(ctxContext, req)
}

func (tm tracerMiddlewareServer) ShowAiFailure(ctx context.Context, req *req.ShowVO) (rsp *res.ShowAiFailureRsp, err error) {
	span, ctxContext := opentracing.StartSpanFromContextWithTracer(ctx, tm.tracer, "ShowAiFailure", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "NewTracerServerMiddleware",
	})
	defer func() {
		span.LogKV("req", req)
		span.Finish()
	}()
	return tm.next.ShowAiFailure(ctxContext, req)
}

func (tm tracerMiddlewareServer) AddAiDistort(ctx context.Context, req *req.AddAiDistortVO) (rsp *res.Ack, err error) {
	span, ctxContext := opentracing.StartSpanFromContextWithTracer(ctx, tm.tracer, "AddAiDistort", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "NewTracerServerMiddleware",
	})
	defer func() {
		span.LogKV("req", req)
		span.Finish()
	}()
	return tm.next.AddAiDistort(ctxContext, req)
}

func (tm tracerMiddlewareServer) AddAiFailure(ctx context.Context, req *req.AddAiFailureVO) (rsp *res.Ack, err error) {
	span, ctxContext := opentracing.StartSpanFromContextWithTracer(ctx, tm.tracer, "AddAiFailure", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "NewTracerServerMiddleware",
	})
	defer func() {
		span.LogKV("req", req)
		span.Finish()
	}()
	return tm.next.AddAiFailure(ctxContext, req)
}

func (tm tracerMiddlewareServer) DeleteAiDistort(ctx context.Context, req *req.DeleteAiDistortVO) (rsp *res.Ack, err error) {
	span, ctxContext := opentracing.StartSpanFromContextWithTracer(ctx, tm.tracer, "DeleteAiDistort", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "NewTracerServerMiddleware",
	})
	defer func() {
		span.LogKV("req", req)
		span.Finish()
	}()
	return tm.next.DeleteAiDistort(ctxContext, req)
}

func (tm tracerMiddlewareServer) DeleteAiFailure(ctx context.Context, req *req.DeleteAiFailureVO) (rsp *res.Ack, err error) {
	span, ctxContext := opentracing.StartSpanFromContextWithTracer(ctx, tm.tracer, "DeleteAiFailure", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "NewTracerServerMiddleware",
	})
	defer func() {
		span.LogKV("req", req)
		span.Finish()
	}()
	return tm.next.DeleteAiFailure(ctxContext, req)
}

