package server

import (
	"github.com/yuxsr/notificator/internal/server/service"
	yuxsrdevpbv1 "github.com/yushi-a/yuxsr-dev-pb/gencode/go/yuxsr_dev_pb/v1"
	"google.golang.org/grpc"
)

func RegisterNewGRPCServer(config service.NotificatorServiceConfig) *grpc.Server {
	notificatorService := service.NewNotificatorService(config)
	s := grpc.NewServer()
	yuxsrdevpbv1.RegisterNotificatorServiceServer(s, notificatorService)
	return s
}
