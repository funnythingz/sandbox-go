package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zenazn/goji/web"
)

func ParseResponse(res *http.Response) (string, int) {
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(contents), res.StatusCode
}

func Test_Hello(t *testing.T) {
	m := web.New()
	Route(m)
	ts := httptest.NewServer(m)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/hello/hoge")
	if err != nil {
		t.Error("unexpected")
	}
	c, s := ParseResponse(res)
	if s != http.StatusOK {
		t.Error("invalid status code")
	}
	if c != `{"Id":1,"Name":"hoge"}` {
		t.Error("invalid response")
	}
}
