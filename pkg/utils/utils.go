package utils

import (
	"github.com/slayv1/crud/pkg/types"
	"crypto/rand"
	"encoding/hex"
	
)


//GenerateTokenStr ...
func GenerateTokenStr() (string, error) {

	buffer := make([]byte, 256)
	n, err := rand.Read(buffer)
	if n != len(buffer) || err != nil {
		return "", types.ErrInternal
	}

	return hex.EncodeToString(buffer), nil
}
