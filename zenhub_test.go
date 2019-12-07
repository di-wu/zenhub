package zenhub

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"os"
	"testing"
)

var (
	zenHubTestClient    *Client
	repoID, issueNumber int
	workspaceID         string
)

func TestMain(m *testing.M) {
	setupGitHub()
	setupZenHub()
	os.Exit(m.Run())
}

func setupGitHub() {
	// get github secret
	githubSecret, ok := os.LookupEnv("GITHUB_SECRET")
	if !ok {
		log.Panicln("could not get github secret")
	}

	// get owner
	owner, ok := os.LookupEnv("TEST_REPO_OWNER")
	if !ok {
		log.Panicln("could not get repo owner")
	}

	// get repo name
	name, ok := os.LookupEnv("TEST_REPO_NAME")
	if !ok {
		log.Println("could not get repo name")
		os.Exit(0)
	}

	// new github client
	gitHubTestClient := github.NewClient(
		oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: githubSecret},
		)),
	)

	// get repo
	githubTestRepo, _, err := gitHubTestClient.Repositories.Get(context.Background(), owner, name)
	if err != nil {
		log.Panicln(err)
	}
	repoID = int(*githubTestRepo.ID)

	// get issues with zenhub-test label
	issues, _, _ := gitHubTestClient.Issues.ListByRepo(context.Background(), owner, name, &github.IssueListByRepoOptions{
		Labels: []string{"zenhub-test"},
	})

	title := "zenhub test issue"
	body := "this issue is used for zenhub tests"

	// create issue if none found
	if len(issues) == 0 {
		issue, _, err := gitHubTestClient.Issues.Create(context.Background(), owner, name, &github.IssueRequest{
			Title:  &title,
			Body:   &body,
			Labels: &[]string{"test-issue"},
		})
		if err != nil {
			log.Panicln(err)
		}
		issueNumber = *issue.Number
	} else {
		// get first issue with correct label and name
		for _, issue := range issues {
			if *issue.Title == title {
				issueNumber = *issue.Number
				break
			}
		}
	}

	// failsafe
	if issueNumber == 0 {
		log.Panicln("no issue number found")
	}
}

func setupZenHub() {
	// get zenhub secret
	zenhubSecret, ok := os.LookupEnv("ZENHUB_SECRET")
	if !ok {
		log.Println("could not get zenhub secret")
		os.Exit(0)
	}

	// new zenhub client
	var err error
	zenHubTestClient, err = NewClient(Options.Secret(zenhubSecret))
	if err != nil {
		log.Panicln(err)
	}

	workspaces, _, err := zenHubTestClient.GetWorkspaces(repoID)
	if err != nil {
		log.Panicln(err)
	}

	if len(*workspaces) == 0 {
		log.Panicln("no workspaces found")
	}

	workspaceID = *(*workspaces)[0].ID
}
