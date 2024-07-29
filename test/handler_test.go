package test

import (
	"bytes"
	gohttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/asinha24/graph-rest-api/graph"
	graphHttp "github.com/asinha24/graph-rest-api/http"
)

func TestCreateNewGraph(t *testing.T) {
	graph := graph.NewgraphInMem()
	graphServer := graphHttp.NewGraphHandler(graph)

	var jsonStr = []byte(`'{"n1":["n2","n3"]}'`)

	req, err := gohttp.NewRequest("POST", "/graph", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := gohttp.HandlerFunc(graphServer.CreateNewGraph)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != gohttp.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, gohttp.StatusCreated)
	}

	// Check the response body is what we expect.
	//It will fail  just because every time the UUID is creating unique id
	expected := `[{"id": "1"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
