package impl

import "github.com/andrespd99/design_patterns/creational-patterns/factory"

func NewInventoryReportFactory() factory.ReportFactory {
	return factory.NewReportFactory(NewInventoryReport)
}
