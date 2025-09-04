package main

import (
	"github.com/alfariiizi/vandor/internal/config"
	"github.com/alfariiizi/vandor/internal/delivery/cmd"
	"github.com/alfariiizi/vandor/internal/pkg/logger"
)

func main() {
	_ = config.GetConfig()
	_ = logger.Init()
	cmd.Execute()
}
