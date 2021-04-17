package pdputils

func Cost(instance *Instance, solution *Solution) int {
	cost := 0

	// Node and travel costs
	for v := 0; v < instance.NumberOfVehicles; v++ {
		previousNode := instance.Vehicles[v].HomeNode
		for _, c := range solution.Routes[v] {
			call := instance.Calls[c]
			if instance.CurrentCalls[c] {
				instance.CurrentCalls[c] = false
				cost += instance.NodeTimesAndAndCosts[v][c].DestinationCost
				cost += instance.TravelTimesAndCosts[v][previousNode][call.DestinationNode].Cost
				previousNode = call.DestinationNode
			} else {
				cost += instance.NodeTimesAndAndCosts[v][c].OriginCost
				cost += instance.TravelTimesAndCosts[v][previousNode][call.OriginNode].Cost
				previousNode = call.OriginNode
				instance.CurrentCalls[c] = true
			}
		}
	}

	// Cost of not transporting
	for _, node := range solution.Routes[instance.NumberOfVehicles] {
		if instance.CurrentCalls[node] {
			instance.CurrentCalls[node] = false
		} else {
			cost += instance.Calls[node].CostOfNotTransporting
			instance.CurrentCalls[node] = true
		}
	}

	return cost
}
