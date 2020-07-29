package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"studentrecord/server"
	"syscall"
)

func main() {

	listenAddr := ":3000"

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	// 创建server:
	serv := server.NewWebServer(listenAddr)

	done := make(chan struct{}, 1)
	quit := make(chan os.Signal, 1)

	// os.Interrupt: syscall.SIGINT
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGQUIT)

	// 启动另一个goroutine， 监听将要关闭的信号：
	go server.Shutdown(serv, logger, quit, done)

	// 启动 server:
	logger.Println("Server is ready to handle request at", listenAddr)
	err := serv.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v \n", listenAddr, err)
	}

	// 等待已经关闭的信号：
	<-done
	logger.Println("Server stopped")
}
