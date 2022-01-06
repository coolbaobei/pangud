package main

import (
	"fmt"
	"github.com/pangud/pangud/pkg/infrastructure/config"
)

func main() {
	fmt.Println(1)

	config.LoadConfigFromFile(".")
	endpoint := initEndpoint()
	endpoint.Serve()
}
