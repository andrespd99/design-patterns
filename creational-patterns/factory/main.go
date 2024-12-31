package factory

func FactoryExampleMain() {
	inventoryReport, _ := getReport("inventory")
	salesReport, _ := getReport("sales")

	inventoryReport.Generate()
	salesReport.Generate()
}
