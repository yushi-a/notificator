/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/yushi-a/notificator/internal/client"
	"github.com/yushi-a/notificator/internal/server"
	"github.com/yushi-a/notificator/internal/server/service"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve notificator server",
	Long:  `Serve notificator server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := getNotificatorConfig()
		return Serve(config)
	},
}

// Serve is run server
func Serve(config NotificatorConfig) error {
	ctx := context.Background()
	clinet, err := client.NewLineClient(ctx, client.LineClientConfig{
		ChannelSecret:      config.LineChannelSecret,
		ChannelAccessToken: config.LineChannelAccessToken,
		UserID:             config.LineUserID,
	})
	if err != nil {
		return err
	}

	handler := server.NewHandler(service.NotificatorServiceConfig{
		Client: clinet,
	})

	return http.ListenAndServe(":50051", handler)
}
