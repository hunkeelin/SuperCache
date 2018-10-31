package main

import (
	"fmt"
	"sync"
	"time"
)

func (c *conn) dowork(m string, stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("done")
			c.wg.Done()
			return
		default:
			time.Sleep(300 * time.Millisecond)
			fmt.Println(m)
		}
	}
}

// c.stopmap = map[string]chan struct{}

func (c *conn) init() {
	c.wg.Add(len(c.foomap))
	for f, m := range c.foomap {
		stop := make(chan struct{})
		c.stopmap[f] = stop
		go c.dowork(m, c.stopmap[f])
	}
}

func (c *conn) stopgoroutine(m string) {
	close(c.stopmap[m])
}

type conn struct {
	stopmap map[string]chan struct{}
	foomap  map[string]string
	wg      sync.WaitGroup
}

func main() {
	sm := make(map[string]chan struct{})
	ff := make(map[string]string)
	ff["noob"] = "hi"
	ff["haha"] = "byebye"
	c := conn{
		stopmap: sm,
		foomap:  ff,
	}
	go c.init()
	time.Sleep(2 * time.Second)
	c.stopgoroutine("noob")
	time.Sleep(2 * time.Second)
	c.stopgoroutine("haha")
	c.wg.Wait()
}
