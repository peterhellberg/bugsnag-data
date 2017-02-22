package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/peterhellberg/link"
)

func main() {
	var key string

	flag.StringVar(&key, "key", "", "Data access API key (required)")

	flag.Parse()

	path := flag.Arg(0)

	if key == "" || path == "" {
		flag.Usage()
		return
	}

	c := newClient(key)

	rawurl := "https://api.bugsnag.com" + path

	enc := json.NewEncoder(os.Stdout)

	for {
		res, err := c.Get(rawurl)
		if err != nil {
			return
		}

		if strings.HasPrefix(res.Header.Get("Content-Type"), "application/json") {
			var v interface{}

			json.NewDecoder(res.Body).Decode(&v)

			enc.Encode(v)
		}

		if res.StatusCode != http.StatusOK {
			break
		}

		io.CopyN(ioutil.Discard, res.Body, 64)
		res.Body.Close()

		group := link.ParseResponse(res)

		if next, ok := group["next"]; ok && next.URI != rawurl {
			rawurl = next.URI
		} else {
			break
		}

		time.Sleep(50 * time.Millisecond)
	}
}

type client struct {
	*http.Client
	key string
}

func newClient(key string) *client {
	return &client{&http.Client{Timeout: 10 * time.Second}, key}
}

func (c *client) Get(rawurl string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, rawurl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+c.key)

	fmt.Fprintln(os.Stderr, "\nGET", req.URL)

	return c.Do(req)
}
