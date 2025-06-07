#!/bin/bash
rm -rf js-client
npx @openapitools/openapi-generator-cli generate \
  -i openapi/openapi.yml \
  -g javascript \
  -o js-client \
  --additional-properties=usePromises=true,projectName=openapi-client

cp tools/js-package.json js-client/package.json