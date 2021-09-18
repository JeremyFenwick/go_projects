package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func main() {
	// Gather all the required inputs and feed them into the relevant functions
	repoNameFlag := flag.String("repo", "N/A", "Input the name of the repository for deletion")
	ownerNameFlag := flag.String("owner", "N/A", "Input the name of the owner of the repository for deletion")
	accessToken := flag.String("token", "N/A", "Enter the access token in a text file here. (Ex token.txt)")
	flag.Parse()
	token := readAccessToken(accessToken)
	ctx, client := setupClient(token)
	repoOwner, repoName := getRepoForDeletion(repoNameFlag, ownerNameFlag)

	// Gather all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		println("Could not retrieve repositories...")
		os.Exit(1)
	}

	// Loop over the repository. If the repo name and owner name provided match the details of the repository,
	// delete the repository and print the response to terminal
	for _, repo := range repos {
		if repo.GetName() == repoName && repo.Owner.GetLogin() == repoOwner {
			response, err := client.Repositories.Delete(ctx, repo.Owner.GetLogin(), repoName)
			if err != nil {
				fmt.Println("Could not delete repository. Exiting...")
				os.Exit(1)
			}
			fmt.Println(response)
		}

	}
}

// Process the token for connecting to github
func readAccessToken(accessToken *string) string {
	tokenFile := *accessToken
	content, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		fmt.Println("Could not retrieve access token")
		os.Exit(1)
	}
	token := string(content)
	return token
}

// Setup the connection to github
func setupClient(tkn string) (context.Context, *github.Client) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tkn},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return ctx, client
}

// From the command line, get the owner and name of the repository for deletion
func getRepoForDeletion(repo *string, owner *string) (string, string) {
	repoName := *repo
	repoOwner := *owner
	// If either value is N/A, exit the program
	if repoName == "N/A" || repoOwner == "N/A" {
		fmt.Println("Please enter both the repo name and the owner name. Exiting...")
		os.Exit(1)
	}

	return repoOwner, repoName
}
