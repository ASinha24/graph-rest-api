package main

import (
	"flag"
	"log"
	gohttp "net/http"

	"github.com/asinha24/graph-rest-api/graph"
	graphHttp "github.com/asinha24/graph-rest-api/http"
	"github.com/gorilla/mux"
)

var port = flag.String("port", "8080", "port to listen")

func main() {
	flag.Parse()
	router := mux.NewRouter()

	graph := graph.NewgraphInMem()
	graphServer := graphHttp.NewGraphHandler(graph)
	graphServer.InstallRoutes(router)

	log.Println("starting http server, listening on port:", *port)
	if err := gohttp.ListenAndServe(":"+*port, router); err != nil {
		log.Fatalf("error in starting server: %v", err)
	}

}
