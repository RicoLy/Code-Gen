package transport

import (
	"Ai/param/erro"
	"Ai/param/global"
	"Ai/param/req"
	"Ai/param/res"
	"Ai/src/endpoint"
	"context"
	"encoding/json"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"net/http"
)



func MakeHttpHandler(endpoints endpoint.AiEndpoints, log *zap.Logger) http.Handler {
	// 配置
	options := []httptransport.ServerOption{
		// 注入错误编码方法
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			log.Warn(fmt.Sprint(ctx.Value(global.ContextReqUUid)), zap.Error(err))
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(erro.ErrorWrapper{Error: err.Error()})
		}),
		// 注入请求前上下文添加UUID
		httptransport.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			UUID := uuid.NewV5(uuid.Must(uuid.NewV4(), nil), "req_uuid").String()
			log.Debug("给请求添加uuid", zap.Any("UUID", UUID))
			ctx = context.WithValue(ctx, global.ContextReqUUid, UUID)
			return ctx
		}),
		// 注入错误处理方法
		httptransport.ServerErrorHandler(NewZapLogErrorHandler(log)),
	}
	r := mux.NewRouter()

	//获取Ai 误报/漏报配置
	r.Methods("POST").Path("/ai/showAiDistort").Handler(httptransport.NewServer(
		endpoints.ShowAiDistortEndpoint,
		(&req.ShowVO{}).HTTPRequestDecode,
		(&res.ShowAiDistortRsp{}).HTTPResponseEncode,
		options...,
	))

	//获取Ai 误报/漏报配置
	r.Methods("POST").Path("/ai/showAiFailure").Handler(httptransport.NewServer(
		endpoints.ShowAiFailureEndpoint,
		(&req.ShowVO{}).HTTPRequestDecode,
		(&res.ShowAiFailureRsp{}).HTTPResponseEncode,
		options...,
	))

	//添加Ai误报信息
	r.Methods("POST").Path("/ai/addAiDistort").Handler(httptransport.NewServer(
		endpoints.ShowAiFailureEndpoint,
		(&req.AddAiDistortVO{}).HTTPRequestDecode,
		(&res.Ack{}).HTTPResponseEncode,
		options...,
	))

	//添加Ai漏报信息
	r.Methods("POST").Path("/ai/addAiFailure").Handler(httptransport.NewServer(
		endpoints.ShowAiFailureEndpoint,
		(&req.AddAiFailureVO{}).HTTPRequestDecode,
		(&res.Ack{}).HTTPResponseEncode,
		options...,
	))

	//删除Ai误报信息
	r.Methods("POST").Path("/ai/deleteAiDistort").Handler(httptransport.NewServer(
		endpoints.ShowAiFailureEndpoint,
		(&req.DeleteAiDistortVO{}).HTTPRequestDecode,
		(&res.Ack{}).HTTPResponseEncode,
		options...,
	))

	//删除Ai误报信息
	r.Methods("POST").Path("/ai/deleteAiFailure").Handler(httptransport.NewServer(
		endpoints.ShowAiFailureEndpoint,
		(&req.DeleteAiFailureVO{}).HTTPRequestDecode,
		(&res.Ack{}).HTTPResponseEncode,
		options...,
	))



	return r
}