package tools

import (
	"github.com/ThinkInAIXYZ/go-mcp/server"

	"github.com/Makia9879/monad-mcp-server-go/internal/handle/tools/current_time"
)

func RegisterTools(server *server.Server) {
	current_time.Register(server)
}
