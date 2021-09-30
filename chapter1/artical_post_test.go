package chapter1

import "testing"

func TestRClient_ArticlePost(t *testing.T) {
	type args struct {
		articel Article
		userID  string
	}
	tests := []struct {
		name    string
		rds     *RClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "article one",
			rds:  NewClient("10.160.69.177:6380", "a2b1c4d3", 1),
			args: args{
				articel: Article{
					Title:  "Three Guns",
					Poster: "Annoymous",
					Link:   "www.sogou.com",
					Votes:  0,
				},
				userID: "user00000001",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rds.ArticlePost(tt.args.articel, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("RClient.ArticlePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
