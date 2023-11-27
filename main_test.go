package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	asrt := assert.New(t)
	arr := []int{1, 2, 3, 33, 12, 14, 461, 11, 1, 144}
	res := linearSearch(arr, 461)
	asrt.Equal(res, 6, "Result should match idx of passed value.")
}
