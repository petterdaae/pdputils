package pdputils

func IsFeasible(instance *Instance, solution *Solution) bool {
	for i := 0; i < instance.NumberOfVehicles; i++ {
		vehicle := instance.Vehicles[i]

		// 1. Check that route is compatible with vehicle
		for _, call := range solution.Routes[i] {
			if !instance.Compatibility[i][call] {
				return false
			}
		}

		// 2. Check vehicle capacity
		currentLoad := 0
		for _, node := range solution.Routes[i] {
			call := instance.Calls[node]
			if instance.CurrentCalls[node] {
				instance.CurrentCalls[node] = false
				currentLoad -= call.Size
				continue
			}
			instance.CurrentCalls[node] = true
			currentLoad += call.Size

			if currentLoad > vehicle.Capacity {
				instance.CleanUpCurrentCalls()
				return false
			}
		}

		// 3. Check time windows
		currentNode := vehicle.HomeNode
		currentTime := vehicle.StartingTime
		for _, c := range solution.Routes[i] {
			call := instance.Calls[c]

			if instance.CurrentCalls[c] {
				instance.CurrentCalls[c] = false
				currentTime += instance.TravelTimesAndCosts[i][currentNode][call.DestinationNode].Time
				currentNode = call.DestinationNode

				if currentTime > call.UpperTimeDelivery {
					instance.CleanUpCurrentCalls()
					return false
				}

				if currentTime < call.LowerTimeDelivery {
					currentTime = call.LowerTimeDelivery
				}

				currentTime += instance.NodeTimesAndAndCosts[i][c].DestinationTime
			} else {
				instance.CurrentCalls[c] = true
				currentTime += instance.TravelTimesAndCosts[i][currentNode][call.OriginNode].Time
				currentNode = call.OriginNode

				if currentTime > call.UpperTimePickup {
					instance.CleanUpCurrentCalls()
					return false
				}

				if currentTime < call.LowerTimePickup {
					currentTime = call.LowerTimePickup
				}

				currentTime += instance.NodeTimesAndAndCosts[i][c].OriginTime
			}
		}
	}

	return true
}
