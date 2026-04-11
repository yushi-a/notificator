// Package service is application logic
package service

import (
	"context"

	yuxsrdevpbv1 "github.com/yushi-a/yuxsr-dev-pb/gencode/go/yuxsr_dev_pb/v1"
)

type NotificatorServiceConfig struct {
	Client Client
}

type notificatorService struct {
	client Client
}

func NewNotificatorService(config NotificatorServiceConfig) yuxsrdevpbv1.NotificatorServiceServer {
	return &notificatorService{
		client: config.Client,
	}
}

func (n *notificatorService) Notify(ctx context.Context, req *yuxsrdevpbv1.NotifyRequest) (*yuxsrdevpbv1.NotifyResponse, error) {
	if err := n.client.Notify(ctx, req.Message); err != nil {
		return nil, err
	}
	return &yuxsrdevpbv1.NotifyResponse{}, nil
}
