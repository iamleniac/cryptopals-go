package challenge2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappyPath(t *testing.T) {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"

	expected := "746865206b696420646f6e277420706c6179"

	result, err := FixedXor(input1, input2)

	assert.NoError(t, err)

	assert.Equal(t, expected, result, "result should be as expected")
}

func TestDifferentInputSizes(t *testing.T) {
	input1 := "e0d524349fea2a02b2ab200a0d51d549"
	input2 := "45be8545834f3584dee63c4c0a27bedc8af"

	_, err := FixedXor(input1, input2)

	assert.Error(t, err)
}
