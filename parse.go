package pdputils

import (
	"io/ioutil"
	"strings"
)

func ParseInstance(filename string) Instance {
	data, err := ioutil.ReadFile(filename)
	check(err)

	input := string(data)
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	// Parse counts
	numberOfNodes := unsafeParse(lines[1])
	numberOfVehicles := unsafeParse(lines[3])
	numberOfCalls := unsafeParse(lines[numberOfVehicles+5+1])

	// Parse vehicles
	vehicles := make([]Vehicle, numberOfVehicles)
	offset1 := 5
	offset2 := 8 + numberOfVehicles
	for i := 0; i < numberOfVehicles; i++ {
		line1 := strings.Split(lines[offset1], ",")
		index := unsafeParse(line1[0]) - 1
		vehicles[index] = Vehicle{
			Index:        index,
			HomeNode:     unsafeParse(line1[1]) - 1,
			StartingTime: unsafeParse(line1[2]),
			Capacity:     unsafeParse(line1[3]),
		}
		line2 := strings.Split(lines[offset2], ",")[1:]
		for _, elem := range line2 {
			call := unsafeParse(elem) - 1
			vehicles[index].PossibleCalls = append(vehicles[index].PossibleCalls, call)
		}
		offset1++
		offset2++
	}

	// Parse calls
	calls := make([]Call, numberOfCalls)
	offset := 9 + (2 * numberOfVehicles)
	for i := 0; i < numberOfCalls; i++ {
		line := strings.Split(lines[offset], ",")
		index := unsafeParse(line[0]) - 1
		calls[index] = Call{
			Index:                 index,
			OriginNode:            unsafeParse(line[1]) - 1,
			DestinationNode:       unsafeParse(line[2]) - 1,
			Size:                  unsafeParse(line[3]),
			CostOfNotTransporting: unsafeParse(line[4]),
			LowerTimePickup:       unsafeParse(line[5]),
			UpperTimePickup:       unsafeParse(line[6]),
			LowerTimeDelivery:     unsafeParse(line[7]),
			UpperTimeDelivery:     unsafeParse(line[8]),
		}
		offset++
	}

	// Parse travel times and costs
	travelTimesAndCosts := make([][][]TravelTimeAndCost, numberOfVehicles)
	for i := 0; i < numberOfVehicles; i++ {
		travelTimesAndCosts[i] = make([][]TravelTimeAndCost, numberOfNodes)
		for j := 0; j < numberOfNodes; j++ {
			travelTimesAndCosts[i][j] = make([]TravelTimeAndCost, numberOfNodes)
		}
	}
	offset = 10 + (numberOfVehicles * 2) + numberOfCalls
	for i := 0; i < numberOfVehicles*numberOfNodes*numberOfNodes; i++ {
		line := strings.Split(lines[offset], ",")
		vehicle := unsafeParse(line[0]) - 1
		origin := unsafeParse(line[1]) - 1
		destination := unsafeParse(line[2]) - 1
		travelTimesAndCosts[vehicle][origin][destination] = TravelTimeAndCost{
			Time: unsafeParse(line[3]),
			Cost: unsafeParse(line[4]),
		}
		offset++
	}

	// Parse node times and costs
	nodeTimesAndCosts := make([][]NodeTimeAndCost, numberOfVehicles)
	for i := 0; i < numberOfVehicles; i++ {
		nodeTimesAndCosts[i] = make([]NodeTimeAndCost, numberOfCalls)
	}
	offset++
	for i := 0; i < numberOfVehicles*numberOfCalls; i++ {
		line := strings.Split(lines[offset], ",")
		vehicle := unsafeParse(line[0]) - 1
		call := unsafeParse(line[1]) - 1
		nodeTimesAndCosts[vehicle][call] = NodeTimeAndCost{
			OriginTime:      unsafeParse(line[2]),
			OriginCost:      unsafeParse(line[3]),
			DestinationTime: unsafeParse(line[4]),
			DestinationCost: unsafeParse(line[5]),
		}
		offset++
	}

	return Instance{
		NumberOfNodes:        numberOfNodes,
		NumberOfVehicles:     numberOfVehicles,
		NumberOfCalls:        numberOfCalls,
		Vehicles:             vehicles,
		Calls:                calls,
		TravelTimesAndCosts:  travelTimesAndCosts,
		NodeTimesAndAndCosts: nodeTimesAndCosts,
	}
}
