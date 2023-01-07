package utils

import "testing"

func TestJoinIntSlice(t *testing.T) {
	type args struct {
		sep string
		s   []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				sep: ",",
				s:   []any{1, 2, 3, "21321", 21321},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(JoinIntSlice(tt.args.sep, tt.args.s...))
		})
	}
}
