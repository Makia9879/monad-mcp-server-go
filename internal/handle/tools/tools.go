package tools

import (
	"github.com/ThinkInAIXYZ/go-mcp/server"

	"github.com/Makia9879/monad-mcp-server-go/internal/handle/tools/current_time"
	"github.com/Makia9879/monad-mcp-server-go/internal/handle/tools/quick_trade"
	"github.com/Makia9879/monad-mcp-server-go/internal/handle/tools/start_meme"
)

func RegisterTools(server *server.Server) {
	current_time.Register(server)
	quick_trade.Register(server)
	start_meme.Register(server)
}
