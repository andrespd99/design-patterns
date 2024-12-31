package abstractfactory

import (
	"github.com/andrespd99/design_patterns/creational-patterns/abstract_factory/civilization"
)

type CivilizationFactory interface {
	GetHouse() civilization.House
	GetMedic() civilization.Medic
	GetTransport() civilization.Transport
}
