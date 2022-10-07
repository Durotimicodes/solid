package main

type Ibuilder interface {
	setWindowType()
	setFloorType()
}

type NormalBuilder struct {
	window string
	floors string
	doors  string
}

type IgahloBuilding struct {
	windowType string
	floorType  string
	doorType   string
}

func getBuilder(buildingType string) Ibuilder {
	if buildingType == "normal" {
		return &NormalBuilder{}
	} else if buildingType == "other" {
		return &IgahloBuilding{}
	}

	return nil
}


func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (n *NormalBuilder) setFloorType() {
	n.doors = "wood"
}

func (n *NormalBuilder) setWindowType() {
	n.window = "glass"
}



func newIgalooBuilding() *IgahloBuilding {
	return &IgahloBuilding{}
}

func (i *IgahloBuilding) setFloorType() {
	i.doorType = "ceramic"
}

func (i *IgahloBuilding) setWindowType() {
	i.floorType = "mabble"
}
