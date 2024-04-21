package main

import "fmt"

// INTERFACE (all concrete products should impl)
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// A concrete type
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) { g.name = name }
func (g *Gun) setPower(power int)  { g.power = power }
func (g *Gun) getName() string     { return g.name }
func (g *Gun) getPower() int       { return g.power }

// Concrete type that directly embeds Gun
type AK47 struct{ Gun }

func newAK47() IGun { return &AK47{Gun{power: 10, name: "Kalashnikov AK47"}} }

// Concrete type that directly embeds Gun
type Musket struct{ Gun }

func newMusket() IGun { return &Musket{Gun{power: 2, name: "Barrett Musket"}} }

// simpleGunFactory is an idiomatic "simple factory"; it merely separates
// creation of a concrete product from the use of the conrete product
func simpleGunFactory(gunType string) (IGun, error) {
	switch gunType {
	case "AK47":
		return newAK47(), nil
	case "Musket":
		return newMusket(), nil
	default:
		return nil, fmt.Errorf("invalid gun type")
	}
}
