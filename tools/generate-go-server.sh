#!/bin/bash
rm -rf go-server
mkdir -p go-server

cd go-server || exit 1
go mod init github.com/Sokol111/ecommerce-product-service-api/go-server
cd ..

oapi-codegen -generate gin-server,types,strict-server -package api \
  -o go-server/server.gen.go openapi/openapi.yml

cd go-server || exit 1
go mod tidy