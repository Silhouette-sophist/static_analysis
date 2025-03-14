package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"io"
	"net/http"
	"time"
)

const (
	GitHubApiBaseUrl     = "https://api.github.com"
	GitHubOwner          = "Silhouette-sophist"
	Accept               = "Accept"
	AcceptType           = "application/vnd.github+json"
	Authorization        = "Authorization"
	AuthorizationToken   = "Bearer %s"
	GitHubApiVersion     = "X-GitHub-Api-Version"
	GitHubApiVersionInfo = "2022-11-28"
	GitHubToken          = "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"
)

var gitHubClient = http.Client{
	Timeout: 30 * time.Second,
}

type GithubRepoParam struct {
	Owner string
	Repo  string
}

type GithubPullRequestParam struct {
	GithubRepoParam
	PullRequestId int32
}

func QueryUserRepos(ctx context.Context, param GithubRepoParam) (any, error) {
	url := fmt.Sprintf("%s/users/%s/repos", GitHubApiBaseUrl, param.Owner)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		hlog.CtxWarnf(ctx, "QueryUserRepos request err %v", err)
		return nil, err
	}
	result, err := sendRequestInternal[any](ctx, request)
	if err != nil {
		hlog.CtxWarnf(ctx, "QueryUserRepos sendRequestInternal err %v", err)
		return nil, err
	}
	return result, nil
}

func QueryPullRequestFiles(ctx context.Context, param GithubPullRequestParam) (any, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/pulls/%d/files", GitHubApiBaseUrl, param.Owner, param.Repo, param.PullRequestId)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		hlog.CtxWarnf(ctx, "QueryPullRequestFiles request err %v", err)
		return nil, err
	}
	result, err := sendRequestInternal[any](ctx, request)
	if err != nil {
		hlog.CtxWarnf(ctx, "QueryPullRequestFiles sendRequestInternal err %v", err)
		return nil, err
	}
	return result, nil
}

func sendRequestInternal[T any](ctx context.Context, req *http.Request) (*T, error) {
	req.Header.Set(Accept, AcceptType)
	req.Header.Set(Authorization, fmt.Sprintf(AuthorizationToken, GitHubToken))
	req.Header.Set(GitHubApiVersion, GitHubApiVersionInfo)
	response, err := gitHubClient.Do(req)
	if err != nil {
		hlog.CtxWarnf(ctx, "sendRequestInternal do err %v", err)
		return nil, err
	}
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		hlog.CtxWarnf(ctx, "sendRequestInternal ReadAll err %v", err)
		return nil, err
	}
	var result T
	if err = json.Unmarshal(respBody, &result); err != nil {
		hlog.CtxWarnf(ctx, "sendRequestInternal Unmarshal err %v", err)
		return nil, err
	}
	return &result, nil
}
