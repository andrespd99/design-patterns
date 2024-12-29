package factory

type ReportFactory struct {
	// Create creates a new Report object that will be used on GenerateReport()
	create func() Report
}

func (rf *ReportFactory) GenerateReport() {
	report := rf.create()
	report.Generate()
}

func NewReportFactory(create func() Report) ReportFactory {
	return ReportFactory{
		create: create,
	}
}
