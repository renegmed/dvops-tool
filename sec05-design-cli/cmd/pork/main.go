package main

/*

	$ go install ./cmd/pork
	$ pork search topic:ruby
	$ port docs myrepository

*/

import (
	"devops_tools/sec05-design-cli"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func main() {
	rootCmd.Execute()
}

func init() {
	rootCmd = &cobra.Command{
		Use:   "pork",
		Short: "Project Forking Tool for Github",
	}
	rootCmd.AddCommand(pork.SearchCmd)
	rootCmd.AddCommand(pork.DocsCmd)
}
