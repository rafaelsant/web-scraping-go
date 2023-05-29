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
	var description []string
	var metadata []string
	GetInformation(url, &response, &images, &description, &metadata, ctx)
	println(metadata[0])
}

func GetInformation(url string, response *[]string, images *[]string, description *[]string, metadata *[]string, ctx context.Context) {
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Evaluate(`
		var a = []
		document.querySelectorAll(".lister-item-header").forEach(i => a.push(i.innerText))
		a`, &response), chromedp.Evaluate(`
		var a = []
		document.querySelectorAll("div > div > div > div > a > img").forEach(i => a.push(i.src))
		a`, &images), chromedp.Evaluate(`
		var a = []
		document.querySelectorAll(".list-description > p").forEach(i => a.push(i.innerText))
		a`, &description),
		chromedp.Evaluate(`
		var a = []
		document.querySelectorAll("div > div > div > .text-muted").forEach(y => {if(y.innerText.startsWith('Director')){a.push(y.innerText)}})
		a`, &metadata))
	if err != nil {
		log.Fatalf("error while reading %v", err)
	}
}
