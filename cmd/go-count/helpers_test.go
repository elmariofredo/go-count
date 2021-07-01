package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// func (app *application) TestPing(t *testing.T) {
func TestPing(t *testing.T) {

	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	app := &application{
		config: defaultConfig,
	}

	app.ping(rr, r)

	rs := rr.Result()

	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	defer rs.Body.Close()

	body, err := ioutil.ReadAll(rs.Body)

	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want body to queal %q", "OK")
	}
}
