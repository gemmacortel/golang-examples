package internal

import (
	"encoding/json"
)

type BeerType int

const (
	Unknown BeerType = iota
	Lager
	Malt
	Ale
	NonAlcoholic
	Stout
	FlavouredMalt
)

var toId = map[string]BeerType{
	"Unknown":        Unknown,
	"Lager":          Lager,
	"Malt":           Malt,
	"Ale":            Ale,
	"NonAlcoholic":   NonAlcoholic,
	"Stout":          Stout,
	"Flavoured Malt": FlavouredMalt,
}

var toString = map[BeerType]string{
	Unknown:       "Unknown",
	Lager:         "Lager",
	Malt:          "Malt",
	Ale:           "Ale",
	NonAlcoholic:  "NonAlcoholic",
	Stout:         "Stout",
	FlavouredMalt: "Flavoured Malt",
}

func NewBeerType(t string) *BeerType {
	beerType := toId[t]
	return &beerType
}

func (t BeerType) String() string {
	return toString[t]
}

func (t *BeerType) UnmarshalJSON(b []byte) error {
	var beerTypeName string
	e := json.Unmarshal(b, &beerTypeName)
	if nil != e {
		return e
	}

	*t = toId[beerTypeName]

	return nil
}
