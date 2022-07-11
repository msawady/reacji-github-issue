package ext_service

import (
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

func NewGitHubService(sc config.SystemConfig) *GitHubService {

	tr := http.DefaultTransport

	pemBin, err := base64.StdEncoding.DecodeString(sc.GitHubPemBinary)
	if err != nil {
		log.Fatalf("Falied to decode GITHUB_PEM_BINARY %s", err)
	}

	itr, err := ghinstallation.New(tr, sc.GitHubAppID, sc.GitHubInstallationID, pemBin)
	return &GitHubService{client: github.NewClient(&http.Client{Transport: itr})}
}
