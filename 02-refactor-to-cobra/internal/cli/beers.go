package cli

import (
	"fmt"
	beer "github.com/CodelyTV/golang-examples/02-refactor-to-cobra/internal"
	"github.com/spf13/cobra"
	"strconv"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"

// InitBeersCmd initialize beers command
func InitBeersCmd(r beer.BeersRepository) *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(r),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id of the beer")

	return beersCmd
}

func runBeersFn(r beer.BeersRepository) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		beers, _ := r.GetBeers()

		i, _ := cmd.Flags().GetString(idFlag)
		id, _ := strconv.Atoi(i)

		for _, b := range beers {
			if b.Id == id {
				fmt.Println(b)
				return
			}
		}
		fmt.Println(beers)
		return
	}
}
