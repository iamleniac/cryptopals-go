package challenge2

import (
	"encoding/hex"
	"errors"
)

func FixedXor(str1 string, str2 string) (string, error) {
	if len(str1) != len(str2) {
		return "", errors.New("inputs must have the same length")
	}

	input1 := []byte(str1)
	input2 := []byte(str2)

	raw1 := make([]byte, hex.DecodedLen(len(input1)))
	raw2 := make([]byte, hex.DecodedLen(len(input2)))

	result := make([]byte, hex.DecodedLen(len(input1)))

	_, err := hex.Decode(raw1, input1)
	if err != nil {
		return "", err
	}

	_, err = hex.Decode(raw2, input2)
	if err != nil {
		return "", err
	}

	for i := range raw1 {
		result[i] = raw1[i] ^ raw2[i]
	}

	return hex.EncodeToString(result), nil
}
