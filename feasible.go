package pdputils

import (
	"fmt"
)

func IsFeasible(instance *Instance, solution []int) (bool, string) {
	startIndex := 0
	zeroIndices := getZeroIndices(solution)
	for i := 0; i < instance.NumberOfVehicles; i++ {
		vehicle := instance.Vehicles[i]
		route := solution[startIndex:zeroIndices[i]]

		// 1. Check that route is compatible with vehicle
		for _, call := range route {
			if !instance.Compatibility[i][call] {
				return false, fmt.Sprintf("incompatible vehicle and call (%d, %d)", i, call)
			}
		}

		// 2. Check vehicle capacity
		currentLoad := 0
		currentCalls := make(map[int]bool)
		for _, node := range route {
			call := instance.Calls[node]
			if value, present := currentCalls[node]; present && value {
				currentCalls[node] = false
				currentLoad -= call.Size
				continue
			}
			currentCalls[node] = true
			currentLoad += call.Size

			if currentLoad > vehicle.Capacity {
				return false, "vehicle capacity exceeded"
			}
		}

		// 3. Check time windows
		currentNode := vehicle.HomeNode
		currentTime := vehicle.StartingTime
		currentCalls = make(map[int]bool)
		for _, c := range route {
			call := instance.Calls[c]

			if value, present := currentCalls[c]; present && value {
				currentCalls[c] = false
				currentTime += instance.TravelTimesAndCosts[i][currentNode][call.DestinationNode].Time
				currentNode = call.DestinationNode

				if currentTime > call.UpperTimeDelivery {
					return false, "delivery time exceeded"
				}

				if currentTime < call.LowerTimeDelivery {
					currentTime = call.LowerTimeDelivery
				}

				currentTime += instance.NodeTimesAndAndCosts[i][c].DestinationTime
			} else {
				currentCalls[c] = true
				currentTime += instance.TravelTimesAndCosts[i][currentNode][call.OriginNode].Time
				currentNode = call.OriginNode

				if currentTime > call.UpperTimePickup {
					return false, "pickup time exceeded"
				}

				if currentTime < call.LowerTimePickup {
					currentTime = call.LowerTimePickup
				}

				currentTime += instance.NodeTimesAndAndCosts[i][c].OriginTime
			}
		}

		startIndex = zeroIndices[i] + 1
	}

	return true, "feasible"
}
