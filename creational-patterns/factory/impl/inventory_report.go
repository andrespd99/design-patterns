package impl

import (
	"fmt"

	"github.com/andrespd99/design_patterns/creational-patterns/factory"
)

type InventoryReport struct{}

func (rp *InventoryReport) Generate() {
	fmt.Println("Generating inventory report...")
}

func NewInventoryReport() factory.Report {
	return &InventoryReport{}
}
