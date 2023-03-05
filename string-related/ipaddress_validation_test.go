package stringrelated

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{ip: "172.16.254.1"},
			want: "IPv4",
		},
		{
			name: "test case 2",
			args: args{ip: "2001:0db8:85a3:0:0:8A2E:0370:7334"},
			want: "IPv6",
		},
		{
			name: "test case 3",
			args: args{ip: "256.256.256.256"},
			want: "Neither",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.ip); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
