package current_time

import (
	"context"
	"fmt"
	"time"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/zeromicro/go-zero/core/logx"

	types "github.com/Makia9879/monad-mcp-server-go/internal/types/tools/current_time"
)

func Register(mcpServer *server.Server) {
	// Register time query tool
	tool, err := protocol.NewTool("current_time", "Get current time for specified timezone", types.TimeRequest{})
	logx.Must(err)
	mcpServer.RegisterTool(tool, handleTimeRequest)
}

func handleTimeRequest(ctx context.Context, req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	var timeReq types.TimeRequest
	if err := protocol.VerifyAndUnmarshal(req.RawArguments, &timeReq); err != nil {
		return nil, err
	}

	loc, err := time.LoadLocation(timeReq.Timezone)
	if err != nil {
		return nil, fmt.Errorf("invalid timezone: %v", err)
	}

	return &protocol.CallToolResult{
		Content: []protocol.Content{
			protocol.TextContent{
				Type: "text",
				Text: time.Now().In(loc).String(),
			},
		},
	}, nil
}
