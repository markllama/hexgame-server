package service
//
//
//
import (
	//	"github.com/unrolled/render"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/markllama/hexgame-server/hexmap"
)

// var (
// 	formatter = render.New(render.Options{ IndentJSON: true })
// )

func TestGetMapList(t *testing.T) {
	var (
		request *http.Request
		recorder *httptest.ResponseRecorder
	)

	server := NewServer()

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/maps/", nil)
	server.ServeHTTP(recorder, request)

	allMaps := []hexmap.Map {
		
	}

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected %v; received %v", http.StatusOK, recorder.Code)
	}
	payload, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}
	err = json.Unmarshal(payload, &allMaps)
	if err != nil {
		t.Errorf("Error unmarshaling response to a list of maps: %v", err)
	}
}
	


