package service

import (
	"context"
	"testing"
)

func TestQueryUserRepos(t *testing.T) {
	repos, err := QueryUserRepos(context.Background(), GithubRepoParam{
		Owner: GitHubOwner,
		Repo:  "static_analysis",
	})
	if err != nil {
		t.Errorf("QueryUserRepos err %v", err)
		return
	}
	t.Logf("QueryUserRepos success %v", repos)
}

func TestQueryPullRequestFiles(t *testing.T) {
	files, err := QueryPullRequestFiles(context.Background(), GithubPullRequestParam{
		GithubRepoParam{
			Owner: GitHubOwner,
			Repo:  "static_analysis",
		},
		3,
	})
	if err != nil {
		t.Errorf("QueryPullRequestFiles err %v", err)
		return
	}
	t.Logf("QueryPullRequestFiles success %v", files)
}
