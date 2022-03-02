package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v43/github"
	"github.com/jessevdk/go-flags"
	"golang.org/x/oauth2"
)

var Options struct {
	Verbose     bool   `short:"v" long:"verbose" description:"show verbose" required:"true"`
	AccessToken string `short:"a" long:"access-token" description:"access token" required:"true"`
	Owner       string `short:"o" long:"owner" description:"owner name" required:"true"`
	Repo        string `short:"r" long:"repo" description:"repo name" required:"true"`
	Ref         string `short:"ref" long:"ref" description:"ref" required:"true"`
	Environment string `short:"e" long:"environment" description:"environment name" required:"true"`
	Description string `short:"d" long:"description" description:"description" required:"true"`
	TargetUrl   string `short:"t" long:"target-url" description:"target url" required:"true"`
}

func main() {
	_, err := flags.ParseArgs(&Options, os.Args)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: Options.AccessToken,
		})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	autoMerge := false
	var requiredContexts []string
	deployment, _, err := client.Repositories.CreateDeployment(ctx, Options.Owner, Options.Repo, &github.DeploymentRequest{
		Ref:              &Options.Ref,
		AutoMerge:        &autoMerge,
		RequiredContexts: &requiredContexts,
		Environment:      &Options.Environment,
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deployment url %+v\n", deployment)
	success := "success"

	_, _, err = client.Repositories.CreateDeploymentStatus(ctx, Options.Owner, Options.Repo, *deployment.ID, &github.DeploymentStatusRequest{
		State:          &success,
		Description:    &Options.Description,
		EnvironmentURL: &Options.TargetUrl,
	})

	if err != nil {
		log.Fatal(err)
	}
}
