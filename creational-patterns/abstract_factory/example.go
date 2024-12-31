package abstractfactory

import (
	"fmt"

	"github.com/andrespd99/design_patterns/creational-patterns/abstract_factory/civilization/impl/contemporary"
	"github.com/andrespd99/design_patterns/creational-patterns/abstract_factory/civilization/impl/medieval"
	"github.com/andrespd99/design_patterns/creational-patterns/abstract_factory/civilization/impl/prehistorical"
)

func AbstractFactoryExample(civilizationType string) (factory CivilizationFactory, err error) {
	if civilizationType == "contemporary" {
		factory = contemporary.CivilizationFactory{}
		return
	}
	if civilizationType == "medieval" {
		factory = medieval.CivilizationFactory{}
		return
	}
	if civilizationType == "prehistorical" {
		factory = prehistorical.CivilizationFactory{}
		return
	}

	err = fmt.Errorf("unknown civilization type: %s", civilizationType)
	return
}
