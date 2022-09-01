FROM golang:1.19 AS build

WORKDIR /app

COPY go.* .
RUN go mod download -x

COPY . .

ENV CGO_ENABLED=0
RUN go build -a -installsuffix cgo -o app main.go

FROM alpine:3.16 AS runtime

RUN apk add ca-certificates
RUN update-ca-certificates

# default env
ENV GIN_MODE=release
ENV SERVER_PORT=8080
RUN touch .env

COPY --from=build /app/app ./
EXPOSE ${SERVER_PORT}/tcp
ENTRYPOINT ["./app"]
