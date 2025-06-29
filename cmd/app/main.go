package main

import (
	"github.com/alfariiizi/go-service/config"
	"github.com/alfariiizi/go-service/internal/delivery/cmd"
)

func main() {
	_ = config.GetConfig()
	cmd.Execute()
}
