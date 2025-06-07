#!/bin/bash
rm -rf go-server
mkdir -p go-server
echo "github.com/Sokol111/ecommerce-product-service-api/go-server

go 1.24.2" > go-server/go.mod
oapi-codegen -generate gin-server,types,strict-server -package api \
  -o go-server/server.gen.go openapi/openapi.yml