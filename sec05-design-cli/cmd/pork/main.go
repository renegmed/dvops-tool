package main

/*

	$ go install ./cmd/pork
	$ pork search topic:ruby
	$ pork docs myrepository

	$ pork clone devops_tools/sec05-design-cli
	$ pork clone devops_tools/sec05-design-cli --ref develop --create
	$ pork clone --help

	$ pork clone mspaulding06/nap --ref testing --create

	$ pork fork myrepository


*/

import (
	"devops_tools/sec05-design-cli"
	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func main() {
	rootCmd.Execute()
}

func init() {
	// Cobra is both a library for creating powerful modern CLI
	// applications as well as a program to generate applications
	// and command files
	rootCmd = &cobra.Command{
		Use:   "pork",
		Short: "Project Forking Tool for Github",
	}
	rootCmd.AddCommand(pork.SearchCmd)
	rootCmd.AddCommand(pork.DocsCmd)
	rootCmd.AddCommand(pork.CloneCmd)
	rootCmd.AddCommand(pork.ForkCmd)

	// Viper is a complete configuration solution for Go applications.
	// When building a modern application, you don’t need to worry
	// about configuration file formats; you want to focus on building
	// awesome software. Viper is here to help with that.
	viper.SetDefault("location", os.Getenv("HOME"))
	viper.SetConfigName("pork")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
}
