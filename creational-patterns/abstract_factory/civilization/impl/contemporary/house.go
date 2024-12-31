package contemporary

import "fmt"

type House struct{}

func (h House) GetCozy() {
	fmt.Println("Turn on electric heater and then some Netflix & chill")
}
