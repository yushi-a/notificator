package server

import (
	"net/http"

	"github.com/yushi-a/notificator/internal/server/service"
	"github.com/yushi-a/yuxsr-dev-pb/gen/go/yuxsr/notification/v1/notificationv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// NewHandler は Connect / gRPC / gRPC-Web を同一ポートでさばくハンドラを返す。
func NewHandler(config service.NotificatorServiceConfig) http.Handler {
	notificatorService := service.NewNotificatorService(config)

	mux := http.NewServeMux()
	mux.Handle(notificationv1connect.NewNotificatorServiceHandler(notificatorService))

	// h2c で TLS なしの HTTP/2 を受ける (gRPC クライアント互換のため)。
	return h2c.NewHandler(mux, &http2.Server{})
}
