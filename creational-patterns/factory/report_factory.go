package factory

import (
	"fmt"

	"github.com/andrespd99/design_patterns/creational-patterns/factory/report"
	"github.com/andrespd99/design_patterns/creational-patterns/factory/report/impl"
)

func getReport(reportType string) (report.Report, error) {
	if reportType == "inventory" {
		return impl.NewInventoryReport(), nil
	}
	if reportType == "sales" {
		return impl.NewSalesReport(), nil
	}
	return nil, fmt.Errorf("unknown report type: %s", reportType)
}
