package graph

import "github.com/neil-vass/advent-of-code-2015-go/shared/priorityqueue"

type NodeCost[TNode any] struct {
	Node TNode
	Cost int
}

// Graph objects passed to A* search need these methods.
type GraphWithCosts[TNode comparable] interface {
	// Nodes you can get to in one step, along with cost to move there.
	Neighbours(node TNode) []NodeCost[TNode]

	// A* needs a heuristic: what is the estimated cost to get from this node
	// to the goal? If your heuristic:
	// - Underestimates costs, then search will find the correct answer
	//     (the more you underestimate, the less efficient search will be).
	// - Estimates perfectly, search will find the correct answer without needing to
	//     explore anything except the best path.
	// - Overestimates costs, then search might not find the correct answer
	//     (can find one path, and ignore others as they look more expensive).
	//
	// You can use A* search with no heuristic (just return 0) to do Dijkstra's algorithm.
	Heuristic(from TNode) float64

	// Confirm whether the given node meets the goal condition.
	GoalReached(candidate TNode) bool
}

func A_StarSearch[TNode comparable](g GraphWithCosts[TNode], start TNode) (goalFound bool, cost int) {
	frontier := priorityqueue.New[TNode]()
	visited := map[TNode]struct {
		costSoFar int
		cameFrom  TNode
	}{start: {costSoFar: 0}}

	frontier.Push(start, 0)

	for !frontier.IsEmpty() {
		current := frontier.Pull()
		if g.GoalReached(current) {
			goalFound = true
			cost = visited[current].costSoFar
			return
		}

		for _, n := range g.Neighbours(current) {
			newCost := visited[current].costSoFar + n.Cost
			old, beenHereBefore := visited[n.Node]

			// If we haven't been here before, _or_ if we've found a cheaper way to get here
			if !beenHereBefore || newCost < old.costSoFar {
				priority := float64(newCost) + g.Heuristic(n.Node)
				frontier.Push(n.Node, priority)
				visited[n.Node] = struct {
					costSoFar int
					cameFrom  TNode
				}{newCost, current}
			}
		}
	}

	// The end of all our exploring.
	goalFound = false
	return
}
