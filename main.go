package main

import (
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/Makia9879/monad-mcp-server-go/internal/handle"
)

func main() {
	// Create SSE transport server
	transportServer, err := transport.NewSSEServerTransport("127.0.0.1:8080")
	logx.Must(err)

	// Initialize MCP server
	mcpServer, err := server.NewServer(transportServer)
	logx.Must(err)

	// Register time query tool
	handle.RegisterHandles(mcpServer)

	// Start server
	if err = mcpServer.Run(); err != nil {
		logx.Must(err)
	}
}
