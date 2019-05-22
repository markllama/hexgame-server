package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/unrolled/render"
)


var (
	formatter = render.New(render.Options{ IndentJSON: true, })
)

func TestCreateGame(t *testing.T) {
	client := &http.Client{}

	server := httptest.NewServer(http.HandlerFunc(createGameHandler(formatter)))
	defer server.Close()

	body := []byte("{\n \"gridsize\": 19,\n \"players\": [\n {\n \"color\": \"white\",\n \"name\": \"bob\"\n },\n {\n \"color\": \"black\",\n \"name\": \"alfred\"\n }\n ]\n}")

	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error in creating POST request for createGameHandler: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in POST to createGameHandler: %v", err)
	}
	defer res.Body.Close()

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: 201, received %s", res.Status)
	}
	fmt.Printf("Payload: %s", string(payload))
}
