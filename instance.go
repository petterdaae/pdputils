package pdputils

type Instance struct {
	NumberOfNodes        int
	NumberOfVehicles     int
	NumberOfCalls        int
	Vehicles             []Vehicle
	Calls                []Call
	TravelTimesAndCosts  [][][]TravelTimeAndCost // [vehicle, origin, destination]
	NodeTimesAndAndCosts [][]NodeTimeAndCost     // [vehicle, call]
}

type Vehicle struct {
	Index         int
	HomeNode      int
	StartingTime  int
	Capacity      int
	PossibleCalls []int
}

type Call struct {
	Index                 int
	OriginNode            int
	DestinationNode       int
	Size                  int
	CostOfNotTransporting int
	LowerTimePickup       int
	UpperTimePickup       int
	LowerTimeDelivery     int
	UpperTimeDelivery     int
}

type TravelTimeAndCost struct {
	Time int
	Cost int
}

type NodeTimeAndCost struct {
	OriginTime      int
	OriginCost      int
	DestinationTime int
	DestinationCost int
}
