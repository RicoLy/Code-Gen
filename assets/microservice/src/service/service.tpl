package service

import (
	"{{.ServiceName}}/param/req"
	"{{.ServiceName}}/param/res"
	"{{.ServiceName}}/param/vo"
	"{{.ServiceName}}/utils/generatorId"
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"strings"
)

// {.ServiceName}业务接口
type {.ServiceName}Service interface {
	ShowAiDistort(ctx context.Context, req *req.ShowVO) (rsp *res.ShowAiDistortRsp, err error)
	{{range $j, $method := .Methods}}{{$method.MethodName}}(ctx context.Context, req *req.{{$method.Param}}) (rsp *res.{{$method.Returns}}, err error)
        {{end -}}
}

type {.ServiceName}ServiceImpl struct {
	logger       *zap.Logger
	// 可以注入DAO
}

func New{.ServiceName}ServiceImpl(counter metrics.Counter, histogram metrics.Histogram, log *zap.Logger) {.ServiceName}Service {
	var service {.ServiceName}Service
	service = &{.ServiceName}ServiceImpl{
		logger:       log,
		// 可以注入DAO
	}

	// 装饰日志中间件
	service = NewLogMiddlewareServer(log)(service)

	// 装饰服务监控中间件
	service = NewMetricsMiddlewareServer(counter, histogram)(service)

	// 装饰链路追踪中间件
	tracer, _, err := NewJaegerTracer("AIServer")
	if err != nil {
		log.Warn(fmt.Sprint(zap.Any("jaeger tracer init failed:", err)))
		return service
	}
	service = NewTracerMiddlewareServer(tracer)(service)

	return service
}

//获取Ai 误报/漏报配置
func (s AIServiceImpl) ShowAiDistort(ctx context.Context, req *req.ShowVO) (rsp *res.ShowAiDistortRsp, err error) {
	rsp = new(res.ShowAiDistortRsp)
	var pagination vo.Pagination
	var whereStr = ""
	var param = make([]interface{}, 0)
	pagination.Page = req.Page
	pagination.PageSize = req.PageSize
	if req.Where != "" {
		kvs := strings.Split(req.Where, "&")
		for _, kv := range kvs {
			skv := strings.Split(kv, "=")
			if whereStr != "" {
				whereStr += " and "
			}
			whereStr = whereStr + skv[0] + " = ?"
			param = append(param, skv[1])
		}
	}

	aiDistort, err := s.aiDistortDAO.Find(whereStr, &pagination, param...)
	if err != nil {
		return
	}

	rsp.AiDistorts = res.BuildDistorts(aiDistort)
	rsp.Pagination = pagination

	return
}

//获取Ai 误报/漏报配置
func (s AIServiceImpl) ShowAiFailure(ctx context.Context, req *req.ShowVO) (rsp *res.ShowAiFailureRsp, err error) {
	rsp = new(res.ShowAiFailureRsp)
	var pagination vo.Pagination
	var whereStr string
	var param = make([]interface{}, 0)
	pagination.Page = req.Page
	pagination.PageSize = req.PageSize
	if req.Where != "" {
		kvs := strings.Split(req.Where, "&")
		for _, kv := range kvs {
			skv := strings.Split(kv, "=")
			if whereStr != "" {
				whereStr += " and "
			}
			whereStr = whereStr + skv[0] + " = ?"
			param = append(param, skv[1])
		}
	}
	AiFailure, err := s.aiFailureDAO.Find(whereStr, &pagination, param...)
	if err != nil {
		return
	}

	rsp.AiFailures = res.BuildAiFailures(AiFailure)
	rsp.Pagination = pagination

	return
}

//添加Ai误报信息
func (s AIServiceImpl) AddAiDistort(ctx context.Context, req *req.AddAiDistortVO) (rsp *res.Ack, err error) {
	rsp = new(res.Ack)
	if err = s.aiDistortDAO.Add(&dao.AiDistort{
		Uin:     generatorId.NextId(), // 雪花算法生成id
		Appid:   req.Appid,
		Domain:  req.Domain,
		Payload: req.Payload,
		From:    req.From,
		Remark:  req.Remark,
		Status:  dao.NotLearned, // 默认未学习
	}); err != nil {
		rsp.IsOk = false
		return
	}
	rsp.IsOk = true
	return
}

//添加Ai漏报信息
func (s AIServiceImpl) AddAiFailure(ctx context.Context, req *req.AddAiFailureVO) (rsp *res.Ack, err error) {
	rsp = new(res.Ack)
	if err = s.aiFailureDAO.Add(&dao.AiFailure{
		Model:   gorm.Model{},
		Uin:     generatorId.NextId(),
		Appid:   req.Appid,
		Domain:  req.Domain,
		Payload: req.Payload,
		Sign:    req.Sign,
		From:    req.From,
		Remark:  req.Remark,
		Status:  dao.NotLearned,
	}); err != nil {
		rsp.IsOk = false
		return
	}
	rsp.IsOk = true
	return
}

//删除Ai误报信息
func (s AIServiceImpl) DeleteAiDistort(ctx context.Context, req *req.DeleteAiDistortVO) (rsp *res.Ack, err error) {
	rsp = new(res.Ack)
	if err = s.aiDistortDAO.Delete("uin = ? ", req.Uin); err != nil {
		rsp.IsOk = false
		return
	}
	rsp.IsOk = true
	return
}

//删除Ai漏报信息
func (s AIServiceImpl) DeleteAiFailure(ctx context.Context, req *req.DeleteAiFailureVO) (rsp *res.Ack, err error) {
	rsp = new(res.Ack)
	if err = s.aiFailureDAO.Delete("uin = ? ", req.Uin); err != nil {
		rsp.IsOk = false
		return
	}
	rsp.IsOk = true
	return
}
