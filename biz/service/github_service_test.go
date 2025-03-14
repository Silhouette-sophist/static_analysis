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
