// @Author: Jiekun
// @Date: 2022/8/8

package sharding

import "testing"

func Test_sharding_Rewrite(t *testing.T) {
	tests := []struct {
		name string
		sql  string
		want string
	}{
		{
			"parse test",
			"select * from t where a = 1;",
			"select * from t where a = 1;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sharding{}
			if got := s.Rewrite(tt.sql); got != tt.want {
				t.Errorf("Rewrite() = %v, want %v", got, tt.want)
			}
		})
	}
}
