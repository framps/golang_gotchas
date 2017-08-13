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
	"os/exec"
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

// GithubClient - Simple github client
type GithubClient struct {
	APIUrl      string
	Client      *http.Client
	AccessToken string
	UserAgent   string
}

// Timestamp - CreatedAt et al
type Timestamp struct {
	time.Time
}

func (t Timestamp) String() string {
	return t.Time.String()
}

// UnmarshalJSON - Unmarshal Timestamp
func (t *Timestamp) UnmarshalJSON(data []byte) (err error) {
	dataAsString := string(data)
	fmt.Printf("--- %s\n", dataAsString)
	timeFromInt, err := strconv.ParseInt(dataAsString, 10, 64)
	if err == nil {
		(*t).Time = time.Unix(timeFromInt, 0)
	} else {
		(*t).Time, err = time.Parse(`"`+time.RFC3339+`"`, dataAsString)
	}
	return
}

// Rate - Rate header returned by git
type Rate struct {
	Limit     int       `json:"limit"`
	Remaining int       `json:"remaining"`
	Reset     time.Time `json:"reset"`
}

// Repository - Git repository
type Repository struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp `json:"updated_at,omitempty"`
}

// Handle - Handle errors and panic
func Handle(err error) {
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
		Handle(err)
	}
	if remaining := r.Header.Get(headerRateLimitRemaining); remaining != "" {
		rate.Remaining, err = strconv.Atoi(remaining)
		Handle(err)
	}
	if reset := r.Header.Get(headerRateLimitReset); reset != "" {
		if v, _ := strconv.ParseInt(reset, 10, 64); v != 0 {
			rate.Reset = time.Unix(v, 0)
		}
	}
	return rate
}

// NewGithubClient - Simple github client
func NewGithubClient(apiURL string, client *http.Client, accessToken, userAgent string) *GithubClient {
	if client == nil {
		client = &http.Client{}
	}
	return &GithubClient{APIUrl: apiURL, Client: client, AccessToken: accessToken, UserAgent: userAgent}
}

func (r GithubClient) executeRequest(url string, additionalHeaderParms ...headerParms) (*[]byte, Rate, error) {

	fmt.Printf("Executing GET %s\n", url)
	req, err := http.NewRequest("GET", url, nil)
	Handle(err)
	req.Header.Set("Authorization", "token "+r.AccessToken)
	req.Header.Set("User-Agent", r.UserAgent)

	if len(additionalHeaderParms) == 1 {
		for k, v := range additionalHeaderParms[0] {
			req.Header.Set(k, v)
		}
	}

	res, err := r.Client.Do(req)
	Handle(err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, Rate{}, fmt.Errorf("Http error %d received", res.StatusCode)
	}

	rsp, err := ioutil.ReadAll(res.Body)
	Handle(err)

	rate := retrieveRate(res)

	return &rsp, rate, nil
}

// GetReadme - Retrieve readme of a repo
func (r *GithubClient) GetReadme(org string, repository string) (*[]byte, Rate, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/readme", gitHost, org, repository)
	addtlHeaderParms := headerParms{"Accept": "application/vnd.github.html"}
	result, rate, err := r.executeRequest(url, addtlHeaderParms)
	Handle(err)
	return result, rate, nil
}

// GetRepositoriesOfOrg - retrieve all repositories of an org
func (r *GithubClient) GetRepositoriesOfOrg(org string, repositoryType string) ([]Repository, Rate, error) {
	url := fmt.Sprintf("%s/orgs/%s/repos", gitHost, org)
	requestResult, rate, err := r.executeRequest(url)
	Handle(err)

	fmt.Printf("Repo response %s\n", string(*requestResult))

	var result []Repository

	json.Unmarshal(*requestResult, &result)
	return result, rate, nil
}

func main() {

	org := flag.String("o", "", "Organization")
	token := flag.String("t", "", "github token")
	repo := flag.String("r", "", "Repository to retrieve the readme")
	userAgent := flag.String("u", "I'm a test user", "github user")

	flag.Parse()

	var orgSet, tokenSet, repoSet bool

	flag.Visit(func(arg *flag.Flag) {
		if arg.Name == "o" {
			orgSet = true
		}
		if arg.Name == "t" {
			tokenSet = true
		}
		if arg.Name == "r" {
			repoSet = true
		}
	})

	if !orgSet || !tokenSet {
		fmt.Printf("Missing -o and/or -p\n")
		os.Exit(1)
	}

	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	repoClient := NewGithubClient(gitHost, client, *token, *userAgent)

	repos, rate, err := repoClient.GetRepositoriesOfOrg(*org, "public")
	Handle(err)

	fmt.Printf("Rate: %v\n", rate)
	fmt.Printf("Repos: %#v\n", repos)

	if repoSet {
		readme, rate, err := repoClient.GetReadme(*org, *repo)
		Handle(err)

		err = ioutil.WriteFile("readme.html", *readme, 0644)
		Handle(err)

		fmt.Printf("Rate: %v\n", rate)

		cmd := exec.Command("/usr/bin/firefox", "readme.html")
		err = cmd.Start()
		Handle(err)
	}
}
