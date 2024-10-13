package product

import "net/url"

type Product struct {
	Name string
	URL  url.URL
}

func NewProduct(name string, url url.URL) Product {
	return Product{
		Name: name,
		URL:  url,
	}
}

func (p Product) String() string {
	return p.Name + "\n" + p.URL.String()
}
