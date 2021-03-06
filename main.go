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
	var (
		key   string
		max   int
		count int
		delay time.Duration
	)

	flag.StringVar(&key, "key", "", "Data access API key (required)")
	flag.IntVar(&max, "max", 0, "Max number of requests to make (0 means unlimited)")
	flag.DurationVar(&delay, "delay", 1*time.Second, "The delay between each request")

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
		if max > 0 && count >= max {
			break
		}

		res, err := c.Get(rawurl)
		if err != nil {
			fmt.Fprintln(os.Stderr, "\nERROR", err.Error())
			return
		}

		if res.Header.Get("X-RateLimit-Remaining") == "1" {
			fmt.Fprintln(os.Stderr, "\nRATE Sleeping for 60s before being rate limited")
			time.Sleep(60 * time.Second)
		}

		if strings.HasPrefix(res.Header.Get("Content-Type"), "application/json") {
			var v interface{}

			json.NewDecoder(res.Body).Decode(&v)

			enc.Encode(v)
		}

		io.CopyN(ioutil.Discard, res.Body, 64)
		res.Body.Close()

		if res.StatusCode != http.StatusOK {
			fmt.Fprintln(os.Stderr, "\nERROR Unexpected status code", res.Status)
			break
		}

		group := link.ParseResponse(res)

		if next, ok := group["next"]; ok && next.URI != rawurl {
			rawurl = next.URI
		} else {
			break
		}

		count++

		time.Sleep(delay)
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
	req, err := c.request(rawurl)
	if err != nil {
		return nil, err
	}

	fmt.Fprintln(os.Stderr, "\nGET", req.URL)

	return c.Do(req)
}

func (c *client) request(rawurl string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, rawurl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+c.key)

	return req, nil
}
