package contemporary

import (
	"github.com/andrespd99/design_patterns/creational-patterns/abstract_factory/civilization"
)

type CivilizationFactory struct{}

func (f CivilizationFactory) GetHouse() civilization.House {
	return House{}
}
func (f CivilizationFactory) GetMedic() civilization.Medic {
	return Medic{}
}
func (f CivilizationFactory) GetTransport() civilization.Transport {
	return GasEngineVehicle{}
}
