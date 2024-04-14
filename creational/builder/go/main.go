package main

import "fmt"

func builder1() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n\n", iglooHouse.floor)
}

func builder2() {
	asusGpuBuilder := NewGPUBuilder("Asus")

	director := NewGPUConstructionDirector(asusGpuBuilder)
	gpu := director.Build()

	fmt.Println("GPU")
	fmt.Printf("    Manufacturer: %s\n", gpu.manufacturer)
	fmt.Printf("    Memory: %d\n", gpu.memory)
	fmt.Printf("    Clock speed: %f\n", gpu.clockSpeed)
	fmt.Printf("    CUDA cores: %d\n", gpu.cudaCores)

	amdGpuBuilder := NewGPUBuilder("AMD")
	director.SetBuilder(amdGpuBuilder)
	gpu = director.Build()

	fmt.Println("GPU")
	fmt.Printf("    Manufacturer: %s\n", gpu.manufacturer)
	fmt.Printf("    Memory: %d\n", gpu.memory)
	fmt.Printf("    Clock speed: %f\n", gpu.clockSpeed)
	fmt.Printf("    CUDA cores: %d\n", gpu.cudaCores)

}

func builder3() {

	printPhoneMetadata := func(mps *[]*MobilePhone) {
		for _, mp := range *mps {
			fmt.Println("Phone")
			fmt.Printf("      Manufacturer: %s\n", mp.manufacturer)
			fmt.Printf("      OS: %s\n", mp.os)
			fmt.Printf("      Screen: %s\n", mp.screen)
			fmt.Printf("      Camera: %d\n", mp.camera)
			fmt.Printf("      Chip: %s\n", mp.chip)
			fmt.Printf("      Telemetry: %s\n", mp.telemetry)
			fmt.Printf("      Battery: %d\n", mp.battery)
			for _, extra := range mp.extras {
				fmt.Printf("      %s: %s\n", extra["key"], extra["value"])
			}
		}
	}

	// store all phones built
	phones := make([]*MobilePhone, 0, 4)

	b := NewMobilePhoneBuilder("Apple")
	director := NewMobilePhoneDirector(b)
	phone := director.Build("EU")
	phones = append(phones, phone)

	b = NewMobilePhoneBuilder("Samsung")
	director.SetBuilder(b)
	phone = director.Build("EU")
	phones = append(phones, phone)

	b = NewMobilePhoneBuilder("Apple")
	director.SetBuilder(b)
	phone = director.Build("Brazil")
	phones = append(phones, phone)

	b = NewMobilePhoneBuilder("Samsung")
	director.SetBuilder(b)
	phone = director.Build("Korea")
	phones = append(phones, phone)

	printPhoneMetadata(&phones)
}

func main() {
	builder1()
	builder2()
	builder3()
}
