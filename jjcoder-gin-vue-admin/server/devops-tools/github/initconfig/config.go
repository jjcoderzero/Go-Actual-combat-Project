package initconfig

import (
	"context"
	"fmt"
	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
	"server/config"
)

func InitConfig() (client *github.Client) {
	token := config.CrcConfig.Github.AccessToken
	fmt.Println(token)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)
	return client
}
