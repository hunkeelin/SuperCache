package supercache

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	m, err := parseconfig("config")
	if err != nil {
		panic(err)
	}
	for _, i := range m {
		fmt.Println(i)
	}
}
