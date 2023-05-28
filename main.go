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
	var images []string
	GetInformation(url, &response, &images, ctx)
	println(images[0])
}

func GetInformation(url string, response *[]string, images *[]string, ctx context.Context) {
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Evaluate(`
		var a = []
		document.querySelectorAll(".lister-item-header").forEach(i => a.push(i.innerText))
		a`, &response), chromedp.Evaluate(`
		var a = []
		document.querySelectorAll("div > div > div > div > a > img").forEach(i => a.push(i.src))
		a`, &images))
	if err != nil {
		log.Fatalf("error while reading %v", err)
	}
}
