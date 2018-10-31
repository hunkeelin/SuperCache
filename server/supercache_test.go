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
	err = c.cacheup("/noob/shit")
	if err != nil {
		panic(err)
	}
}
