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

func complex(a uint32) {
	b := 3
	c := 4
	d := 3
	e := 4
	f := 3
	g := 4
	h := 3
	i := 4

	if a == 10 {
		fmt.Println(a)
	} else if a == 9 {
		fmt.Println(a)
	} else if a == 8 {
		fmt.Println(a)
	} else if a == 7 {
		fmt.Println(a)
	} else if a == 6 {
		fmt.Println(a)
	} else if a == 5 {
		fmt.Println(a)
	} else if a == 4 {
		fmt.Println(a)
	} else if a == 3 {
		fmt.Println(a)
	} else if a == 2 {
		fmt.Println(a)
	} else if a == 1 {
		fmt.Println(a)
	}

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

}

const ITEM_SELECTOR = "div[data-product]"
const TITLE_SELECTOR = "h3.searchProductTitle"
const PRICE_SELECTOR = "div.searchTilePriceMobile .price-new"
const IMG_SELECTOR = "div.searchProductImage img"

func main() {
	c := colly.NewCollector()
	colly.AllowedDomains("gamestop.ca")
	complex(1)

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
