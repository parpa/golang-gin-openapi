#!/bin/bash

cd "$(dirname "$0")"

# init folder
rm -rf /out/org 2>/dev/null
mkdir -p /out/org
mkdir -p /out/dist/controller
mkdir -p /out/dist/model
mkdir -p /out/dist/router
# do generate
/usr/local/bin/docker-entrypoint.sh generate -i /openapi/api.yaml -t /mustache -g go-gin-server -o /out/org --additional-properties="packageName=model"
# cp files
cp /out/org/go/routers.go /out/dist/router
cp /out/org/go/model_*.go /out/dist/model
cp /out/org/go/api_*.go /out/dist/controller
# csplit controller
cd /out/dist/controller
cat api_*.go | csplit -zf controller - '/filename/' {*} 1>/dev/null
rm -f api_*.go
rm -f controller00
for f in controller*; do
    filename="$(head -n 1 $f | sed -r 's/^filename=//' | sed -r 's/([A-Z])/_\L\1/g')"
    sed -i -e '1d' $f
    mv $f $filename
done
