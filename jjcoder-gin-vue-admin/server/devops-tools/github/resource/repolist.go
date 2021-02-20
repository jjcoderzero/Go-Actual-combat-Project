package resource

import (
	"context"
	"github.com/google/go-github/v32/github"
	"log"
	"server/devops-tools/github/initconfig"
)

func ListRepos() (repos []*github.Repository) {
	client := initconfig.InitConfig()
	ctx := context.Background()
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		log.Fatal(err)
	}
	return repos
}
