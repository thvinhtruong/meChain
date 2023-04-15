package utils

import "fmt"

func Logger(data string, args ...interface{}) {
	fmt.Printf(data, args...)
}
