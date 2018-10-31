package supercache

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	fmt.Println("testing parse")
	c := conn{
		confpath: "config",
	}
	m, err := c.parseconfig()
	c.config = m
	if err != nil {
		panic(err)
	}
	for _, f := range m {
		fmt.Println(f)
	}
}
