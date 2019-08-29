package services

import (
	"fmt"
	helper "github.com/nikola43/go_github/helpers"
)

// :create (POST) functions
func CreateGithubRepository() {
	fmt.Println("CreateGithubRepository")
	helper.GithubHostName()
}

func CreateGithubRepositoryPayload() {
	fmt.Println("CreateGithubRepositoryPayload")
	helper.GithubHostName()
}

func CreateGithubRepositoryRequestBuilder() {
	fmt.Println("CreateGithubRepositoryRequestBuilder")
	helper.GithubHostName()
}

func CreateGithubRepositoryUrl() {
	fmt.Println("CreateGithubRepositoryUrl")
	helper.GithubHostName()
}
