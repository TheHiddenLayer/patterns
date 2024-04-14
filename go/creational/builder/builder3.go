package main

// IMobilePhoneBuilder interface
type IMobilePhoneBuilder interface {
	Manufacturer()
	Chip()
	Telemetry(level string)
	OS()
	Camera()
	Screen()
	Battery()
	Build() *MobilePhone
	Extras()
}

type mobilePhoneExtras []map[string]string

// MobilePhone struct
type MobilePhone struct {
	manufacturer, chip, telemetry, os, screen string
	camera, battery                           int
	extras                                    mobilePhoneExtras
}

// Concrete Product: iPhone
type iPhone struct{ MobilePhone }

// Concrete Product: Samsung Galaxy
type SamsungGalaxy struct{ MobilePhone }

// Concrete Builder: IPhoneBuilder
type IPhoneBuilder struct{ iPhone }

func (i *IPhoneBuilder) Manufacturer()          { i.manufacturer = "Apple" }
func (i *IPhoneBuilder) Chip()                  { i.chip = "Apple Silicon" }
func (i *IPhoneBuilder) Telemetry(level string) { i.telemetry = level }
func (i *IPhoneBuilder) OS()                    { i.os = "iOS" }
func (i *IPhoneBuilder) Screen()                { i.screen = "Retina Display" }
func (i *IPhoneBuilder) Camera()                { i.camera = 3 }
func (i *IPhoneBuilder) Battery()               { i.battery = 10 }
func (i *IPhoneBuilder) Extras() {
	i.extras = []map[string]string{
		{"key": "Apple Care", "value": "2 years"},
		{"key": "Free", "value": "EarPods"},
	}
}
func (i *IPhoneBuilder) Build() *MobilePhone {
	return &MobilePhone{
		manufacturer: i.manufacturer,
		chip:         i.chip,
		telemetry:    i.telemetry,
		os:           i.os,
		screen:       i.screen,
		camera:       i.camera,
		battery:      i.battery,
		extras:       i.extras,
	}
}

// Concrete Builder: SamsungGalaxyBuilder
type SamsungGalaxyBuilder struct{ SamsungGalaxy }

func (g *SamsungGalaxyBuilder) Manufacturer()          { g.manufacturer = "Samsung" }
func (g *SamsungGalaxyBuilder) Chip()                  { g.chip = "Qualcomm Snapdragon" }
func (g *SamsungGalaxyBuilder) Telemetry(level string) { g.telemetry = level }
func (g *SamsungGalaxyBuilder) OS()                    { g.os = "Android" }
func (g *SamsungGalaxyBuilder) Screen()                { g.screen = "OLED" }
func (g *SamsungGalaxyBuilder) Camera()                { g.camera = 4 }
func (g *SamsungGalaxyBuilder) Battery()               { g.battery = 9 }
func (g *SamsungGalaxyBuilder) Extras() {
	g.extras = []map[string]string{
		{"key": "Expandable Storage", "value": "10 GB"},
	}
}
func (g *SamsungGalaxyBuilder) Build() *MobilePhone {
	return &MobilePhone{
		manufacturer: g.manufacturer,
		chip:         g.chip,
		telemetry:    g.telemetry,
		os:           g.os,
		screen:       g.screen,
		camera:       g.camera,
		battery:      g.battery,
		extras:       g.extras,
	}
}

func NewMobilePhoneBuilder(brand string) IMobilePhoneBuilder {
	switch brand {
	case "Apple":
		return &IPhoneBuilder{}
	case "Samsung":
		return &SamsungGalaxyBuilder{}
	default:
		return nil
	}
}

// Director
type MobilePhoneDirector struct {
	mpb IMobilePhoneBuilder
}

func NewMobilePhoneDirector(b IMobilePhoneBuilder) *MobilePhoneDirector {
	return &MobilePhoneDirector{mpb: b}
}

func (mpd *MobilePhoneDirector) SetBuilder(b IMobilePhoneBuilder) {
	mpd.mpb = b
}

func (mpd *MobilePhoneDirector) Build(region string) *MobilePhone {
	mpd.mpb.Manufacturer()
	mpd.mpb.Chip()
	mpd.mpb.OS()
	mpd.mpb.Screen()
	mpd.mpb.Camera()
	mpd.mpb.Battery()
	mpd.mpb.Extras()

	// phones in European Union must abide by GDPR
	telemetryLevel := "unfettered access"
	if region == "EU" {
		telemetryLevel = "limited"
	}
	mpd.mpb.Telemetry(telemetryLevel)

	return mpd.mpb.Build()
}
