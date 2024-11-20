package common

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceRmDup(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "remove duplicate elements from slice",
			args: args{slice: []int{1, 2, 2, 3, 4, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "remove duplicate elements from slice, unordered",
			args: args{slice: []int{5, 3, 4, 4, 2, 1, 2, 3}},
			want: []int{5, 3, 4, 2, 1},
		},
		{
			name: "slice with negative numbers",
			args: args{slice: []int{-1, -2, -3, -3, -4, -5, -5}},
			want: []int{-1, -2, -3, -4, -5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SliceRmDup(tt.args.slice)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal([]int{3, 2, 1}, Reverse([]int{1, 2, 3}))
	assert.Equal([]string{"a", "b", "c", "d"}, Reverse([]string{"d", "c", "b", "a"}))
	assert.Equal([]float64{}, Reverse([]float64{}))
}

func TestUniq(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	result1 := Uniq([]int{1, 2, 2, 1})
	assert.Equal(len(result1), 2)
	assert.Equal(result1, []int{1, 2})
}

func TestDifferenceSet(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				b: []int{2, 4},
			},
			want: []int{1, 3, 5},
		},
		{
			name: "Test 2",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				b: []int{},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test 3",
			args: args{
				a: []int{},
				b: []int{1, 2, 3, 4, 5},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DifferenceSet(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DifferenceSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
