package pdputils

func Cost(instance *Instance, solution *Solution) int {
	cost := 0
	currentCalls := make([]bool, instance.NumberOfCalls)

	// Node and travel costs
	for v := 0; v < instance.NumberOfVehicles; v++ {
		previousNode := instance.Vehicles[v].HomeNode
		ln := len(solution.Routes[v])
		for i := 0; i < ln; i++ {
			call := instance.Calls[solution.Routes[v][i]]
			if currentCalls[solution.Routes[v][i]] {
				currentCalls[solution.Routes[v][i]] = false
				cost += instance.NodeTimesAndAndCosts[v][solution.Routes[v][i]].DestinationCost
				cost += instance.TravelTimesAndCosts[v][previousNode][call.DestinationNode].Cost
				previousNode = call.DestinationNode
			} else {
				cost += instance.NodeTimesAndAndCosts[v][solution.Routes[v][i]].OriginCost
				cost += instance.TravelTimesAndCosts[v][previousNode][call.OriginNode].Cost
				previousNode = call.OriginNode
				currentCalls[solution.Routes[v][i]] = true
			}
		}
	}

	// Cost of not transporting
	ln := len(solution.Routes[instance.NumberOfVehicles])
	for i := 0; i < ln; i++ {
		if currentCalls[solution.Routes[instance.NumberOfVehicles][i]] {
			currentCalls[solution.Routes[instance.NumberOfVehicles][i]] = false
		} else {
			cost += instance.Calls[solution.Routes[instance.NumberOfVehicles][i]].CostOfNotTransporting
			currentCalls[solution.Routes[instance.NumberOfVehicles][i]] = true
		}
	}

	return cost
}
