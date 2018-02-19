package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	tests := []int{3, 2, 4}
	target := 6

	ret := twoSum1(tests, target)
	assert.Len(t, ret, 2, "return 2")
	assert.Contains(t, ret, 1, "")
	assert.Contains(t, ret, 2, "")

	ret = twoSum2(tests, target)
	assert.Len(t, ret, 2, "return 2")
	assert.Contains(t, ret, 1, "")
	assert.Contains(t, ret, 2, "")
}
