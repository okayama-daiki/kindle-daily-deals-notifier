package crawler

import (
	"net/url"

	"github.com/gocolly/colly/v2"
	"github.com/okayama-daiki/kindle-daily-deals-notifier/libs/product"
)

const (
	targetURL = "https://www.amazon.co.jp/kindle-dbs/browse?widgetId=ebooks-deals-storefront_KindleDailyDealsStrategy&sourceType=recs"
)

func Crawl() ([]product.Product, error) {

	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	products := []product.Product{}

	c.OnHTML("div#browse-views-area ul > div#sponsoredLabel-title > a", func(e *colly.HTMLElement) {
		productName := e.Attr("aria-label")
		if productName == "" {
			return
		}

		rawProductPath := e.Attr("href")
		productPath, err := url.Parse(rawProductPath)
		if err != nil {
			return
		}

		productUrl := url.URL{
			Scheme: "https",
			Host:   "www.amazon.co.jp",
			Path:   productPath.Path,
		}

		product := product.New(productName, productUrl)
		products = append(products, product)
	})

	if err := c.Visit(targetURL); err != nil {
		return nil, err
	}

	return products, nil
}
