package pdputils

func Cost(instance *Instance, solution []int) int {
	// Strip outsourced calls
	zeroIndices := getZeroIndices(solution)
	outSourced := solution[zeroIndices[len(zeroIndices)-1]+1:]
	solution = solution[:zeroIndices[len(zeroIndices)-1]]

	cost := 0
	startIndex := 0

	// Node and travel costs
	for v := 0; v < instance.NumberOfVehicles; v++ {
		route := solution[startIndex:zeroIndices[v]]
		startedCalls := make(map[int]bool)
		previousNode := instance.Vehicles[v].HomeNode
		for _, c := range route {
			call := instance.Calls[c]
			if _, present := startedCalls[c]; present {
				cost += instance.NodeTimesAndAndCosts[v][c].DestinationCost
				cost += instance.TravelTimesAndCosts[v][previousNode][call.DestinationNode].Cost
				previousNode = call.DestinationNode
			} else {
				cost += instance.NodeTimesAndAndCosts[v][c].OriginCost
				cost += instance.TravelTimesAndCosts[v][previousNode][call.OriginNode].Cost
				previousNode = call.OriginNode
				startedCalls[c] = true
			}
		}
		startIndex = zeroIndices[v] + 1
	}

	// Cost of not transporting
	currentCalls := make(map[int]bool)
	for _, node := range outSourced {
		if _, present := currentCalls[node]; !present {
			cost += instance.Calls[node].CostOfNotTransporting
			currentCalls[node] = true
		}
	}

	return cost
}
