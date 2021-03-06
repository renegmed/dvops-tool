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

	$ pork search topic:go

	$ pork clone avelino/awesome-go
*/

import (
	"devops_tools/sec05-design-cli"
	"fmt"
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
	viper.SetConfigName("pork") // refers to ./pork.yaml file
	viper.AddConfigPath(".")    // refers to ./pork.yaml file
	viper.ReadInConfig()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No configuration file found")
	}
	viper.SetDefault("location", os.Getenv("HOME")) // see pork.yaml file for 'location' setup. If none, this set to user's home directory
}
