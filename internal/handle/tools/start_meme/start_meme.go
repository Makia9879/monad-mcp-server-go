package start_meme

import (
	"context"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/Makia9879/monad-mcp-server-go/internal/helper"
	types "github.com/Makia9879/monad-mcp-server-go/internal/types/tools/start_meme"
)

const DefaultMemeURL = "https://testnet.nad.fun/"

func Register(mcpServer *server.Server) {
	tool, err := protocol.NewTool("start_meme", "快速交易，用于快速买入指定合约的代币。关键词：打狗", types.StartMemeRequest{})
	logx.Must(err)
	mcpServer.RegisterTool(tool, handleStartMeme)
}

func handleStartMeme(ctx context.Context, req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	var startMemeRequest types.StartMemeRequest
	if err := protocol.VerifyAndUnmarshal(req.RawArguments, &startMemeRequest); err != nil {
		return nil, err
	}

	// 启动浏览器，访问 startMemeRequest.MemePumpURL
	var pumpURL string
	pumpURL = DefaultMemeURL
	helper.RunChromeDaemon(ctx, pumpURL)

	return &protocol.CallToolResult{
		Content: []protocol.Content{
			protocol.TextContent{
				Text: "执行完毕，提醒用户手动连接好钱包，准备打金狗",
				Type: "text",
			},
		},
	}, nil
}
