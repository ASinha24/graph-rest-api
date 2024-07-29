package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/asinha24/graph-rest-api/graph"
	"github.com/asinha24/graph-rest-api/graph/model"
	"github.com/asinha24/graph-rest-api/http/utils"
)

type GraphHandler struct {
	graphHandler graph.Graph
}

func (g *GraphHandler) InstallRoutes(mux *mux.Router) {
	// It will Create a undirected non weighted graph
	mux.Methods(http.MethodPost).Path("/graph").HandlerFunc(g.CreateNewGraph)
	// find the shorted path of between 2 graphs
	mux.Methods(http.MethodGet).Path("/graph/shortest_path/{graphID}/{start/{end}}").HandlerFunc(g.GetShortestPath)
	// delete graph
	mux.Methods(http.MethodPost).Path("/graph/delete_graph/{graphID}").HandlerFunc(g.DeleteGraph)
}

func NewGraphHandler(graph graph.Graph) *GraphHandler {
	return &GraphHandler{
		graphHandler: graph,
	}
}

func (g *GraphHandler) CreateNewGraph(w http.ResponseWriter, r *http.Request) {
	var nodes model.Graph
	if err := json.NewDecoder(r.Body).Decode(&nodes); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := g.graphHandler.CreateGraph(r.Context(), &nodes)
	if err != nil {
		utils.WriteErrorResponse(http.StatusInternalServerError, err, w)
	}

	utils.WriteResponse(http.StatusCreated, nodes, w)
	w.Write([]byte(id))
}

func (g *GraphHandler) GetShortestPath(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["graphID"]
	start := mux.Vars(r)["start"]
	end := mux.Vars(r)["end"]

	path, err := g.graphHandler.GetShortestPath(r.Context(), id, start, end)
	if path == nil {
		utils.WriteErrorResponse(http.StatusNotFound, err, w)
		return
	}

	response, err := json.Marshal(path)
	if err != nil {
		utils.WriteErrorResponse(http.StatusInternalServerError, err, w)
		return
	}

	utils.WriteResponse(http.StatusOK, response, w)
}

func (g *GraphHandler) DeleteGraph(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["graphID"]

	if err := g.graphHandler.DeleteGraph(r.Context(), id); err != nil {
		utils.WriteErrorResponse(http.StatusInternalServerError, err, w)
		return
	}

	utils.WriteResponse(http.StatusNoContent, nil, w)
}
