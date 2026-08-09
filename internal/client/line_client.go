package client

import (
	"context"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/yushi-a/notificator/internal/logging"
	"github.com/yushi-a/notificator/internal/server/service"
)

type LineClientConfig struct {
	ChannelSecret      string
	ChannelAccessToken string
	UserID             string
}

type lineClient struct {
	client *linebot.Client
	userID string
}

func NewLineClient(ctx context.Context, config LineClientConfig) (service.Client, error) {
	client, err := linebot.New(config.ChannelSecret, config.ChannelAccessToken)
	if err != nil {
		return nil, err
	}
	return &lineClient{
		client: client,
		userID: config.UserID,
	}, nil
}

func (l *lineClient) Notify(ctx context.Context, message string) error {
	logger := logging.NewLogger()
	_, err := l.client.PushMessage(l.userID, linebot.NewTextMessage(message)).Do()
	if err != nil {
		logger.Error("failed to push message", "error", err)
		return err
	}
	return nil
}
