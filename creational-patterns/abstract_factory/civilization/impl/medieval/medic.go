package medieval

import "fmt"

type Medic struct{}

func (m Medic) Tend() {
	fmt.Println("Ampute leg and pray to god")
}
