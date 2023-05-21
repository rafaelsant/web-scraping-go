package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	url := "https://www.imdb.com/list/ls569932833/?ref_=hm_edcft_ft_lst_csesmg23_1_i"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var response []string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Evaluate(`
		var a = []
		document.querySelectorAll(".lister-item-header").forEach(i => a.push(i.innerText))
		a`, &response))
	if err != nil {
		log.Fatalf("error while reading %v", err)
	}
	println(response)
}
