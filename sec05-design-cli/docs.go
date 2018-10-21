package pork

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var DocsCmd = &cobra.Command{
	Use:   "docs",
	Short: "read the documentation for a repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("You must supply repository argument")
		}
		content := GetRepositoryReadme(args[0])
		fmt.Println(content)
	},
}

func GetRepositoryReadme(repository string) string {
	return "docs - " + repository
}

/*
$ pork -h
Project Forking Tool for Github

Usage:
  pork [command]

Available Commands:
  help        Help about any command
  search      search for GitHub repositories by keyword

Flags:
  -h, --help   help for pork

Use "pork [command] --help" for more information about a command.

*/
