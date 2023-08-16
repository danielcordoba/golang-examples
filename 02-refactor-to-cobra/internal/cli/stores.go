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

func InitStoresCmd() *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "stores",
		Short: "Print data about stores",
		Run:   runStoresFn(),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id of the beer")

	return beersCmd
}

func runStoresFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		//obtenemos el id que viene por línea de comando
		id, _ := cmd.Flags().GetString(idFlag)

		if stores[id] == "" {
			fmt.Println("No stores found")
		} else {
			fmt.Println(stores[id])
		}

	}
}
