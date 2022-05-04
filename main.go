package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			req, _ := client.NewRequest("GET", "https://charts.youtube.com", nil)
			req.Rendered = true
			req.Actions = []chromedp.Action{
				chromedp.Navigate("https://charts.youtube.com"),
				chromedp.WaitReady(":root"),
				chromedp.ActionFunc(func(ctx context.Context) error {
					node, err := dom.GetDocument().Do(ctx)
					if err != nil {
						return err
					}
					body, err := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
					fmt.Println("Helloooo", body)
					return err
				}),
			}
			g.Do(req, g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			fmt.Println(string(r.Body))
			fmt.Println(r.Request.URL.String(), r.Header)
		},
	}).Start()
}
