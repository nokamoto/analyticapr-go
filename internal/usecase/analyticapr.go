//go:generate go run go.uber.org/mock/mockgen -source=$GOFILE -destination=$GOFILE.mock_test.go -package=$GOPACKAGE
package usecase

import (
	"fmt"

	v1 "github.com/nokamoto/analyticapr-go/pkg/api/v1"
)

type gh interface {
	ListPulls(*v1.Repository, string) ([]*v1.PullRequest, error)
}

type Analyticapr struct {
	gh     gh
	config *v1.Config
}

func NewAnalyticapr(gh gh, config *v1.Config) *Analyticapr {
	return &Analyticapr{gh: gh, config: config}
}

func (a *Analyticapr) GetAnalytica() (*v1.Analyticapr, error) {
	var res v1.Analyticapr
	for _, repo := range a.config.GetRepositories() {
		pulls, err := a.gh.ListPulls(repo, a.config.GetSince())
		if err != nil {
			return nil, fmt.Errorf("failed to list pulls: %w", err)
		}
		res.Repositories = append(res.Repositories, &v1.RepositoryAnalytica{
			Repository: repo,
			Pulls:      pulls,
		})
	}
	return &res, nil
}
