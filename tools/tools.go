package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadInput(ex string) (string, error) {
	dir := os.Getenv("INPUTS")
	file := fmt.Sprintf("%s/%s.txt", dir, ex)
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(bytes), "\n"), nil
}
