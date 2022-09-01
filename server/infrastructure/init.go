package infrastructure

import (
	lg "github.com/parpa/golang-gin-openapi/server/logger"
	"go.uber.org/zap"
)

var logger *zap.Logger

func Init() {
	initLogger()
	// todo init static clients
}

func initLogger() {
	logger = lg.Init()
}
