package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/itchio/go-brotli/dec"
)

var channelID string

func init() {
	channelID = os.Getenv("CHANNEL_ID")
	if channelID == "" {
		log.Fatal("Channel ID is empty.")
	}
}

var query = map[string]string{
	"view":      "2",
	"flow":      "list",
	"live_view": "501",
	"pbj":       "1",
}

var header = map[string]string{
	"content-type":             "application/json; charset=UTF-8",
	"accept-encoding":          "br",
	"x-youtube-client-name":    "2",
	"x-youtube-client-version": "2.20200513.00.00",
	"user-agent":               "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1",
}

func main() {
	u, err := url.Parse("https://m.youtube.com/channel/" + channelID + "/videos")
	if err != nil {
		panic(err)
	}
	q := u.Query()

	for k, v := range query {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		panic(err)
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	c := new(http.Client)
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.ContentLength == -1 {
		panic("content-length is unknown")
	}

	if e := resp.Header.Get("content-encoding"); e != "br" {
		panic(e + " is not supported encoding")
	}

	r := dec.NewBrotliReader(resp.Body)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
