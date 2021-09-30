package chapter1

import (
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestRClient_AddAndRemoveGroups(t *testing.T) {
	type fields struct {
		Client *redis.Client
	}
	type args struct {
		articleID string
		toAdd     []string
		toRemove  []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "novel case",
			fields: fields{
				Client: redis.NewClient(&redis.Options{
					Addr:     "10.160.69.177:6380",
					Password: "a2b1c4d3",
					DB:       1,
				}),
			},
			args: args{
				articleID: "article:00000001",
				toAdd:     []string{"novel"},
				toRemove:  []string{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rdb := &RClient{
				Client: tt.fields.Client,
			}
			if err := rdb.AddAndRemoveGroups(tt.args.articleID, tt.args.toAdd, tt.args.toRemove); (err != nil) != tt.wantErr {
				t.Errorf("RClient.AddAndRemoveGroups() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
