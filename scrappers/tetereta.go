package scrappers

import (
	urls "tea_webscraper/urls"

	"github.com/gocolly/colly"
	"golang.org/x/exp/slices"
)

type item struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

func Tetereta() int {
	var items []item
	c := colly.NewCollector(
		colly.AllowedDomains("tetereta.com"),
	)
	baseUrl := "https://tetereta.com/"

	for _, category := range urls.TeteretaMap {
		c.OnHTML("div.astra-shop-summary-wrap", func(h *colly.HTMLElement) {
			item := item{
				Name:  h.ChildText("h2.woocommerce-loop-product__title"),
				Price: h.ChildText("bdi"),
			}
			if !(slices.Contains(items, item)) {
				items = append(items, item)
			}
		})
		c.OnHTML("ul.page-numbers", func(h *colly.HTMLElement) {
			nuevaPagina := h.ChildAttr("a.next", "href")
			c.Visit(nuevaPagina)
		})
		c.Visit(baseUrl + category)
	}
	//fmt.Println(len(items))
	return len(items)
}
