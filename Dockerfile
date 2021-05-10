FROM timbru31/java-node:11-jre-14 AS pre-build

WORKDIR /app

RUN npm install @openapitools/openapi-generator-cli -g
RUN openapi-generator-cli version-manager set 5.1.0

COPY external ./openapi
COPY .openapi-generator-ignore .

RUN openapi-generator-cli generate -i openapi/api.yaml -g go-gin-server -o ./ -c openapi/api-server-config.json

FROM golang:1.16 AS build

WORKDIR /app

COPY go.* .
RUN go mod download -x

COPY server ./server
COPY --from=pre-build /app/server ./server
COPY infrastructure ./infrastructure
COPY logger ./logger
COPY migration ./migration
COPY main.go ./main.go

ENV CGO_ENABLED=0
RUN go build -a -installsuffix cgo -o app main.go

FROM alpine:3.13 AS runtime

RUN apk add ca-certificates
RUN update-ca-certificates

# default env
ENV GIN_MODE=release
ENV SERVER_PORT=8080
RUN touch .env

COPY --from=build /app/app ./
EXPOSE ${SERVER_PORT}/tcp
ENTRYPOINT ["./app"]
