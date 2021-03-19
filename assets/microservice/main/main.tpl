package main

import (
	"Ai/param/pb"
	"Ai/src/dao"
	"Ai/src/endpoint"
	"Ai/src/service"
	"Ai/src/transport"
	"Ai/utils/log"
	"context"
	"flag"
	"fmt"
	log2 "github.com/go-kit/kit/log"
	metricsprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/sd/etcdv3"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"hash/crc32"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// grpc 监听地址
var grpcAddr = flag.String("g", "127.0.0.1:8881", "grpcAddr")

// http 监听地址
var httpAddr = flag.String("h", "127.0.0.1:8882", "httpAddr")

// prometheus 监听地址
var prometheusAddr = flag.String("p", "127.0.0.1:9100", "prometheus addr")

// 退出信号管道
var quitChan = make(chan error, 1)

func main() {
	flag.Parse()

	// 初始化日志
	log.InitLoggerServer()
	logger := log.GetLogger()

	var (
		etcdAddr = []string{"127.0.0.1:2379"}
		srvName  = "svc.ai.agent"
		ttl      = 5 * time.Second
		err      error
	)

	// ETCD 配置
	options := etcdv3.ClientOptions{
		DialTimeout:   ttl,
		DialKeepAlive: ttl,
	}
	// 获取ETCD客户端
	etcdClient, err := etcdv3.NewClient(context.Background(), etcdAddr, options)
	if err != nil {
		logger.Error("[user_agent]  NewClient", zap.Error(err))
		return
	}

	// 服务注册
	Register := etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
		Key:   fmt.Sprintf("%s/%d", srvName, crc32.ChecksumIEEE([]byte(*grpcAddr))),
		Value: *grpcAddr,
	}, log2.NewNopLogger())

	// 初始化数据库
	if err = dao.InitMysql("192.168.142.128", "3306", "root", "mysqlly", "ai"); err != nil {
		logger.Debug("err: ", zap.Any(" mysql", err))
	}

	count := metricsprometheus.NewCounterFrom(prometheus.CounterOpts{
		Subsystem: "user_agent",
		Name:      "request_count",
		Help:      "Number of requests",
	}, []string{"method"})

	histogram := metricsprometheus.NewHistogramFrom(prometheus.HistogramOpts{
		Subsystem: "user_agent",
		Name:      "request_consume",
		Help:      "Request consumes time",
	}, []string{"method"})

	srv := service.NewAIServiceImpl(dao.AiDistort{}, dao.AiFailure{}, count, histogram, logger)

	// 令牌桶服务限流 每秒产生 cpu个数令牌，存储10 * cpu个数个令牌
	golangLimit := rate.NewLimiter(rate.Limit(runtime.NumCPU()), 10*runtime.NumCPU())
	aiEndpoints := endpoint.NewAiEndpoints(srv, logger, golangLimit)

	// 开启协程监听http请求
	go func() {
		httpServer := transport.MakeHttpHandler(aiEndpoints, logger)
		logger.Info("[Ai_Server] http run " + *httpAddr)
		quitChan <- http.ListenAndServe(*httpAddr, httpServer)
	}()

	// 开启协程监听grpc请求
	go func() {
		grpcServer := transport.NewGRPCServer(aiEndpoints, logger)
		grpcListener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			logger.Warn("[Ai_Server] Listen", zap.Error(err))
			quitChan <- err
			return
		}
		//服务注册
		Register.Register()
		baseServer := grpc.NewServer(grpc.UnaryInterceptor(grpctransport.Interceptor))
		pb.RegisterAiServer(baseServer, grpcServer)
		logger.Info("[Ai_Server] grpc run " + *grpcAddr)
		quitChan <- baseServer.Serve(grpcListener)
	}()

	// 服务监控
	go func() {
		logger.Info("[Ai_Server] prometheus run " + *prometheusAddr)
		m := http.NewServeMux()
		m.Handle("/metrics", promhttp.Handler())
		quitChan <- http.ListenAndServe(*prometheusAddr, m)
	}()

	// 开启协程监听退出
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		quitChan <- fmt.Errorf("%s", <-c)
	}()

	// 主协程阻塞等待
	err = <-quitChan
	logger.Debug("err: ", zap.Any("error", err))
}
