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

	"github.com/markllama/hexgame-server/hexgame"
)

func TestGetGameList(t *testing.T) {
	var (
		request *http.Request
		recorder *httptest.ResponseRecorder
	)

	server := NewServer()

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/maps/", nil)
	server.ServeHTTP(recorder, request)

	allGames := []hexgame.Game {
		
	}

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected %v; received %v", http.StatusOK, recorder.Code)
	}
	payload, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}
	err = json.Unmarshal(payload, &allGames)
	if err != nil {
		t.Errorf("Error unmarshaling response to a list of maps: %v", err)
	}
}
	
func TestGetGame(t *testing.T) {
	var (
		request *http.Request
		recorder *httptest.ResponseRecorder
	)

	server := NewServer()

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/maps/12345", nil)
	server.ServeHTTP(recorder, request)

	aGame := hexgame.Game {
		
	}

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected %v; received %v", http.StatusOK, recorder.Code)
	}
	payload, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}
	err = json.Unmarshal(payload, &aGame)
	if err != nil {
		t.Errorf("Error unmarshaling response to a request for a map: %v", err)
	}

	if aGame.Name != "12345" {
		t.Errorf("Wrong ID: %s: %v", aGame.Name, err)		
	}
}


