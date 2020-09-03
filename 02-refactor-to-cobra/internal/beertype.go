package internal

type BeerType int

const (
	Unknown BeerType = iota
	Larger
	Malt
	Ale
	NonAlcoholic
)

var toId = map[string]BeerType{
	"Unknown":      Unknown,
	"Larger":       Larger,
	"Malt":         Malt,
	"Ale":          Ale,
	"NonAlcoholic": NonAlcoholic,
}

var toString = map[BeerType]string{
	Unknown:      "Unknown",
	Larger:       "Larger",
	Malt:         "Malt",
	Ale:          "Ale",
	NonAlcoholic: "NonAlcoholic",
}

func NewBeerType(t string) *BeerType {
	beerType := toId[t]
	return &beerType
}

func (b BeerType) String() string {
	return toString[b]
}
