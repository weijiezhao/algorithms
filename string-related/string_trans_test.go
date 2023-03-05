package stringrelated

import "testing"

func Test_trans(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Test case 1",
			args: args{s: "Hello World", n: 11},
			want: "wORLD hELLO",
		},
		{
			name: "Test case 2",
			args: args{s: "Abc a", n: 5},
			want: "A aBC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trans(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("trans() = %v, want %v", got, tt.want)
			}
		})
	}
}
