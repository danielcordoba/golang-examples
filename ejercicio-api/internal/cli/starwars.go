package cli

import (
	"fmt"
	"github.com/CodelyTV/golang-examples/ejercicio-api/internal/api"
	"github.com/CodelyTV/golang-examples/ejercicio-api/internal/parser"
	"github.com/spf13/cobra"
)

const idFlag = "id"
const csvFlag = "csv"

type cobraFn func(cmd *cobra.Command, args []string)

func InitStarWarsCommand(repository api.StarWarsRepo) *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "people",
		Short: "Downloads people from films",
		Run:   runStarWarsFn(repository),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id of the person")
	beersCmd.Flags().BoolP(csvFlag, "c", false, "saves as csv")

	return beersCmd
}

func runStarWarsFn(repository api.StarWarsRepo) cobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)
		csv, _ := cmd.Flags().GetBool(csvFlag)

		var people []api.Person
		var err error

		if id != "" {
			people, err = repository.GetPerson(id)
		} else {
			people, err = repository.GetPeople()
		}

		if err == nil {
			fmt.Println(people)
		}

		if csv {
			err = parser.SaveJson(people)
		}

		if err != nil {
			fmt.Print(err)
		}
	}
}
