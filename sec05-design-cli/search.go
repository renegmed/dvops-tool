package pork

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
)

type SearchResponse struct {
	Results []*SearchResult `json:"items"`
}

type SearchResult struct {
	FullName string `json:"full_name"`
}

var SearchCmd = &cobra.Command{
	Use:   "search", // this is a sub-command
	Short: "search for GitHub repositories by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		if err := SearchByKeyword; err != nil {
			log.Fatalln("Search Failed:", err)
		}
	},
}

func SearchByKeyword(keywords []string) error {
	return GitHubAPI().Call("search", map[string]string{
		"query": strings.Join(keywords, "+"),
	})
}
