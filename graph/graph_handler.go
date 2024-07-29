package graph

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/asinha24/graph-rest-api/graph/model"
)

type Graph interface {
	CreateGraph(ctx context.Context, graph *model.Graph) (string, error)
	GetShortestPath(ctx context.Context, id, start, end string) ([]string, error)
	DeleteGraph(ctx context.Context, id string) error
}

type graphInMem struct {
	graphs  map[string]model.Graph
	mu      sync.Mutex
	graphID int
}

func NewgraphInMem() *graphInMem {
	return &graphInMem{
		graphs:  make(map[string]model.Graph),
		graphID: 1,
	}
}

func (n *graphInMem) CreateGraph(ctx context.Context, graph *model.Graph) (string, error) {
	n.mu.Lock()
	id := fmt.Sprintf("%d", n.graphID)
	n.graphID++
	n.graphs[id] = model.Graph{Nodes: graph.Nodes}
	n.mu.Unlock()
	return id, nil
}

func (n *graphInMem) GetShortestPath(ctx context.Context, id, start, end string) ([]string, error) {
	n.mu.Lock()
	graph, exists := n.graphs[id]
	n.mu.Unlock()
	if !exists {
		return nil, errors.New("graph not found")
	}

	path := n.findShortestPath(graph, start, end)
	if path == nil {
		return nil, errors.New("path not found")
	}

	return path, nil
}

func (n *graphInMem) DeleteGraph(ctx context.Context, id string) error {
	n.mu.Lock()
	_, exists := n.graphs[id]
	if exists {
		delete(n.graphs, id)
	}
	n.mu.Unlock()

	if !exists {
		return errors.New("graph not found")
	}

	return nil
}

func (s *graphInMem) findShortestPath(graph model.Graph, start, end string) []string {
	if start == end {
		return []string{start}
	}

	queue := [][]string{{start}}
	visited := map[string]bool{start: true}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		node := path[len(path)-1]

		for _, neighbor := range graph.Nodes[node] {
			if visited[neighbor] {
				continue
			}

			newPath := append([]string{}, path...)
			newPath = append(newPath, neighbor)
			if neighbor == end {
				return newPath
			}

			visited[neighbor] = true
			queue = append(queue, newPath)
		}
	}

	return nil
}
