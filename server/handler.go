package supercache

import (
	"github.com/json-iterator/go"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (c *conn) MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if cache, ok := c.cacherdy[r.URL.Path]; ok {
			if cache.ready {
				w.WriteHeader(200)
				w.Write(cache.data)
				return
			}
		}
	}
	http.Redirect(w, r, c.defpath+r.URL.Path, 301)
}
