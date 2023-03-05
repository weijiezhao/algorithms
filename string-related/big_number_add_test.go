package stringrelated

import "testing"

func Test_solveBigNumberAdd(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{s: "99", t: "1"},
			want: "100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveBigNumberAdd(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("solveBigNumberAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
