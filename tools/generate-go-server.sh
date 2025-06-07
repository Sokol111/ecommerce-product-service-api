#!/bin/bash

rm -rf server go-client
mkdir -p server go-client

oapi-codegen -generate types \
  -package model \
  -o internal/api/model/types.gen.go \
  openapi/openapi.yml

oapi-codegen -generate gin-server,strict-server \
  -package server \
  -o server/server.gen.go \
  openapi/openapi.yml

oapi-codegen -generate client \
  -package client \
  -o client/client.gen.go \
  openapi/openapi.yml