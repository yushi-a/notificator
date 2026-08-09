package server

import (
	"net/http"

	"github.com/yushi-a/notificator/internal/server/service"
	"github.com/yushi-a/yuxsr-dev-pb/gen/go/yuxsr/notification/v1/notificationv1connect"
)

// NewHandler は Connect / gRPC / gRPC-Web を同一ポートでさばくハンドラを返す。
func NewHandler(config service.NotificatorServiceConfig) http.Handler {
	notificatorService := service.NewNotificatorService(config)

	mux := http.NewServeMux()
	mux.Handle(notificationv1connect.NewNotificatorServiceHandler(notificatorService))
	return mux
}

// Protocols は HTTP/1.1 と TLS なしの HTTP/2 を受け付ける設定を返す。
// gRPC クライアントは TLS なしの HTTP/2 (h2c) で接続してくるため必要。
func Protocols() *http.Protocols {
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true)
	return protocols
}

// NewServer は addr で待ち受ける notificator サーバを返す。
func NewServer(addr string, config service.NotificatorServiceConfig) *http.Server {
	return &http.Server{
		Addr:      addr,
		Handler:   NewHandler(config),
		Protocols: Protocols(),
	}
}
