package transport

import (
	"Ai/param/global"
	"Ai/param/pb"
	"Ai/param/req"
	"Ai/param/res"
	"Ai/src/endpoint"
	"Ai/utils/log"
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

//实现protobuf中定义的接口
type grpcServer struct {
	showAiDistort   grpctransport.Handler
	showAiFailure   grpctransport.Handler
	addAiDistort    grpctransport.Handler
	addAiFailure    grpctransport.Handler
	deleteAiDistort grpctransport.Handler
	deleteAiFailure grpctransport.Handler
}

func NewGRPCServer(endpoints endpoint.AiEndpoints, log *zap.Logger) pb.AiServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
			ctx = context.WithValue(ctx, global.ContextReqUUid, md.Get(global.ContextReqUUid))
			return ctx
		}),
		grpctransport.ServerErrorHandler(NewZapLogErrorHandler(log)),
	}

	return &grpcServer{
		showAiDistort: grpctransport.NewServer( //获取Ai 误报/漏报配置
			endpoints.ShowAiDistortEndpoint,
			(&req.ShowVO{}).GRPCRequestDecode,
			(&res.ShowAiDistortRsp{}).GRPCResponseEncode,
			options...,
		),
		showAiFailure: grpctransport.NewServer( //获取Ai 误报/漏报配置
			endpoints.ShowAiFailureEndpoint,
			(&req.ShowVO{}).GRPCRequestDecode,
			(&res.ShowAiFailureRsp{}).GRPCResponseEncode,
			options...,
		),
		addAiDistort: grpctransport.NewServer( //添加Ai误报信息
			endpoints.AddAiDistortEndpoint,
			(&req.AddAiDistortVO{}).GRPCRequestDecode,
			(&res.Ack{}).GRPCResponseEncode,
			options...,
		),
		addAiFailure: grpctransport.NewServer( //添加Ai漏报信息
			endpoints.AddAiFailureEndpoint,
			(&req.AddAiFailureVO{}).GRPCRequestDecode,
			(&res.Ack{}).GRPCResponseEncode,
			options...,
		),
		deleteAiDistort: grpctransport.NewServer( //删除Ai误报信息
			endpoints.DeleteAiDistortEndpoint,
			(&req.DeleteAiDistortVO{}).GRPCRequestDecode,
			(&res.Ack{}).GRPCResponseEncode,
			options...,
		),
		deleteAiFailure: grpctransport.NewServer( //删除Ai误报信息
			endpoints.DeleteAiFailureEndpoint,
			(&req.DeleteAiFailureVO{}).GRPCRequestDecode,
			(&res.Ack{}).GRPCResponseEncode,
			options...,
		),
	}
}

func (s *grpcServer) ShowAiDistort(ctx context.Context, req *pb.ShowReq) (*pb.ShowAiDistortRsp, error) {
	_, rep, err := s.showAiDistort.ServeGRPC(ctx, req)
	if err != nil {
		log.GetLogger().Warn("s.ShowAiDistort.ServeGRPC", zap.Error(err))
		return nil, err
	}
	return rep.(*pb.ShowAiDistortRsp), nil
}

func (s *grpcServer) ShowAiFailure(ctx context.Context, req *pb.ShowReq) (*pb.ShowAiFailureRsp, error) {
	_, rep, err := s.showAiFailure.ServeGRPC(ctx, req)
	if err != nil {
		log.GetLogger().Warn("s.ShowAiFailure.ServeGRPC", zap.Error(err))
		return nil, err
	}
	return rep.(*pb.ShowAiFailureRsp), nil
}

func (s *grpcServer) AddAiDistort(ctx context.Context, req *pb.AddAiDistortReq) (*pb.Ack, error) {
	_, rep, err := s.addAiDistort.ServeGRPC(ctx, req)
	if err != nil {
		log.GetLogger().Warn("s.AddAiDistort.ServeGRPC", zap.Error(err))
		return nil, err
	}
	return rep.(*pb.Ack), nil
}

func (s *grpcServer) AddAiFailure(ctx context.Context, req *pb.AddAiFailureReq) (*pb.Ack, error) {
	_, rep, err := s.addAiFailure.ServeGRPC(ctx, req)
	if err != nil {
		log.GetLogger().Warn("s.AddAiFailure.ServeGRPC", zap.Error(err))
		return nil, err
	}
	return rep.(*pb.Ack), nil
}

func (s *grpcServer) DeleteAiDistort(ctx context.Context, req *pb.DeleteAiDistortReq) (*pb.Ack, error) {
	_, rep, err := s.deleteAiDistort.ServeGRPC(ctx, req)
	if err != nil {
		log.GetLogger().Warn("s.DeleteAiDistort.ServeGRPC", zap.Error(err))
		return nil, err
	}
	return rep.(*pb.Ack), nil
}

func (s *grpcServer) DeleteAiFailure(ctx context.Context, req *pb.DeleteAiFailureReq) (*pb.Ack, error) {
	_, rep, err := s.deleteAiFailure.ServeGRPC(ctx, req)
	if err != nil {
		log.GetLogger().Warn("s.DeleteAiFailure.ServeGRPC", zap.Error(err))
		return nil, err
	}
	return rep.(*pb.Ack), nil
}

