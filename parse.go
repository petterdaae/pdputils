package pdputils

import (
	"io/ioutil"
	"strings"
)

func ParseInstance(instance *Instance, filename string) {
	data, err := ioutil.ReadFile(filename)
	check(err)

	input := string(data)
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	// Parse counts
	numberOfNodes := unsafeParse(lines[1])
	numberOfVehicles := unsafeParse(lines[3])
	numberOfCalls := unsafeParse(lines[numberOfVehicles+5+1])
	calls := make([]Call, numberOfCalls)
	instance.Compatibility = make([][]bool, numberOfVehicles+1)
	for i := 0; i < numberOfVehicles+1; i++ {
		instance.Compatibility[i] = make([]bool, numberOfCalls)
	}
	for i := 0; i < numberOfCalls; i++ {
		instance.Compatibility[numberOfVehicles][i] = true
		calls[i].PossibleVehicles = append(calls[i].PossibleVehicles, numberOfVehicles)
	}

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
			calls[call].PossibleVehicles = append(calls[call].PossibleVehicles, index)
			instance.Compatibility[index][call] = true
		}
		offset1++
		offset2++
	}

	// Parse calls
	offset := 9 + (2 * numberOfVehicles)
	for i := 0; i < numberOfCalls; i++ {
		line := strings.Split(lines[offset], ",")
		index := unsafeParse(line[0]) - 1
		calls[index].Index = index
		calls[index].OriginNode = unsafeParse(line[1]) - 1
		calls[index].DestinationNode = unsafeParse(line[2]) - 1
		calls[index].Size = unsafeParse(line[3])
		calls[index].CostOfNotTransporting = unsafeParse(line[4])
		calls[index].LowerTimePickup = unsafeParse(line[5])
		calls[index].UpperTimePickup = unsafeParse(line[6])
		calls[index].LowerTimeDelivery = unsafeParse(line[7])
		calls[index].UpperTimeDelivery = unsafeParse(line[8])
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

	instance.NumberOfNodes = numberOfNodes
	instance.NumberOfVehicles = numberOfVehicles
	instance.NumberOfCalls = numberOfCalls
	instance.Vehicles = vehicles
	instance.Calls = calls
	instance.TravelTimesAndCosts = travelTimesAndCosts
	instance.NodeTimesAndAndCosts = nodeTimesAndCosts
}
