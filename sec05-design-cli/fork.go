package pork

import (
	nap "devops_tools/sec03-nap/cmd"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

type ForkResponse struct {
	CloneURL string `json:"clone_url"`
	FullName string `json:"full_name"`
}

var ForkCmd = &cobra.Command{
	Use:   "fork",
	Short: "fork a GitHub repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("You must supply a repository")
		}
		if err := ForkRepository(args[0]); err != nil {
			log.Fatalln("Unable to fork repository:", err)
		}
	},
}

func ForkRepository(repository string) error {
	values := strings.Split(repository, "/")
	if len(values) != 2 {
		return fmt.Errorf("Repository must be in the format owner/project")
	}

	return GitHubAPI().Call("fork", map[string]string{
		"owner": values[0],
		"repo":  values[1],
	})
}

func ForkSuccess(resp *http.Response, _ interface{}) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respContent := ForkResponse{}
	json.Unmarshal(content, &respContent)
	fmt.Printf("\t[ForkSuccess]Forked to repository: %s\n", respContent.FullName)
	return nil
}

// GetForkResource is a Builder function
func GetForkResource() *nap.RestResource {
	forkRouter := nap.NewRouter()
	forkRouter.RegisterFunc(202, ForkSuccess) // 202 - resource was created
	forkRouter.RegisterFunc(401, func(_ *http.Response, _ interface{}) error {
		return fmt.Errorf("You must set an authentication token")
	})
	fork := nap.NewResource("/repos/{{.owner}}/{{.repo}}/forks", "POST", forkRouter)
	return fork
}
