package main

import (
	"github.com/alfariiizi/vandor/config"
	"github.com/alfariiizi/vandor/internal/delivery/cmd"
)

func main() {
	_ = config.GetConfig()
	cmd.Execute()
}
