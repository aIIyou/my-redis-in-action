package chapter1

import "testing"

func TestRClient_ArticalVote(t *testing.T) {
	type args struct {
		articalID string
		userID    string
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
				articalID: "article:00000001",
				userID:    "user00000001",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rds.ArticalVote(tt.args.articalID, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("RClient.ArticalVote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
