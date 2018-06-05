package mage

import (
	"fmt"
)

type Destination struct {
	value string
}

func DestinationForMac() Destination {
	return Destination{"'arch=x86_64'"}
}

func DestinationForSimulator(osVersion string, simulator string) Destination {
	return Destination{fmt.Sprintf("'OS=%s,name=%s'", osVersion, simulator)}
}

func DestinationGeneric(platform Platform) Destination {
	return Destination{fmt.Sprintf("'generic/platform=%s'", string(platform))}
}

func (d Destination) String() string {
	return d.value
}
