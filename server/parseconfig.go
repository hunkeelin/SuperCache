package supercache

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strconv"
)

func (c *conn) parseconfig() (map[string]conf, error) {
	toreturn := make(map[string]conf)
	f, err := ioutil.ReadFile(c.confpath)
	if err != nil {
		return toreturn, err
	}
	blines := bytes.Split(f, []byte("\n"))
	var t conf
	var current string
	for _, line := range blines {
		tlines := bytes.Replace(line, []byte(" "), []byte(""), -1)
		switch {
		case bytes.HasPrefix(tlines, []byte("path")):
			f := bytes.Split(tlines, []byte(":"))
			if len(f) != 2 {
				return toreturn, errors.New("config error at: " + string(tlines))
			}
			current = string(f[1])
			t = toreturn[current]
			t.path = current
			toreturn[current] = t
		case bytes.HasPrefix(tlines, []byte("-url")):
			f := bytes.Split(tlines, []byte("-url:"))
			if len(f) != 2 {
				return toreturn, errors.New("config error at: " + string(tlines))
			}
			t = toreturn[current]
			t.url = string(f[1])
			toreturn[current] = t
		case bytes.HasPrefix(tlines, []byte("-time")):
			f := bytes.Split(tlines, []byte("-time:"))
			if len(f) != 2 {
				return toreturn, errors.New("config error at: " + string(tlines))
			}
			t = toreturn[current]
			time, err := strconv.Atoi(string(f[1]))
			if err != nil {
				return toreturn, errors.New("config error at: " + string(tlines))
			}
			t.time = time
			toreturn[current] = t
		default:
			continue
		}
	}
	return toreturn, nil
}
