package chapter2

import (
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestAddToCart(t *testing.T) {
	type args struct {
		rdb   *redis.Client
		user  string
		item  string
		count int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "item to cart",
			args: args{
				rdb: redis.NewClient(&redis.Options{
					Addr:     "10.160.69.177:6380",
					Password: "a2b1c4d3",
					DB:       1,
				}),
				user:  "abcd1234",
				item:  "00000001",
				count: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddToCart(tt.args.rdb, tt.args.user, tt.args.item, tt.args.count); (err != nil) != tt.wantErr {
				t.Errorf("AddToCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
