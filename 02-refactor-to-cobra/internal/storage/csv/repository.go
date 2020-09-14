package csv

import (
	"bufio"
	beer "github.com/CodelyTV/golang-examples/02-refactor-to-cobra/internal"
	"os"
	"strconv"
	"strings"
)

type repository struct{}

func NewRepository() beer.BeersRepository {
	return &repository{}
}

func (r repository) GetBeers() ([]beer.Beer, error) {
	f, _ := os.Open("../../data/beers.csv")
	reader := bufio.NewReader(f)
	beers := make([]beer.Beer, 3)

	for line := readLine(reader); line != nil; line = readLine(reader) {
		values := strings.Split(string(line), ",")

		beerId, _ := strconv.Atoi(values[0])
		beerPrice, _ := strconv.ParseFloat(values[3], 8)
		beers = append(beers, beer.NewBeer(beerId, values[1], values[2], beerPrice, values[4], values[5], values[6]))
	}

	return beers, nil
}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}
