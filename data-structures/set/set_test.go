package set_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/MueR/adventofcode.go/data-structures/set"
	"github.com/stretchr/testify/assert"
)

func TestIntSet(t *testing.T) {
	tests := []struct {
		valToAdd int
		wantHas  []int
	}{
		{5, []int{5}},
		{5, []int{5}},
		{10, []int{5, 10}},
		{20, []int{5, 10, 20}},
		{20, []int{5, 10, 20}},
		{2230, []int{5, 10, 20, 2230}},
		{123, []int{5, 10, 20, 123, 2230}},
	}

	intSet := set.NewIntSet(nil)
	for _, tt := range tests {
		intSet.Add(tt.valToAdd)
		for _, want := range tt.wantHas {
			assert.Truef(t, intSet.Has(want), "want IntSet.Has(%d) = true, got false", want)
		}
	}

	got := intSet.Keys()
	sort.Ints(got)
	want := []int{5, 10, 20, 123, 2230}
	assert.Truef(t, reflect.DeepEqual(got, want), "expected inset.Keys() to be %v, got %v", want, got)

	valsToRemove := []int{5, 10, 20, 123, 2230}
	for _, tt := range valsToRemove {
		intSet.Remove(tt)
		assert.Falsef(t, intSet.Has(tt), "want IntSet.Has(%d) = false, got true", tt)
	}

	got = intSet.Keys()
	assert.Truef(t, len(got) == 0, "want zero-length slice after removing all keys, got %v", got)
}
