package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

var beers = map[string]string{
	"01D9X58E7NPXX5MVCR9QN794CH": "Mad Jack Mixer",
	"01D9X5BQ5X48XMMVZ2F2G3R5MS": "Keystone Ice",
	"01D9X5CVS1M9VR5ZD627XDF6ND": "Belgian Moon",
}

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
		println(id)

		if id != 0 {
			fmt.Println(beers[id])
		} else {
			fmt.Println(beers)
		}
	}
}

func readFile() map[int]string {
	f, _ := os.Open("../../data/beers.csv")
	reader := bufio.NewReader(f)
	beers := make(map[int]string)

	for line := readLine(reader); line != nil; line = readLine(reader) {
		values := strings.Split(string(line), ",")
		beerId, _ := strconv.Atoi(values[0])

		beers[beerId] = values[1]
	}
	return beers
}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}
