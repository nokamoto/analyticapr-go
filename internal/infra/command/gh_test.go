package command

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "github.com/nokamoto/analyticapr-go/pkg/api/v1"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/testing/protocmp"
)

func mustJSON(ps pulls) []byte {
	bs, _ := json.Marshal(ps)
	return bs
}

func TestGh_ListPulls(t *testing.T) {
	internalErr := errors.New("internal error")
	since := "2024-03-06"
	tests := []struct {
		name    string
		mock    func(*Mockrunner)
		repo    *v1.Repository
		want    []*v1.PullRequest
		wantErr error
	}{
		{
			name: "ok",
			mock: func(m *Mockrunner) {
				m.EXPECT().runO(
					"gh", "pr", "list",
					"--repo", "nokamoto/analyticapr-go",
					"--state", "merged",
					"--search", "created:>=2024-03-06",
					"--json", "number,title,author,comments,mergedAt,reviews",
				).Return(mustJSON(pulls{
					{
						Number: 1,
						Title:  "title1",
						Author: user{
							Login: "author1",
						},
						Comments: comments{
							{
								Author: user{
									Login: "comment1",
								},
							},
						},
						MergedAt: "2024-03-06T00:00:00Z",
						Reviews: reviews{
							{
								Author: user{
									Login: "review1",
								},
								State: "APPROVED",
							},
							{
								Author: user{
									Login: "review2",
								},
								State: "COMMENTED",
							},
						},
					},
				}), nil)
			},
			repo: &v1.Repository{
				Owner: "nokamoto",
				Repo:  "analyticapr-go",
			},
			want: []*v1.PullRequest{
				{
					Number: 1,
					Title:  "title1",
					Author: "author1",
					Comments: []*v1.Comment{
						{
							Author: "comment1",
						},
					},
					Reviews: []*v1.Review{
						{
							Author: "review1",
							State:  v1.Review_STATE_APPROVED,
						},
						{
							Author: "review2",
							State:  v1.Review_STATE_COMMENTED,
						},
					},
				},
			},
		},
		{
			name: "error if failed to run gh pr list",
			mock: func(m *Mockrunner) {
				m.EXPECT().runO(gomock.Any(), gomock.Any()).Return(nil, internalErr)
			},
			wantErr: internalErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			m := NewMockrunner(ctrl)
			if tt.mock != nil {
				tt.mock(m)
			}
			g := &Gh{runner: m}
			got, err := g.ListPulls(tt.repo, since)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Gh.ListPulls() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want, protocmp.Transform()); diff != "" {
				t.Errorf("Gh.ListPulls() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
