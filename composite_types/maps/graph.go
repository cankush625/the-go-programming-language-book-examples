package main

import "fmt"

func addEdge(graph map[string]map[string]bool, from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(graph map[string]map[string]bool, from, to string) bool {
	return graph[from][to]
}

func main() {
	graph := make(map[string]map[string]bool)
	// Add edges
	addEdge(graph, "a", "b")
	addEdge(graph, "a", "c")
	addEdge(graph, "b", "c")
	addEdge(graph, "c", "a")
	// Check if edge is present
	fmt.Println(hasEdge(graph, "a", "c")) // "true"
	fmt.Println(hasEdge(graph, "b", "a")) // "false"
}
