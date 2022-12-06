package util

import "os"

func Load(filename string) string {
	// Load input
	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(input)
}
