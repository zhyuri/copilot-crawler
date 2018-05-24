package lib

import (
	"testing"
)

func TestNewGithubUser(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    GithubUser
		wantErr bool
	}{
		{
			name:    "query Github User",
			args:    args{name: "zhyuri"},
			want:    GithubUser{Login: "zhyuri"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGithubUser(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGithubUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Login != tt.want.Login {
				t.Errorf("NewGithubUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
