package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
	"user-service/config"
	"user-service/internal/db"
	"user-service/internal/service"
	"user-service/proto"
)

// Init 读取配置信息，并解析，完成日志初始化
func Init() {
	// 日志初始化
	//log.InitLog()
	if err := config.Init(); err != nil {
		log.Fatalf("init config failed, err:%v\n", err)
	}

	log.Info("log init success...")
	//utils.InitSvrConn()
	log.Info("InitSvrConn success...")
}

func Run() error {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.GetGlobalConfig().SvrConfig.Host, config.GetGlobalConfig().SvrConfig.Port))
	if err != nil {
		log.Fatalf("listen: error %v", err)
		return fmt.Errorf("listen: error %v", err)
	}

	// 端口监听成功，启动 grpc
	server := grpc.NewServer()
	// 注册服务
	proto.RegisterUserServiceServer(server, &service.UserService{})
	// 注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 启动
	go func() {
		log.Infof("Weather listening on %s:%d", config.GetGlobalConfig().SvrConfig.Host,
			config.GetGlobalConfig().SvrConfig.Port)
		reflection.Register(server)
		err = server.Serve(listen)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()
	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	return nil
}

func main() {
	Init()
	//defer log.Sync()
	defer db.CloseDB()
	if err := Run(); err != nil {
		log.Errorf("usersvr run err:%v", err)
	}
}
