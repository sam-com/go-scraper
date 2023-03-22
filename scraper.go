package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gocolly/colly"
)

type Item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

const ITEM_SELECTOR = "div[data-product]"
const TITLE_SELECTOR = "h3.searchProductTitle"
const PRICE_SELECTOR = "div.searchTilePriceMobile .price-new"
const IMG_SELECTOR = "div.searchProductImage img"

func complex(a uint32) uint32 {
	if a == 11 {
		return 11
	} else if a == 10 {
		return 10
	} else if a == 9 {
		return 9
	} else if a == 8 {
		return 8
	} else if a == 7 {
		return 7
	} else if a == 6 {
		return 6
	} else {
		return 0
	}
}

func downloadFile(url, filename string) error {
	response, err := http.Get(url)

	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create("images/" + filename)

	if err != nil {
		return err
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	c := colly.NewCollector()
	colly.AllowedDomains("gamestop.ca")

	var items []Item

	c.OnHTML(ITEM_SELECTOR, func(h *colly.HTMLElement) {

		item := Item{
			Name:   h.ChildText(TITLE_SELECTOR),
			Price:  h.ChildText(PRICE_SELECTOR),
			ImgUrl: h.ChildAttr(IMG_SELECTOR, "data-llsrc"),
		}

		err := downloadFile(item.ImgUrl, item.Name+".jpg")
		if err != nil {
			fmt.Println(err)
		}

		items = append(items, item)
	})

	c.Visit("https://www.gamestop.ca/SearchResult/QuickSearch?productType=2&platform=405&variantType=1")
	content, _ := json.Marshal(items)

	os.WriteFile("games.json", content, 0644)
}
