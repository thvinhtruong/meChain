package utils

import "bytes"

func IsEqual(a, b []byte) bool {
	return bytes.Equal(a, b)
}
