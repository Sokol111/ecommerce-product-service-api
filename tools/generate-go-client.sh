#!/bin/bash
rm -rf go-client
mkdir -p go-client
echo "github.com/Sokol111/ecommerce-product-service-api/go-client

go 1.24.2" > go-client/go.mod
oapi-codegen -generate client -package client \
  -o go-client/client.gen.go openapi/openapi.yml