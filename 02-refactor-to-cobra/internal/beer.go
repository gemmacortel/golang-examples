package internal

import "fmt"

type Beer struct {
	Id       int       `json:"product_id"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
	Price    string    `json:"price"`
	BeerType *BeerType `json:"type"`
	Brewer   string    `json:"brewer"`
	Country  string    `json:"country"`
}

func NewBeer(id int, name string, category string, price string, beerType string, brewer string, country string) (b Beer) {
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

type BeersRepository interface {
	GetBeers() ([]Beer, error)
}
