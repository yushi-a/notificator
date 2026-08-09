// Package service is application logic
package service

import (
	"context"

	"connectrpc.com/connect"
	notificationv1 "github.com/yushi-a/yuxsr-dev-pb/gen/go/yuxsr/notification/v1"
)

type NotificatorServiceConfig struct {
	Client Client
}

type notificatorService struct {
	client Client
}

func NewNotificatorService(config NotificatorServiceConfig) *notificatorService {
	return &notificatorService{
		client: config.Client,
	}
}

func (n *notificatorService) Notify(
	ctx context.Context,
	req *connect.Request[notificationv1.NotifyRequest],
) (*connect.Response[notificationv1.NotifyResponse], error) {
	if err := n.client.Notify(ctx, req.Msg.Message); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&notificationv1.NotifyResponse{}), nil
}
