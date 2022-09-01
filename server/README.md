# Server

## Local start

```sh
# start
sh start.sh
```

## Docker start

```sh
docker build -t golang-gin-openapi -f Dockerfile .

docker run -p 9090:8080 --rm golang-gin-openapi
```

## Environment variables

| #   | name        | exp.       | required | desp |
| --- | ----------- | ---------- | -------- | ---- |
| 1   | SERVER_PORT | 8080       | ×        |      |
| 2   | LOG_LEVEL   | info       | ×        |      |
| 3   | LOG_FILE    | server.log | ×        |      |
