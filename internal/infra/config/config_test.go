package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "github.com/nokamoto/analyticapr-go/pkg/api/v1"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name string
		file string
		want *v1.Config
	}{
		{
			name: "ok",
			file: "testdata/config.yaml",
			want: &v1.Config{
				Repositories: []*v1.Repository{
					{
						Owner: "foo",
						Repo:  "bar",
					},
					{
						Owner: "baz",
						Repo:  "qux",
						Gh:    "example.com",
					},
				},
				Since: "2024-03-01",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.file)
			if err != nil {
				t.Fatalf("NewConfig(%s) = _, %v, want _, nil", tt.file, err)
			}
			if diff := cmp.Diff(tt.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("NewConfig(%s) mismatch (-want +got):\n%s", tt.file, diff)
			}
		})
	}
}
