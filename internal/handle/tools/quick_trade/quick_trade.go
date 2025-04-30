package quick_trade

import (
	"context"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/chromedp/chromedp"
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

	// 执行浏览器操作
	err := chromedp.Run(global.SvcCtx.ChromeCtx,
		chromedp.Navigate("https://testnet.nad.fun/tokens/"+quickTradeRequest.TargetAbstract),
		chromedp.WaitEnabled("body > div.flex.min-h-dvh.w-full.flex-col > main > div > div > div.relative.mx-auto.w-full.px-0.lg\\:px-8.pb-\\[60px\\].lg\\:pb-\\[80px\\].max-w-\\[1800px\\] > div > main > div.w-full.px-0.sm\\:px-4.py-2.pb-20.lg\\:pb-2 > div > div.hidden.w-full.space-y-6.lg\\:block.lg\\:w-96 > div.my-0.mt-4.max-h-\\[440px\\].rounded-lg.bg-gray-800.p-4.sm\\:mt-0 > div.space-y-1 > div.flex.gap-2 > button:nth-child(2)", chromedp.ByQueryAll),
		chromedp.Click("body > div.flex.min-h-dvh.w-full.flex-col > main > div > div > div.relative.mx-auto.w-full.px-0.lg\\:px-8.pb-\\[60px\\].lg\\:pb-\\[80px\\].max-w-\\[1800px\\] > div > main > div.w-full.px-0.sm\\:px-4.py-2.pb-20.lg\\:pb-2 > div > div.hidden.w-full.space-y-6.lg\\:block.lg\\:w-96 > div.my-0.mt-4.max-h-\\[440px\\].rounded-lg.bg-gray-800.p-4.sm\\:mt-0 > div.space-y-1 > div.flex.gap-2 > button:nth-child(2)", chromedp.ByQueryAll),
		chromedp.Click("body > div.flex.min-h-dvh.w-full.flex-col > main > div > div > div.relative.mx-auto.w-full.px-0.lg\\:px-8.pb-\\[60px\\].lg\\:pb-\\[80px\\].max-w-\\[1800px\\] > div > main > div.w-full.px-0.sm\\:px-4.py-2.pb-20.lg\\:pb-2 > div > div.hidden.w-full.space-y-6.lg\\:block.lg\\:w-96 > div.my-0.mt-4.max-h-\\[440px\\].rounded-lg.bg-gray-800.p-4.sm\\:mt-0 > button", chromedp.ByQueryAll),
	)
	if err != nil {
		logx.Errorf("浏览器操作失败: %v", err)
		return nil, err
	}

	// 完成
	return &protocol.CallToolResult{
		Content: []protocol.Content{
			protocol.TextContent{
				Type: "text",
				Text: "请点击钱包插件确认交易",
			},
		},
	}, nil
}
