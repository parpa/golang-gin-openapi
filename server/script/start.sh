#!/bin/bash

cd "$(dirname "$0")"

cd ../

export GIN_MODE=debug
export SERVER_PORT=9090

go run main.go
