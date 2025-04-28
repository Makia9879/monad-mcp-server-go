package quick_trade

import (
	"context"
	"time"

	"github.com/ThinkInAIXYZ/go-mcp/client"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/Makia9879/monad-mcp-server-go/internal/global"
	types "github.com/Makia9879/monad-mcp-server-go/internal/types/tools/quick_trade"
)

func Register(mcpServer *server.Server) {
	// Register time query tool
	tool, err := protocol.NewTool("quick_trade", "快速交易，用于快速买入指定合约的代币", types.QuickTradeRequest{})
	logx.Must(err)
	mcpServer.RegisterTool(tool, handleQuickTrade)
}

func handleQuickTrade(ctx context.Context, req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	var quickTradeRequest types.QuickTradeRequest
	if err := protocol.VerifyAndUnmarshal(req.RawArguments, &quickTradeRequest); err != nil {
		return nil, err
	}

	// 先浏览器访问 https://testnet.nad.fun/
	t, err := transport.NewStdioClientTransport("/usr/local/bin/node", []string{global.SvcCtx.RunningPathCtx + "/scripts/puppeteer/index.js"})
	if err != nil {
		logx.Errorf("transport.NewStdioClientTransport err: %v", err)
		return nil, err
	}

	cli, err := client.NewClient(t)
	if err != nil {
		logx.Errorf("client.NewClient: %+v", err)
		return nil, err
	}
	defer func() {
		_ = cli.Close()
	}()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// List Tools
	tools, err := cli.ListTools(ctx)
	if err != nil {
		logx.Errorf("Failed to list tools: %v", err)
		return nil, err
	}
	for _, tool := range tools.Tools {
		logx.Infof("- %s: %s", tool.Name, tool.Description)
	}

	return &protocol.CallToolResult{
		Content: []protocol.Content{
			protocol.TextContent{
				Type: "text",
				Text: "完成",
			},
		},
	}, nil
}
