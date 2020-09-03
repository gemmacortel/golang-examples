package cli

import (
	"bufio"
	"fmt"
	beer "github.com/CodelyTV/golang-examples/02-refactor-to-cobra/internal"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"

// InitBeersCmd initialize beers command
func InitBeersCmd() *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id of the beer")

	return beersCmd
}

func runBeersFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		beers := readFile()

		i, _ := cmd.Flags().GetString(idFlag)
		id, _ := strconv.Atoi(i)

		if id != 0 {
			fmt.Println(beers[id].String())
		} else {
			fmt.Println(beers)
		}
	}
}

func readFile() map[int]beer.Beer {
	f, _ := os.Open("../../data/beers.csv")
	reader := bufio.NewReader(f)
	beers := make(map[int]beer.Beer)

	for line := readLine(reader); line != nil; line = readLine(reader) {
		values := strings.Split(string(line), ",")

		beerId, _ := strconv.Atoi(values[0])
		beerPrice, _ := strconv.ParseFloat(values[3], 8)
		beers[beerId] = beer.NewBeer(beerId, values[1], values[2], beerPrice, values[4], values[5], values[6])
	}

	return beers
}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}
