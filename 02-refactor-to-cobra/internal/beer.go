package internal

import "fmt"

type Beer struct {
	Id       int
	Name     string
	Category string
	Price    float64
	BeerType *BeerType
	Brewer   string
	Country  string
}

func NewBeer(id int, name string, category string, price float64, beerType string, brewer string, country string) (b Beer) {
	b = Beer{
		id,
		name,
		category,
		price,
		NewBeerType(beerType),
		brewer,
		country,
	}

	return
}

func (b Beer) String() string {
	return fmt.Sprintf("id: %v, name: %v, category: %v, price: %v, beerType: %v, brewer: %v, country: %v",
		b.Id, b.Name, b.Category, b.Price, b.BeerType, b.Brewer, b.Country)
}
