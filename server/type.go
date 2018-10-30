package supercache

import (
	"sync"
)

type pcache struct {
	data  []byte
	err   error
	ready bool
}
type conf struct {
	path string
	url  string
	time int
}
type conn struct {
	kb, tb, cb []byte // key,cert,ca in bytes
	cacheMu    sync.Mutex
	cacherdy   map[string]pcache
	config     map[string]conf // path:url
	confpath   string          // path:url
	defpath    string
}
