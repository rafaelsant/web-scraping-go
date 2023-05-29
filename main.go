package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
)

type Movie struct {
	Name             string
	Image            string
	Description      string
	MovieInformation string
}

func main() {
	url := "https://www.imdb.com/list/ls569932833/?ref_=hm_edcft_ft_lst_csesmg23_1_i"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	movies := GetMovies(url, ctx)
	fmt.Println(movies[0])
}

func GetMovies(url string, ctx context.Context) []Movie {
	var titles []string
	var images []string
	var description []string
	var information []string
	movies := make([]Movie, 100)
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Evaluate(`
		var a = []
		document.querySelectorAll(".lister-item-header").forEach(i => a.push(i.innerText))
		a`, &titles), chromedp.Evaluate(`
		var a = []
		document.querySelectorAll("div > div > div > div > a > img").forEach(i => a.push(i.src))
		a`, &images), chromedp.Evaluate(`
		var a = []
		document.querySelectorAll(".list-description > p").forEach(i => a.push(i.innerText))
		a`, &description),
		chromedp.Evaluate(`
		var a = []
		document.querySelectorAll("div > div > div > .text-muted").forEach(y => {if(y.innerText.startsWith('Director')){a.push(y.innerText)}})
		a`, &information))
	if err != nil {
		log.Fatalf("error while reading %v", err)
	}

	for num, title := range titles {
		movies[num] = Movie{Name: title, Image: images[num], Description: description[num+1], MovieInformation: information[num]}
	}
	return movies
}
