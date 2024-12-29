package impl

import "github.com/andrespd99/design_patterns/creational-patterns/factory"

func NewSalesReportFactory() factory.ReportFactory {
	return factory.NewReportFactory(NewSalesReport)
}
