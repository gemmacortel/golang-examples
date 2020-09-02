package cli

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func InitPricesCmd() *cobra.Command {
	pricesCmd := &cobra.Command{
		Use:   "prices",
		Short: "Find price",
		Run:   runPricesFn(),
	}

	pricesCmd.Flags().StringP("id", "i", "", "Find price")
	pricesCmd.Flags().Float64P("price", "p", 0, "Find beers below a certain price")

	return pricesCmd
}

func runPricesFn() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		prices := readCsv()

		id, _ := cmd.Flags().GetString("id")
		price, _ := cmd.Flags().GetFloat64("price")

		if id != "" {
			fmt.Println(prices[id])
		} else if price != 0 {
			for id, p := range prices {
				if p <= price {
					fmt.Println(id, p)
				}
			}
		} else {
			fmt.Println(prices)
		}
	}
}

func readCsv() map[string]float64 {
	f, _ := os.Open("../../data/beers.csv")
	reader := csv.NewReader(f)
	prices := make(map[string]float64)

	for line, _ := reader.Read(); line != nil; line, _ = reader.Read() {
		beerPrice, _ := strconv.ParseFloat(line[3], 8)
		prices[line[0]] = beerPrice
	}
	return prices
}
