package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func complex() {
	var x bool
	if x == true {
	}

}

const ITEM_SELECTOR = "div[data-product]"
const TITLE_SELECTOR = "h3.searchProductTitle"
const PRICE_SELECTOR = "div.searchTilePriceMobile .price-new"
const IMG_SELECTOR = "div.searchProductImage img"

func main() {
	c := colly.NewCollector()
	colly.AllowedDomains("gamestop.ca")
	complex()

	c.OnHTML(ITEM_SELECTOR, func(h *colly.HTMLElement) {

		item := Item{
			Name:   h.ChildText(TITLE_SELECTOR),
			Price:  h.ChildText(PRICE_SELECTOR),
			ImgUrl: h.ChildAttr(IMG_SELECTOR, "data-llsrc"),
		}

		fmt.Println(item)
		fmt.Println("allo")
	})

	c.Visit("https://www.gamestop.ca/SearchResult/QuickSearch?productType=2&platform=405&variantType=1")
}
