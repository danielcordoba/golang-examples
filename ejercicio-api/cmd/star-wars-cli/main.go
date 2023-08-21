package main

import (
	api "github.com/CodelyTV/golang-examples/ejercicio-api/internal/api"
	"github.com/CodelyTV/golang-examples/ejercicio-api/internal/cli"
	"github.com/spf13/cobra"
)

func main() {

	var repo api.StarWarsRepo
	repo = api.NewStarWarsRepo()

	rootCmd := &cobra.Command{Use: "starwars-cli"}
	rootCmd.AddCommand(cli.InitStarWarsCommand(repo))
	rootCmd.Execute()
}
