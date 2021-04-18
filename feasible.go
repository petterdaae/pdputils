package pdputils

func IsFeasible(instance *Instance, solution *Solution) bool {
	currentCalls := make([]bool, instance.NumberOfCalls)
	for i := 0; i < instance.NumberOfVehicles; i++ {
		vehicle := instance.Vehicles[i]
		ln := len(solution.Routes[i])

		// 1. Check that route is compatible with vehicle
		for j := 0; j < ln; j++ {
			if !instance.Compatibility[i][solution.Routes[i][j]] {
				return false
			}
		}

		// 2. Check vehicle capacity
		currentLoad := 0
		for j := 0; j < ln; j++ {
			call := instance.Calls[solution.Routes[i][j]]
			if currentCalls[solution.Routes[i][j]] {
				currentCalls[solution.Routes[i][j]] = false
				currentLoad -= call.Size
				continue
			}
			currentCalls[solution.Routes[i][j]] = true
			currentLoad += call.Size

			if currentLoad > vehicle.Capacity {
				return false
			}
		}

		// 3. Check time windows
		currentNode := vehicle.HomeNode
		currentTime := vehicle.StartingTime
		for j := 0; j < ln; j++ {
			call := instance.Calls[solution.Routes[i][j]]

			if currentCalls[solution.Routes[i][j]] {
				currentCalls[solution.Routes[i][j]] = false
				currentTime += instance.TravelTimesAndCosts[i][currentNode][call.DestinationNode].Time
				currentNode = call.DestinationNode

				if currentTime > call.UpperTimeDelivery {
					return false
				}

				if currentTime < call.LowerTimeDelivery {
					currentTime = call.LowerTimeDelivery
				}

				currentTime += instance.NodeTimesAndAndCosts[i][solution.Routes[i][j]].DestinationTime
			} else {
				currentCalls[solution.Routes[i][j]] = true
				currentTime += instance.TravelTimesAndCosts[i][currentNode][call.OriginNode].Time
				currentNode = call.OriginNode

				if currentTime > call.UpperTimePickup {
					return false
				}

				if currentTime < call.LowerTimePickup {
					currentTime = call.LowerTimePickup
				}

				currentTime += instance.NodeTimesAndAndCosts[i][solution.Routes[i][j]].OriginTime
			}
		}
	}

	return true
}
