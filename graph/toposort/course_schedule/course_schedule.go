/*
Problem Url: https://leetcode.com/problems/course-schedule/

Constraints:

The number of nodes in the graph is in the range [1, 2000].
0 <= Node.val <= 1000
It is guaranteed that the graph represents a directed acyclic graph (DAG).

*/

package main

import (
	"maps"
	"fmt"
)

type Set map[int]struct{}

type Graph struct {
	adjacencyList map[int]Set
	indegrees     map[int]int
}

func NewGraph() Graph {
	return Graph{
		adjacencyList: make(map[int]Set),
		indegrees: make(map[int]int),
	}
}

func (graph *Graph) AddNode(node int) {
	if _, nodeInGraph := graph.adjacencyList[node]; !nodeInGraph {
		graph.adjacencyList[node] = Set{}
		graph.indegrees[node] = 0
	}
}

func (graph *Graph) AddEdge(sourceNode, destinationNode int) {
	graph.AddNode(sourceNode)
	graph.AddNode(destinationNode)
	if _, edgeExists := graph.adjacencyList[sourceNode][destinationNode]; !edgeExists {
		graph.adjacencyList[sourceNode][destinationNode] = struct{}{}
		graph.indegrees[destinationNode] += 1
	}
}

func (graph *Graph) ToposortUsingKahnsAlgorithm() (isDirectedAcyclicGraph bool) {
	nodesWithZeroIndegree := make([]int, 0, len(graph.adjacencyList))
	for node, indegree := range(graph.indegrees) {
		if indegree == 0 {
			nodesWithZeroIndegree = append(nodesWithZeroIndegree, node)
		}
	}

	indegrees := maps.Clone(graph.indegrees)
	toposort := make([]int, 0, len(graph.adjacencyList))

	for len(nodesWithZeroIndegree) > 0 {

		node := nodesWithZeroIndegree[0]
		// nodesWithZeroIndegree = slices.Delete(nodesWithZeroIndegree, 0, 1) -> O(N) operation
		nodesWithZeroIndegree = nodesWithZeroIndegree[1:] // O(1) operation
		toposort = append(toposort, node)

		for destinationNode, _ := range(graph.adjacencyList[node]) {
			indegrees[destinationNode] -= 1
			if indegrees[destinationNode] == 0 {
				nodesWithZeroIndegree = append(nodesWithZeroIndegree, destinationNode)
			}
		}
	}
	isDirectedAcyclicGraph = len(toposort) == len(graph.adjacencyList)
	return isDirectedAcyclicGraph
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 || len (prerequisites) == 0 {
		return true
	}
	graph := generateGraph(numCourses, prerequisites)
	return graph.ToposortUsingKahnsAlgorithm()
}

func generateGraph(numCourses int, prerequisites [][]int) (graph Graph) {
	graph = NewGraph()
	for node := range(numCourses) {
		graph.AddNode(node)
	}
	for _, dependency := range(prerequisites) {
		sourceNode, destinationNode := dependency[0], dependency[1]
		graph.AddEdge(sourceNode, destinationNode)
	}
	return graph
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}
