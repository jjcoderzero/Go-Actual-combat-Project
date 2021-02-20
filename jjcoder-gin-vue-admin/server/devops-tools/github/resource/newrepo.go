package resource

import (
	"context"
	"github.com/google/go-github/v32/github"
	"log"
	"server/devops-tools/github/initconfig"
)

func NewRepo(name *string, private *bool, description *string) (repo *github.Repository) {
	client := initconfig.InitConfig()
	ctx := context.Background()
	r := &github.Repository{
		Name:        name,
		Private:     private,
		Description: description,
	}
	repo, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}
