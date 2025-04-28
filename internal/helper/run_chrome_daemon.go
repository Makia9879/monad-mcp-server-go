package helper

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Makia9879/monad-mcp-server-go/internal/global"
	"github.com/chromedp/chromedp"
	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/logx"
)

func RunChromeDaemon(ctx context.Context, pumpURL string) {
	if global.SvcCtx.ChromeCtx != nil {
		return
	}

	userDataDir := filepath.Join(global.SvcCtx.RunningPathCtx, "meme_pump")
	_, err := os.Stat(userDataDir)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(userDataDir, 0755)
		} else {
			logx.Errorf("创建用户文件夹错误：%v", err)
			return
		}
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserDataDir(userDataDir),
		chromedp.Flag("headless", false),
		chromedp.Flag("proxy-server", "http://127.0.0.1:7890"),
		// 新增JS执行相关参数
		//chromedp.Flag("disable-web-security", true),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("enable-automation", false),
	)

	// 修改allocCtx为全局可访问
	chromeAllocCtx, allocCancel := chromedp.NewExecAllocator(contextx.ValueOnlyFrom(ctx), opts...)

	// 创建可复用的浏览器上下文
	global.SvcCtx.ChromeCtx, global.SvcCtx.ChromeCancelFunc = chromedp.NewContext(chromeAllocCtx,
		chromedp.WithLogf(log.Printf))

	// 保持浏览器常驻
	go func() {
		defer allocCancel()
		err := chromedp.Run(global.SvcCtx.ChromeCtx, chromedp.Tasks{
			chromedp.Navigate(pumpURL),
			// 添加页面加载完成检测
			chromedp.WaitReady("body", chromedp.ByQuery),
			chromedp.ActionFunc(func(ctx context.Context) error {
				// 等待JavaScript初始化完成
				time.Sleep(2 * time.Second)
				return nil
			}),
		})
		if err != nil {
			logx.Errorf("浏览器启动失败%v", err)
			global.SvcCtx.ChromeCancelFunc()
			global.SvcCtx.ChromeCtx = nil
			return
		}

		for {
			select {
			case <-global.SvcCtx.ChromeCtx.Done():
				return
			default:
				logx.Infof("浏览器保持活跃")
				// 保持连接，每5分钟刷新一次
				time.Sleep(5 * time.Minute)
			}
		}
	}()
}

// GetBrowserContext 新增获取浏览器上下文的公共方法
func GetBrowserContext() (context.Context, context.CancelFunc) {
	return global.SvcCtx.ChromeCtx, global.SvcCtx.ChromeCancelFunc
}
