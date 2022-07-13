package ext_service

import (
	"context"
	"encoding/base64"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v45/github"
	"log"
	"net/http"
	"reacji-github-issue/config"
)

type GitHubService struct {
	client *github.Client
}

type IssueParam struct {
	Owner   string
	Repo    string
	Request github.IssueRequest
}

func NewGitHubService(sc config.SystemConfig) *GitHubService {

	tr := http.DefaultTransport

	pemBin, err := base64.StdEncoding.DecodeString(sc.GitHubPemBinary)
	if err != nil {
		log.Fatalf("Falied to decode GITHUB_PEM_BINARY %s", err)
	}

	itr, err := ghinstallation.New(tr, sc.GitHubAppID, sc.GitHubInstallationID, pemBin)
	return &GitHubService{client: github.NewClient(&http.Client{Transport: itr})}
}

// CreateIssue create GitHub issue and returns created issue url.
func (ghs GitHubService) CreateIssue(param IssueParam) (*string, error) {

	log.Println(*param.Request.Title)
	log.Println(*param.Request.Body)

	i, res, err := ghs.client.Issues.Create(context.Background(), param.Owner, param.Repo, &param.Request)

	log.Println(res.Status)
	if err != nil {
		log.Printf("failed to create issue %v", err)
		return nil, err
	}
	log.Printf("creates issue %v", i.URL)

	return i.HTMLURL, nil
}
