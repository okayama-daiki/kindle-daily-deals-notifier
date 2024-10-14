package product

import "net/url"

type Product struct {
	name string
	url  url.URL
}

func New(name string, url url.URL) Product {
	return Product{
		name,
		url,
	}
}

func (p Product) String() string {
	return p.name + "\n" + p.url.String()
}
