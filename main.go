package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var listenAddr string

func init() {
	flag.StringVar(&listenAddr, "listen-addr", "5000", "server listen address")
}

func main() {
	flag.Parse() // 外部参数解析
	listenAddr = fmt.Sprintf(":%s", listenAddr)

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	// 创建server:
	server := newWebServer(logger)

	done := make(chan struct{}, 1)
	quit := make(chan os.Signal, 1)

	// os.Interrupt: syscall.SIGINT
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGQUIT)

	// 启动另一个goroutine， 监听将要关闭的信号：
	go shutdown(server, logger, quit, done)

	// 启动 server:
	logger.Println("Server is ready to handle request at", listenAddr)
	err := server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v \n", listenAddr, err)
	}

	// 等待已经关闭的信号：
	<-done
	logger.Println("Server stopped")
}

func newWebServer(logger *log.Logger) *http.Server {

	// 路由：
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//time.Sleep(50 * time.Second)

		w.WriteHeader(http.StatusOK)
	})

	// http 服务配置:
	server := &http.Server{
		Addr:         listenAddr,
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return server
}

// 关闭server
// quit: 接收关闭信号
// done: 发出已关闭信号
func shutdown(server *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- struct{}) {
	// 等待接收到退出信号
	<-quit
	logger.Println("Sever is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	err := server.Shutdown(ctx)
	if err != nil {
		logger.Fatalf("could not gracefully shutdown the server: %v \n", err)
	}

	// do somthing :

	fmt.Println("do something start ...", time.Now())
	time.Sleep(5 * time.Second)
	fmt.Println("do something end ...", time.Now())

	close(done)
}
