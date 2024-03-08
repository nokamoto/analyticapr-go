package usecase

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "github.com/nokamoto/analyticapr-go/pkg/api/v1"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestAnalyticapr_GetAnalytica(t *testing.T) {
	internalErr := errors.New("internal")
	config := &v1.Config{
		Repositories: []*v1.Repository{
			{
				Owner: "test1",
			},
			{
				Owner: "test2",
			},
		},
		Since: "test",
	}
	tests := []struct {
		name    string
		mock    func(*Mockgh)
		want    *v1.Analyticapr
		wantErr error
	}{
		{
			name: "ok",
			mock: func(gh *Mockgh) {
				gh.EXPECT().ListPulls(config.GetRepositories()[0], config.GetSince()).Return([]*v1.PullRequest{
					{
						Number: 1,
					},
				}, nil)
				gh.EXPECT().ListPulls(config.GetRepositories()[1], config.GetSince()).Return([]*v1.PullRequest{
					{
						Number: 2,
					},
				}, nil)
			},
			want: &v1.Analyticapr{
				Repositories: []*v1.RepositoryAnalytica{
					{
						Repository: config.GetRepositories()[0],
						Pulls: []*v1.PullRequest{
							{
								Number: 1,
							},
						},
					},
					{
						Repository: config.GetRepositories()[1],
						Pulls: []*v1.PullRequest{
							{
								Number: 2,
							},
						},
					},
				},
			},
		},
		{
			name: "error if ListPulls fails",
			mock: func(gh *Mockgh) {
				gh.EXPECT().ListPulls(gomock.Any(), gomock.Any()).Return(nil, internalErr)
			},
			wantErr: internalErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			gh := NewMockgh(ctrl)
			if tt.mock != nil {
				tt.mock(gh)
			}
			a := &Analyticapr{
				gh:     gh,
				config: config,
			}
			got, err := a.GetAnalytica()
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetAnalytica() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want, protocmp.Transform()); diff != "" {
				t.Errorf("GetAnalytica() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
