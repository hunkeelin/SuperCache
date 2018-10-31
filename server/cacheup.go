package supercache

import (
	"errors"
	"github.com/hunkeelin/mtls/klinreq"
	"strings"
)

type epayload struct {
	empty []byte `json:"empty"`
}

func (c *conn) cacheup(path string) {
	var cache pcache
	dest := strings.Split(c.config[path].url, "://")
	if len(dest) != 2 {
		cache.err = errors.New("something is wrong with " + c.config[path].url)
		cache.ready = false
		c.cacherdy[path] = cache
		return
	}
	p := epayload{}
	i := &klinreq.ReqInfo{
		Dest:    dest[1],
		Method:  "GET",
		Payload: p,
		Route:   path,
	}
	if !strings.HasPrefix(c.config[path].url, "https") {
		i.Http = true
	} else {
		i.TrustBytes = c.config[path].ca
		i.CertBytes = c.config[path].cert
		i.KeyBytes = c.config[path].key
	}
	resp, err := klinreq.SendPayload(i)
	if err != nil {
		cache.err = err
		cache.ready = false
		c.cacherdy[path] = cache
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cache.err = err
		cache.ready = false
		c.cacherdy[path] = cache
		return
	}
	c.cacheMu.Lock()
	cache.ready = true
	cache.data = body
	cache.err = nil
	c.cacheready[path] = cache
	c.cacheMu.Unlock()
}
