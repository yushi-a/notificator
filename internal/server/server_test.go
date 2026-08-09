package server

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"connectrpc.com/connect"
	"github.com/yushi-a/notificator/internal/server/service"
	notificationv1 "github.com/yushi-a/yuxsr-dev-pb/gen/go/yuxsr/notification/v1"
	"github.com/yushi-a/yuxsr-dev-pb/gen/go/yuxsr/notification/v1/notificationv1connect"
	"golang.org/x/net/http2"
)

type fakeClient struct {
	messages []string
}

func (f *fakeClient) Notify(_ context.Context, message string) error {
	f.messages = append(f.messages, message)
	return nil
}

// h2cClient は TLS なしで HTTP/2 を話すクライアントを返す。
func h2cClient() *http.Client {
	return &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLSContext: func(ctx context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
				var d net.Dialer
				return d.DialContext(ctx, network, addr)
			},
		},
	}
}

// TestNewHandler は Connect / gRPC / ブラウザ相当の JSON POST が
// 同一ハンドラで処理されることを確認する。
func TestNewHandler(t *testing.T) {
	client := &fakeClient{}
	srv := httptest.NewServer(NewHandler(service.NotificatorServiceConfig{Client: client}))
	defer srv.Close()

	t.Run("browser JSON POST", func(t *testing.T) {
		res, err := http.Post(
			srv.URL+"/yuxsr.notification.v1.NotificatorService/Notify",
			"application/json",
			strings.NewReader(`{"message":"json"}`),
		)
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusOK)
		}
	})

	t.Run("connect protocol", func(t *testing.T) {
		c := notificationv1connect.NewNotificatorServiceClient(srv.Client(), srv.URL)
		if _, err := c.Notify(t.Context(), connect.NewRequest(&notificationv1.NotifyRequest{Message: "connect"})); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("grpc protocol over h2c", func(t *testing.T) {
		c := notificationv1connect.NewNotificatorServiceClient(h2cClient(), srv.URL, connect.WithGRPC())
		if _, err := c.Notify(t.Context(), connect.NewRequest(&notificationv1.NotifyRequest{Message: "grpc"})); err != nil {
			t.Fatal(err)
		}
	})

	if want := []string{"json", "connect", "grpc"}; len(client.messages) != len(want) {
		t.Fatalf("messages = %v, want %v", client.messages, want)
	}
}
