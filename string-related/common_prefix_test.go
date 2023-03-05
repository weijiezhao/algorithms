package stringrelated

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{ss: []string{"abc", "ab", "a"}},
			want: "a",
		},
		{
			name: "test case 2",
			args: args{ss: []string{"abc", "ab", "a", ""}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.ss); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
