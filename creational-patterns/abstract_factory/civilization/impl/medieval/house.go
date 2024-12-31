package medieval

import "fmt"

type CobblestoneHouse struct{}

func (h CobblestoneHouse) GetCozy() {
	fmt.Println("Light fireplace 🪵")
}
