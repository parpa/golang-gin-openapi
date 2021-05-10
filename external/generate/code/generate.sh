#!/bin/bash

cd "$(dirname "$0")"
# folder
rm -rf .dist 2>/dev/null
mkdir -p .dist/server
mkdir -p .dist/client
# generate
docker-compose -f openapi.compose.yaml up --build
# cp
server="$(cd ../../../server && pwd)"
service="$server/service"
cp .dist/server/dist/model/* $service/model
cp .dist/server/dist/router/* $service/router
for f in .dist/server/dist/controller/*; do
    if [ ! -f "$service/controller/$(basename $f)" ]; then
        cp $f $service/controller
    fi
done
# formate
gofmt -w $server
