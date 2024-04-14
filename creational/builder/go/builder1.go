package main

// BUILDER INTERFACE
type IHouseBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IHouseBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}

// CONCRETE BUILDER: For Normal House Product
type NormalBuilder struct {
	windowType, doorType string
	floor                int
}

func newNormalBuilder() *NormalBuilder  { return &NormalBuilder{} }
func (b *NormalBuilder) setWindowType() { b.windowType = "Wooden Window" }
func (b *NormalBuilder) setDoorType()   { b.doorType = "Wooden Door" }
func (b *NormalBuilder) setNumFloor()   { b.floor = 2 }
func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// CONCRETE BUILDER: For Igloo House Product
type IglooBuilder struct {
	windowType, doorType string
	floor                int
}

func newIglooBuilder() *IglooBuilder   { return &IglooBuilder{} }
func (b *IglooBuilder) setWindowType() { b.windowType = "Snow Window" }
func (b *IglooBuilder) setDoorType()   { b.doorType = "Snow Door" }
func (b *IglooBuilder) setNumFloor()   { b.floor = 1 }
func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// PRODUCT: House
type House struct {
	windowType string
	doorType   string
	floor      int
}

// DIRECTOR
type Director struct {
	builder IHouseBuilder
}

func newDirector(b IHouseBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IHouseBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
