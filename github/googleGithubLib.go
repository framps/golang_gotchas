package main

// Use google git-hub to retrieve public repositories of an organization
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de
//
// See github.com/framps/golang_gotchas for latest code

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {

	org := flag.String("o", "", "Organization")
	token := flag.String("t", "", "github token")
	flag.Parse()

	var orgSet, tokenSet bool

	flag.Visit(func(arg *flag.Flag) {
		if arg.Name == "o" {
			orgSet = true
		}
		if arg.Name == "t" {
			tokenSet = true
		}
	})

	if !orgSet || !tokenSet {
		fmt.Printf("Missing -o and/or -p\n")
		os.Exit(1)
	}

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	client.BaseURL = &url.URL{Scheme: "https", Host: "api.github.com"}
	client.UserAgent = "framp@linux-tips-andtricks.de"

	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, rsp, err := client.Repositories.ListByOrg(ctx, *org, opt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Github Limit: %d - Remaining: %d\n\n", rsp.Limit, rsp.Remaining)
	fmt.Printf("Found %d repositories\n\n", len(repos))

	for _, r := range repos {
		var description string
		if r.Description != nil {
			description = *r.Description
		}
		fmt.Printf("Name: %s - CreatedAt: %s - UpdatedAt: %s - Description: %s\n", *r.Name, r.CreatedAt, r.UpdatedAt, description)
	}
}
