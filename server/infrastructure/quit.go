package infrastructure

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Quit - Wait for interrupt signal to gracefully shutdown the server.
func Quit(q func()) chan bool {
	quit := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		<-quit
		defer func() {
			done <- true
		}()
		defer logger.Sync()
		defer func() {
			log.Println("[server] Close XXX Client ...")
			// todo close some static client
		}()

		defer q()
	}()

	return done
}
