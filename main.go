package main

import (
	"context"
	"fmt"
	"goDemo/router"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 设置路由
	router.SetupRoutes()

	server := &http.Server{
		Addr: ":8080",
	}

	// 启动服务器
	fmt.Println("Server started at http://localhost:8080")
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				// 处理监听失败的错误
				// 记录错误
				fmt.Printf("HTTP服务器失败: %v", err)
			}
		}
	}()

	// 等待中断信号来优雅地关闭服务器
	stop := make(chan os.Signal, 1)
	// 用 signal.Notify 来监听 os.Interrupt 信号，这是用户向程序发送中断信号（如Ctrl+C）时产生的信号
	signal.Notify(stop, os.Interrupt)

	<-stop // 程序在此处阻塞，直到接收到一个中断信号
	fmt.Println("stop here")

	//当有中断信号来，创建一个带有超时的 context.Context 对象，超时时间为5秒
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 确保在函数返回时取消这个上下文，释放相关资源
	defer cancel()

	//当接收到中断信号时，调用 server.Shutdown 方法并传入上面创建的 ctx 对象，以优雅地关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		// 如果在关闭过程中出现错误
		fmt.Println("处理关闭服务器时的错误")
	}

}
