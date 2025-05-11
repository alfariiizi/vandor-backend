package main

import (
	"github.com/alfariiizi/go-service/cmd"
	"github.com/alfariiizi/go-service/config"
)

func main() {
	_ = config.GetConfig()
	cmd.Execute()
}
