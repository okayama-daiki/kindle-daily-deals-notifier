package crawler

import (
	"net/url"

	"github.com/gocolly/colly/v2"
	"github.com/okayama-daiki/kindle-daily-deals-notifier/libs/product"
)

func Crawl() ([]product.Product, error) {
	const (
		TARGET_URL = "https://www.amazon.co.jp/kindle-dbs/browse?widgetId=ebooks-deals-storefront_KindleDailyDealsStrategy&sourceType=recs"
	)

	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	productList := []product.Product{}

	c.OnHTML("div#browse-views-area ul > div#sponsoredLabel-title > a", func(e *colly.HTMLElement) {
		productName := e.Attr("aria-label")

		rawProductPath := e.Attr("href")
		productPath, err := url.Parse(rawProductPath)

		if productName == "" || err != nil {
			return
		}

		productUrl := url.URL{
			Scheme: "https",
			Host:   "www.amazon.co.jp",
			Path:   productPath.Path,
		}

		product := product.New(productName, productUrl)
		productList = append(productList, product)
	})

	err := c.Visit(TARGET_URL)
	if err != nil {
		return nil, err
	}

	return productList, nil
}
