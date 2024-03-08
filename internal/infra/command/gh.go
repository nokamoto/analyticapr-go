package command

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	v1 "github.com/nokamoto/analyticapr-go/pkg/api/v1"
)

type Gh struct {
	runner runner
}

func NewGh() *Gh {
	return &Gh{runner: &runnerImpl{}}
}

func (g *Gh) ListPulls(repo *v1.Repository, since string) ([]*v1.PullRequest, error) {
	var host string
	if repo.GetGh() != "" {
		host = fmt.Sprintf("%s/", repo.GetGh())
	}
	r := fmt.Sprintf("%s%s/%s", host, repo.GetOwner(), repo.GetRepo())
	fields := []string{
		"number",
		"title",
		"author",
		"comments",
		"mergedAt",
		"reviews",
	}
	// https://docs.github.com/ja/search-github/searching-on-github/searching-issues-and-pull-requests
	q := fmt.Sprintf("created:>=%s", since)
	bs, err := g.runner.runO("gh", "pr", "list", "--repo", r, "--state", "merged", "--search", q, "--json", strings.Join(fields, ","))
	if err != nil {
		return nil, fmt.Errorf("failed to run gh pr list: %w", err)
	}
	var res pulls
	if err := json.Unmarshal(bs, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}
	slog.Debug("gh pr list", "output", res)
	return res.toPullRequests(), nil
}
