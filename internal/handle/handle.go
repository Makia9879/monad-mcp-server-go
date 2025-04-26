package handle

import (
	"github.com/ThinkInAIXYZ/go-mcp/server"

	"github.com/Makia9879/monad-mcp-server-go/internal/handle/tools"
)

func RegisterHandles(mcpServer *server.Server) {
	tools.RegisterTools(mcpServer)
}
