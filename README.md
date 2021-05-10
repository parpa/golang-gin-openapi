# GIN API

## Setup API

```sh
# openapi generate
docker-compose -f init-api.compose.yaml up

# setup
sh setup.sh
```

## Docker

```sh
docker build -t golang-gin-openapi -f Dockerfile .

docker run -p 8080:8080 --rm golang-gin-openapi
```

## Environment variables

| #   | name        | exp. | required | desp |
| --- | ----------- | ---- | -------- | ---- |
| 1   | SERVER_PORT | 8080 | ×        |      |
