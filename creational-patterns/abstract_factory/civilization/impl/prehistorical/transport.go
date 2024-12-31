package prehistorical

import "fmt"

type GoodOldFoot struct{}

func (t GoodOldFoot) Ready() {
	fmt.Println("Ready up your happy feet!")
}

func (t GoodOldFoot) Transport() {
	fmt.Println("Literally just walk, lol")
}
