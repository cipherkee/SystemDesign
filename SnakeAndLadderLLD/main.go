package main

func main() {
	// Initialize orchestrator
	var (
		players = []string{"jack", "reacher"}
		size    = 10
		snakes  = [][]int{{5, 1}}
		ladders = [][]int{{3, 7}}
	)

	orch := InitializeOrchestrator(players, size, snakes, ladders)

	orch.RunGame()
}
