package medieval

import "fmt"

type Horse struct{}

func (t Horse) Ready() {
	fmt.Println("Put saddle on")
}

func (t Horse) Transport() {
	fmt.Println("Giddy up! 🐴")
}
