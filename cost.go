package pdputils

func Cost(instance *Instance, solution *Solution) int {
	cost := 0

	// Node and travel costs
	for v := 0; v < instance.NumberOfVehicles; v++ {
		startedCalls := make(map[int]bool)
		previousNode := instance.Vehicles[v].HomeNode
		for _, c := range solution.Routes[v] {
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
	}

	// Cost of not transporting
	currentCalls := make(map[int]bool)
	for _, node := range solution.Routes[instance.NumberOfVehicles] {
		if _, present := currentCalls[node]; !present {
			cost += instance.Calls[node].CostOfNotTransporting
			currentCalls[node] = true
		}
	}

	return cost
}
