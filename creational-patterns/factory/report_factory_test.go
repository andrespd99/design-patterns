package factory

import (
	"fmt"
	"testing"

	report "github.com/andrespd99/design_patterns/creational-patterns/factory/report/impl"
	"github.com/stretchr/testify/assert"
)

func TestGetReport(t *testing.T) {
	t.Run("should return InventoryReport on \"inventory\" reportType input", func(t *testing.T) {
		reportType := "inventory"
		rp, err := getReport(reportType)
		if err != nil {
			assert.Fail(t, fmt.Sprintf("got error when a result was expected from \"%s\" reportType", reportType))
		}

		assert.IsType(t, &report.InventoryReport{}, rp)
	})

	t.Run("should return InventoryReport on \"sales\" reportType input", func(t *testing.T) {
		reportType := "sales"
		rp, err := getReport(reportType)
		if err != nil {
			assert.Fail(t, fmt.Sprintf("got error when a result was expected from \"%s\" reportType", reportType))
		}

		assert.IsType(t, &report.SalesReport{}, rp)
	})
	t.Run("should return an error when an unknown reportType is given", func(t *testing.T) {
		reportType := "foo"
		rp, err := getReport(reportType)
		if err == nil {
			assert.Fail(t, "got nil when error was expected")
		}
		if rp != nil {
			assert.Fail(t, "got Report when nil was expected")
		}
		expectedErr := fmt.Errorf("unknown report type: %s", reportType)
		assert.Equal(t, expectedErr, err)
	})
}
