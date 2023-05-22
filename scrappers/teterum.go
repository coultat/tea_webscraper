package scrappers

// Código que hace webscrapping de teterum.com
// - falta añadir más categorías. Por ahora solo está la de collections
// - falta mandar las categorías por parámetros
// - falta devolver el json.
import (
	urls "tea_webscraper/urls"

	"github.com/gocolly/colly"
)

/*
type item struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}*/

func TeterumWebScraper() []item {
	c := colly.NewCollector(
		colly.AllowedDomains("teterum.com"),
	)
	baseUrl := "https://teterum.com"
	//var TeterumMap = [3]string{"/collections/te", "/collections/infusiones", "/collections/accesorios"}

	var items []item

	for _, categories := range urls.TeterumMap {

		c.OnHTML("div.mt-3", func(h *colly.HTMLElement) {
			if len(h.ChildText(".sf__pcard-name")) != 0 {
				item := item{
					Name:  h.ChildText(".sf__pcard-name"),
					Price: h.ChildText(".f-price-item--sale"),
				}
				items = append(items, item)

			}

		})
		c.OnHTML("span.next", func(h *colly.HTMLElement) {
			nuevaPagina := h.Request.AbsoluteURL(h.ChildAttr("a", "href"))
			c.Visit(nuevaPagina)
		})

		c.Visit(baseUrl + categories)

	}
	//fmt.Println(items)
	return items
}
