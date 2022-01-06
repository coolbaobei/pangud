//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/pangud/pangud/pkg/infrastructure/endpoint"
)

func initEndpoint() *endpoint.RestfulEndpoint {
	wire.Build(endpoint.RestfulEndpointSet)
	return &endpoint.RestfulEndpoint{}
}
