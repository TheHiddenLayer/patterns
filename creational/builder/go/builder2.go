package main

// Concrete Product
type GPU struct {
	memory, cudaCores int
	clockSpeed        float32
	manufacturer      string
}

type IGPUBuilder interface {
	Manufacturer()
	Memory()
	CUDACores()
	ClockSpeed()
	Build() *GPU
}

func NewGPUBuilder(manufacturer string) IGPUBuilder {
	switch manufacturer {
	case "Asus":
		return &AsusGPUBuilder{}
	case "AMD":
		return &AMDGPUBuilder{}
	default:
		return nil
	}
}

// Concrete Builder for Concrete Product: Asus GPU
type AsusGPUBuilder struct {
	memory, cudaCores int
	clockSpeed        float32
	manufacturer      string
}

func (agb *AsusGPUBuilder) Manufacturer() { agb.manufacturer = "Asus" }
func (agb *AsusGPUBuilder) Memory()       { agb.memory = 8 }
func (agb *AsusGPUBuilder) CUDACores()    { agb.cudaCores = 20e3 }
func (agb *AsusGPUBuilder) ClockSpeed()   { agb.clockSpeed = 3.5 }
func (agb *AsusGPUBuilder) Build() *GPU {
	return &GPU{
		manufacturer: agb.manufacturer,
		memory:       agb.memory,
		cudaCores:    agb.cudaCores,
		clockSpeed:   agb.clockSpeed,
	}
}

// Concrete Builder for Concrete Product: AMD GPU
type AMDGPUBuilder struct {
	memory, cudaCores int
	clockSpeed        float32
	manufacturer      string
}

func (agb *AMDGPUBuilder) Manufacturer() { agb.manufacturer = "Advanced Micro Devices" }
func (agb *AMDGPUBuilder) Memory()       { agb.memory = 16 }
func (agb *AMDGPUBuilder) CUDACores()    { agb.cudaCores = 15e3 }
func (agb *AMDGPUBuilder) ClockSpeed()   { agb.clockSpeed = 4 }
func (agb *AMDGPUBuilder) Build() *GPU {
	return &GPU{
		manufacturer: agb.manufacturer,
		memory:       agb.memory,
		cudaCores:    agb.cudaCores,
		clockSpeed:   agb.clockSpeed,
	}
}

// Director
type GPUConstructionDirector struct {
	gb IGPUBuilder
}

func NewGPUConstructionDirector(gb IGPUBuilder) *GPUConstructionDirector {
	return &GPUConstructionDirector{gb: gb}
}
func (d *GPUConstructionDirector) SetBuilder(gb IGPUBuilder) { d.gb = gb }
func (d *GPUConstructionDirector) Build() *GPU {
	d.gb.Manufacturer()
	d.gb.Memory()
	d.gb.CUDACores()
	d.gb.ClockSpeed()
	return d.gb.Build()
}
