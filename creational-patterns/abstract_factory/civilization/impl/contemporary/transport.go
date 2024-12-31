package contemporary

import "fmt"

type GasEngineVehicle struct{}

func (t GasEngineVehicle) Ready() {
	fmt.Println("Turn engine on")
}

func (t GasEngineVehicle) Transport() {
	fmt.Println("Brummmmmmm 🚗")
}
