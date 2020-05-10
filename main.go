package main

import (
	"context"
	"fmt"
	"github.com/idoqo/reads/handler"
	"github.com/idoqo/reads/logger"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	port = ":9500"
)

func main() {
	httpHandler := handler.NewHandler()
	server := &http.Server{
		Handler: httpHandler,
	}
	listener, err := net.Listen("tcp", port)
	if err != nil {
		logger.Panic(err.Error())
	}
	go func () {
		server.Serve(listener)
	}()
	defer Stop(server)

	logger.Info(fmt.Sprintf("Started API server on %s", port))
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	logger.Info(fmt.Sprint(<-ch))
	logger.Info("Shutting down API Server")
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Info("Can't shutdown API server correctly")
	}
}
