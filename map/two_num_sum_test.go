package mapp

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numers []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{numers: []int{3, 2, 4}, target: 6},
			want: []int{2, 3},
		},
		{
			name: "test case 2",
			args: args{numers: []int{8, 2, 4}, target: 10},
			want: []int{1, 2},
		},
		{
			name: "test case 3",
			args: args{numers: []int{8, 2, 4}, target: 9},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
