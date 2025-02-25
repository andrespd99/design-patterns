package impl

import (
	"fmt"

	"github.com/andrespd99/design_patterns/creational-patterns/factory/report"
)

type SalesReport struct{}

func (rp *SalesReport) Generate() {
	fmt.Println("Generating sales report...")
}

func NewSalesReport() report.Report {
	return &SalesReport{}
}
