// Command visible is a chromedp example demonstrating how to wait until an
// element is visible.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func main() {
	port := flag.Int("port", 8544, "port")
	flag.Parse()

	dir, err := os.MkdirTemp("", "chromedp-example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserDataDir(dir),
		chromedp.Flag("headless", false),
	)

	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer allocCancel()

	taskCtx, taskCtxCancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer taskCtxCancel()

	// run server
	go func() {
		time.Sleep(3 * time.Second)
		err = chromedp.Run(taskCtx, visible(fmt.Sprintf("http://localhost:%d", *port)))
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = testServer(fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
}

func visible(host string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(host),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(makeVisibleScript).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf("waiting 3s for box to become visible")
			return nil
		}),
		chromedp.WaitVisible(`#box1`),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf(">>>>>>>>>>>>>>>>>>>> BOX1 IS VISIBLE")
			return nil
		}),
		chromedp.WaitVisible(`#box2`),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf(">>>>>>>>>>>>>>>>>>>> BOX2 IS VISIBLE")
			return nil
		}),
	}
}

const (
	makeVisibleScript = `setTimeout(function() {
	document.querySelector('#box1').style.display = '';
}, 3000);`
)

// testServer is a simple HTTP server that serves a static html page.
func testServer(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(res, indexHTML)
	})
	return http.ListenAndServe(addr, mux)
}

const indexHTML = `<!doctype html>
<html>
<head>
  <title>example</title>
</head>
<body>
  <div id="box1" style="display:none">
    <div id="box2">
      <p>box2</p>
    </div>
  </div>
  <div id="box3">
    <h2>box3</h3>
    <p id="box4">
      box4 text
      <input id="input1" value="some value"><br><br>
      <textarea id="textarea1" style="width:500px;height:400px">textarea</textarea><br><br>
      <input id="input2" type="submit" value="Next">
      <select id="select1">
        <option value="one">1</option>
        <option value="two">2</option>
        <option value="three">3</option>
        <option value="four">4</option>
      </select>
    </p>
  </div>
</body>
</html>`
