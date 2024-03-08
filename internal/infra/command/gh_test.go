package command

import (
	"errors"
	"testing"

	v1 "github.com/nokamoto/analyticapr-go/pkg/api/v1"
	gomock "go.uber.org/mock/gomock"
)

func TestGh_ListPulls(t *testing.T) {
	internalErr := errors.New("internal error")
	since := "2024-03-06"
	tests := []struct {
		name    string
		mock    func(*Mockrunner)
		repo    *v1.Repository
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
					"--json", "author,mergedAt",
				).Return([]byte(`[]`), nil)
			},
			repo: &v1.Repository{
				Owner: "nokamoto",
				Repo:  "analyticapr-go",
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
			_, err := g.ListPulls(tt.repo, since)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Gh.ListPulls() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
