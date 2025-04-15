package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	vertices map[int][]int
}

func (g *Graph) AddEdge(src, dest int) {
	g.vertices[src] = append(g.vertices[src], dest)
	g.vertices[dest] = append(g.vertices[dest], src)
}

// DFS - Depth First Search
func (g *Graph) DFS(start int, visited map[int]bool) {
	visited[start] = true
	fmt.Printf("%d ", start)

	for _, neighbor := range g.vertices[start] {
		if !visited[neighbor] {
			g.DFS(neighbor, visited)
		}
	}
}

// BFS - Breadth First Search
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := list.New()

	visited[start] = true
	queue.PushBack(start)

	for queue.Len() > 0 {
		element := queue.Front()
		node := element.Value.(int)
		queue.Remove(element)

		fmt.Printf("%d ", node)

		for _, neighbor := range g.vertices[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
			}
		}
	}
}

func main() {
	graph := Graph{vertices: make(map[int][]int)}

	// Build sample graph
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 5)
	graph.AddEdge(4, 6)
	graph.AddEdge(5, 6)

	fmt.Println("DFS traversal starting from node 1:")
	visitedDFS := make(map[int]bool)
	graph.DFS(1, visitedDFS)

	fmt.Println("\nBFS traversal starting from node 1:")
	graph.BFS(1)
}
