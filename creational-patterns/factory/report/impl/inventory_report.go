package impl

import (
	"fmt"

	"github.com/andrespd99/design_patterns/creational-patterns/factory/report"
)

type InventoryReport struct{}

func (rp *InventoryReport) Generate() {
	fmt.Println("Generating inventory report...")
}

func NewInventoryReport() report.Report {
	return &InventoryReport{}
}
