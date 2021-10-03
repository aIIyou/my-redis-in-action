package chapter2

import (
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestCheckToken(t *testing.T) {
	type args struct {
		rdb   *redis.Client
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "zhangsan case",
			args: args{
				rdb: redis.NewClient(&redis.Options{
					Addr:     "10.160.69.177:6380",
					Password: "a2b1c4d3",
					DB:       1,
				}),
				token: "abcd1234",
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckToken(tt.args.rdb, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
