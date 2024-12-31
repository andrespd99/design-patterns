package contemporary

import "fmt"

type Medic struct{}

func (m Medic) Tend() {
	fmt.Println("Stitch wound up and give antibiotics 💊")
}
