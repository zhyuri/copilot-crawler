package lib

import (
	"testing"
)

func TestNewGithubRepo(t *testing.T) {
	type args struct {
		name  string
		owner string
	}
	tests := []struct {
		name    string
		args    args
		want    GithubRepo
		wantErr bool
	}{
		{
			name:    "query Github Repo",
			args:    args{name: "copilot-crawler", owner: "zhyuri"},
			want:    GithubRepo{ID: "MDEwOlJlcG9zaXRvcnkxMzQ2NDE2NTU="},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGithubRepo(tt.args.name, tt.args.owner)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGithubUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ID != tt.want.ID {
				t.Errorf("NewGithubUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
