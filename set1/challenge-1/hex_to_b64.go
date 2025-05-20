package challenge1

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToB64(input string) (string, error) {
	strToBytes := []byte(input)

	raw := make([]byte, hex.DecodedLen(len(strToBytes)))
	_, err := hex.Decode(raw, strToBytes)
	if err != nil {
		return "", err
	}

	b64 := base64.StdEncoding.EncodeToString(raw)
	return b64, nil
}
