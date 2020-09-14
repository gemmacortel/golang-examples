package main

import (
	"github.com/CodelyTV/golang-examples/02-refactor-to-cobra/internal/cli"
	"github.com/CodelyTV/golang-examples/02-refactor-to-cobra/internal/storage/ontario"
	"github.com/spf13/cobra"
)

func main() {
	r := ontario.NewOntarioRepository()
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(r))
	rootCmd.AddCommand(cli.InitStoresCmd())
	rootCmd.AddCommand(cli.InitPricesCmd())
	rootCmd.Execute()
}
