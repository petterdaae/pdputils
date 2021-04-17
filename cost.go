package pdputils

func Cost(instance *Instance, solution *Solution) int {
	cost := 0

	currentCalls := make([]bool, instance.NumberOfCalls)

	// Node and travel costs
	for v := 0; v < instance.NumberOfVehicles; v++ {
		previousNode := instance.Vehicles[v].HomeNode
		for _, c := range solution.Routes[v] {
			call := instance.Calls[c]
			if currentCalls[c] {
				cost += instance.NodeTimesAndAndCosts[v][c].DestinationCost
				cost += instance.TravelTimesAndCosts[v][previousNode][call.DestinationNode].Cost
				previousNode = call.DestinationNode
			} else {
				cost += instance.NodeTimesAndAndCosts[v][c].OriginCost
				cost += instance.TravelTimesAndCosts[v][previousNode][call.OriginNode].Cost
				previousNode = call.OriginNode
				currentCalls[c] = true
			}
		}
	}

	// Cost of not transporting
	for _, node := range solution.Routes[instance.NumberOfVehicles] {
		if !currentCalls[node] {
			cost += instance.Calls[node].CostOfNotTransporting
			currentCalls[node] = true
		}
	}

	return cost
}
