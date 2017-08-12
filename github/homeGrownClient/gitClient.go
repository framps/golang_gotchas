package main

// retrieve public repositories of an organization using a home grown github client
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de
//
// See github.com/framps/golang_gotchas for latest code

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	gitHost = "https://api.github.com"
)

const (
	headerRateLimit          = "X-RateLimit-Limit"
	headerRateLimitRemaining = "X-RateLimit-Remaining"
	headerRateLimitReset     = "X-RateLimit-Reset"
)

type GitRepository struct {
	APIUrl      string
	Client      *http.Client
	AccessToken string
}

type Rate struct {
	Limit     int       `json:"limit"`
	Remaining int       `json:"remaining"`
	Reset     time.Time `json:"reset"`
}

type Repository struct {
	Name string `json:"name"`
}

func CheckForErrors(err error) {
	if err != nil {
		panic(err)
	}
}

type headerParms map[string]string

func (r Rate) String() string {
	return fmt.Sprintf("Limit: %d - Remaining: %d - Reset: %s", r.Limit, r.Remaining, r.Reset.String())
}

func retrieveRate(r *http.Response) Rate {
	var (
		rate Rate
		err  error
	)
	if limit := r.Header.Get(headerRateLimit); limit != "" {
		rate.Limit, err = strconv.Atoi(limit)
		CheckForErrors(err)
	}
	if remaining := r.Header.Get(headerRateLimitRemaining); remaining != "" {
		rate.Remaining, err = strconv.Atoi(remaining)
		CheckForErrors(err)
	}
	if reset := r.Header.Get(headerRateLimitReset); reset != "" {
		if v, _ := strconv.ParseInt(reset, 10, 64); v != 0 {
			rate.Reset = time.Unix(v, 0)
		}
	}
	return rate
}

func NewGitRepository(apiUrl string, client *http.Client, accessToken string) *GitRepository {
	if client == nil {
		client = &http.Client{}
	}
	return &GitRepository{APIUrl: apiUrl, Client: client, AccessToken: accessToken}
}

func (r GitRepository) ExecuteRequest(url string, additionalHeaderParms ...headerParms) (*[]byte, Rate, error) {

	req, err := http.NewRequest("GET", url, nil)
	CheckForErrors(err)
	req.Header.Add("Authorization", "token "+r.AccessToken)

	if len(additionalHeaderParms) == 1 {
		for k, v := range additionalHeaderParms[0] {
			req.Header.Add(k, v)
		}
	}

	res, err := r.Client.Do(req)
	CheckForErrors(err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, Rate{}, fmt.Errorf("Http error %d received", res.StatusCode)
	}

	rsp, err := ioutil.ReadAll(res.Body)
	CheckForErrors(err)

	rate := retrieveRate(res)

	return &rsp, rate, nil
}

func (r *GitRepository) GetReadme(org string, repository string) (*[]byte, Rate, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/readme", gitHost, org, repository)
	addtlHeaderParms := headerParms{"Accept": "application/vnd.github.html"}
	result, rate, err := r.ExecuteRequest(url, addtlHeaderParms)
	CheckForErrors(err)
	return result, rate, nil
}

func (r *GitRepository) GetRepositoriesOfOrg(org string, repositoryType string) ([]Repository, Rate, error) {
	url := fmt.Sprintf("%s/orgs/%s/repos", gitHost, org)
	requestResult, rate, err := r.ExecuteRequest(url)
	CheckForErrors(err)

	var result []Repository

	json.Unmarshal(*requestResult, &result)
	return result, rate, nil
}

func main() {

	org := flag.String("o", "", "Organization")
	token := flag.String("t", "", "github token")
	// repo := flag.String("r", "", "Repository to retrieve the readme")
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

	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	repoClient := NewGitRepository(gitHost, client, *token)

	repos, rate, err := repoClient.GetRepositoriesOfOrg(*org, "public")
	CheckForErrors(err)

	fmt.Printf("Rate: %v\n", rate)
	fmt.Printf("Repos: %v\n", repos)

	/*
		readme, rate, err := repoClient.GetReadme(*org, *repo)
		CheckForErrors(err)

		err = ioutil.WriteFile("readme.html", *readme, 0644)
		CheckForErrors(err)

		fmt.Printf("Rate: %v\n", rate)

		cmd := exec.Command("/usr/bin/firefox", "readme.html")
		err = cmd.Start()
		CheckForErrors(err)
	*/
}
