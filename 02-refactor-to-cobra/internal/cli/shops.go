package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var stores = map[string]string{
	"01DC9ZAPGKEQJS4P4A48EG3P43": "Mercadona",
	"01DC9ZB23EW0J0ARAER09SJDKC": "Carrefour",
	"01DC9ZB89V1PQD977ZE6QXSQHH": "Alcampo",
}

func InitStoresCmd() *cobra.Command{
	shopsCmd := &cobra.Command{
		Use:   "stores",
		Short: "Where they are sold",
		Run:   runStoresFn(),
	}

	shopsCmd.Flags().StringP("id", "i", "", "Beer id")

	return shopsCmd
}

func runStoresFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")

		if id != "" {
			fmt.Println(stores[id])
		} else {
			fmt.Println(stores)
		}
	}
}
