package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/Makia9879/monad-mcp-server-go/internal/global"
	"github.com/Makia9879/monad-mcp-server-go/internal/handle"
)

func main() {
	initSvcCtx()

	// Create SSE transport server
	transportServer, err := transport.NewSSEServerTransport("127.0.0.1:8080")
	logx.Must(err)

	// Initialize MCP server
	mcpServer, err := server.NewServer(transportServer)
	logx.Must(err)

	// Register time query tool
	handle.RegisterHandles(mcpServer)

	// 注册 ctrl-c信号量回调
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-done
		logx.Info("接收到终止信号，正在优雅关闭...")

		// 执行其他清理操作
		global.SvcCtx.Close()

		os.Exit(0)
	}()

	// Start server
	if err = mcpServer.Run(); err != nil {
		logx.Must(err)
	}
}

func initSvcCtx() {
	var runningPathCtx string
	flag.StringVar(&runningPathCtx, "context", "./", "指明程序运行的上下文路径，用于脚本和日志输出等场景")
	flag.Parse()

	global.SvcCtx = global.NewServiceContext()
	global.SvcCtx.RunningPathCtx = runningPathCtx
}
