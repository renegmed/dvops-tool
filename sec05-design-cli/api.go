package pork

import (
	nap "devops_tools/sec03-nap/cmd"

	"github.com/spf13/viper"
)

var api *nap.API

// GitHubAPI is a singleton design pattern
func GitHubAPI() *nap.API {
	if api == nil {
		api = nap.NewAPI("https://api.github.com") // this is the base url for api
		token := viper.GetString("token")          // retrieve token field from the config file pork.yaml
		api.SetAuth(nap.NewAuthToken(token))
		api.AddResource("fork", GetForkResource())
	}
	return api
}
