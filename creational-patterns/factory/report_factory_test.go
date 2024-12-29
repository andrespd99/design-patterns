package factory

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockReport struct {
	mock.Mock
}

func (m *MockReport) Generate() {
	m.Called()
}

func TestReportFactory(t *testing.T) {
	t.Run("should call Report.Generate of inner Report", func(t *testing.T) {
		rpMock := new(MockReport)
		createReport := func() Report {
			return rpMock
		}

		rpMock.On("Generate").Return()

		rf := NewReportFactory(createReport)

		rf.GenerateReport()

		rpMock.AssertNumberOfCalls(t, "Generate", 1)
	})
}
