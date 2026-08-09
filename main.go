/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/yushi-a/notificator/cmd"
	"github.com/yushi-a/notificator/internal/logging"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logging.NewLogger().Error("cmd execution error", "error", err)
	}
}
