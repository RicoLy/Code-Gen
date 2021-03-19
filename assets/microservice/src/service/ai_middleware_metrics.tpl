package service

import (
	"Ai/param/req"
	"Ai/param/res"
	"context"
	"github.com/go-kit/kit/metrics"
	"time"
)

// 服务监控中间件
type metricsMiddlewareServer struct {
	next      AIService
	counter   metrics.Counter
	histogram metrics.Histogram
}

func NewMetricsMiddlewareServer(counter metrics.Counter, histogram metrics.Histogram) NewMiddlewareServer {
	return func(service AIService) AIService {
		return metricsMiddlewareServer{
			next: service,
			counter: counter,
			histogram: histogram,
		}
	}
}

func (m metricsMiddlewareServer) ShowAiDistort(ctx context.Context, req *req.ShowVO) (rsp *res.ShowAiDistortRsp, err error) {
	defer func(start time.Time) {
		method := []string{"method", "ShowAiDistort"}
		m.counter.With(method...).Add(1)
		m.histogram.With(method...).Observe(time.Since(start).Seconds())
	}(time.Now())
	return m.next.ShowAiDistort(ctx, req)
}

func (m metricsMiddlewareServer) ShowAiFailure(ctx context.Context, req *req.ShowVO) (rsp *res.ShowAiFailureRsp, err error) {
	defer func(start time.Time) {
		method := []string{"method", "ShowAiFailure"}
		m.counter.With(method...).Add(1)
		m.histogram.With(method...).Observe(time.Since(start).Seconds())
	}(time.Now())
	return m.next.ShowAiFailure(ctx, req)
}

func (m metricsMiddlewareServer) AddAiDistort(ctx context.Context, req *req.AddAiDistortVO) (rsp *res.Ack, err error) {
	defer func(start time.Time) {
		method := []string{"method", "AddAiDistort"}
		m.counter.With(method...).Add(1)
		m.histogram.With(method...).Observe(time.Since(start).Seconds())
	}(time.Now())
	return m.next.AddAiDistort(ctx, req)
}

func (m metricsMiddlewareServer) AddAiFailure(ctx context.Context, req *req.AddAiFailureVO) (rsp *res.Ack, err error) {
	defer func(start time.Time) {
		method := []string{"method", "AddAiFailure"}
		m.counter.With(method...).Add(1)
		m.histogram.With(method...).Observe(time.Since(start).Seconds())
	}(time.Now())
	return m.next.AddAiFailure(ctx, req)
}

func (m metricsMiddlewareServer) DeleteAiDistort(ctx context.Context, req *req.DeleteAiDistortVO) (rsp *res.Ack, err error) {
	defer func(start time.Time) {
		method := []string{"method", "DeleteAiDistort"}
		m.counter.With(method...).Add(1)
		m.histogram.With(method...).Observe(time.Since(start).Seconds())
	}(time.Now())
	return m.next.DeleteAiDistort(ctx, req)
}

func (m metricsMiddlewareServer) DeleteAiFailure(ctx context.Context, req *req.DeleteAiFailureVO) (rsp *res.Ack, err error) {
	defer func(start time.Time) {
		method := []string{"method", "DeleteAiFailure"}
		m.counter.With(method...).Add(1)
		m.histogram.With(method...).Observe(time.Since(start).Seconds())
	}(time.Now())
	return m.next.DeleteAiFailure(ctx, req)
}



