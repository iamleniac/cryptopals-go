package challenge1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexToB64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	result, err := HexToB64(input)

	assert.NoError(t, err)

	assert.Equal(t, expected, result, "result should be as expected")
}
