package prehistorical

import (
	"github.com/andrespd99/design_patterns/creational-patterns/abstract_factory/civilization"
)

type CivilizationFactory struct{}

func (f CivilizationFactory) GetHouse() civilization.House {
	return Cave{}
}
func (f CivilizationFactory) GetMedic() civilization.Medic {
	return Shaman{}
}
func (f CivilizationFactory) GetTransport() civilization.Transport {
	return GoodOldFoot{}
}
