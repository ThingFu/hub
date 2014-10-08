package protocol

import (
	"errors"
	serial "github.com/tarm/goserial"
	"github.com/thingfu/hub/api"
	"github.com/thingfu/hub/utils"
	"log"
	"strconv"
	"strings"
	"time"
)

// Data given by the MCU
type RF433Data struct {
	Protocol int
	BinData  string
	DecData  string
}

func (r *RF433Data) GetData() map[string]interface{} {
	ret := make(map[string]interface{})

	ret["protocol"] = r.Protocol
	ret["bin"] = r.BinData
	ret["dec"] = r.DecData

	return ret
}

// Handles 433MHZ signals via Serial
type RF433ProtocolHandler struct {
	factory      api.Factory
	environment  api.Environment
	thingManager api.ThingManager
	config       api.ProtocolConfiguration
}

func (p *RF433ProtocolHandler) Stop() {

}

func (p *RF433ProtocolHandler) Start() {
	log.Println("[INFO] Starting Serial for RF433 ..")

	c := &serial.Config{Name: p.config.Port, Baud: int(p.config.Baud)}
	s, err := serial.OpenPort(c)
	if err == nil {
		buf := ""

		for {
			serialContent := make([]byte, 256)
			_, err := s.Read(serialContent)
			if err != nil {
				log.Fatal(err)
			}

			c := string(serialContent)
			if strings.Index(c, "\n") != -1 {
				str := strings.Split(c, "\n")
				buf += str[0]
				buf = strings.Replace(buf, "\x00", "", -1)
				p.Handle(buf)
				buf = str[1]
			} else {
				buf += c
			}
		}
	} else {
		log.Print("[ERROR] Unable to start RF433 Serial -- %s", err)
	}
}

func (p *RF433ProtocolHandler) Handle(payload interface{}) {
	buf := payload.(string)
	spl := strings.Split(buf, "|")

	if len(spl) < 5 {
		return
	}

	data := new(RF433Data)
	prot, _ := strconv.Atoi(spl[1])
	data.Protocol = prot
	data.BinData = spl[2]
	data.DecData = spl[3]

	if data.Protocol == 5 && len(data.BinData) == 36 {
		go p.handleWT450(data)
	} else {
		go p.handleCodeMatch(data)
	}
}

func (p *RF433ProtocolHandler) handleWT450(data *RF433Data) {
	dec, _ := strconv.Atoi(data.DecData)
	ser := "nb-wt450-" + strconv.Itoa(dec>>26)

	dev, service, err := p.getThing(ser)

	if err == nil {
		drv := p.factory.CreateThingAdapter("433mhz-wt450")

		go func() {
			state := drv.OnSense(dev, service, data)

			lastEvent := service.LastEvent
			desc := dev.Descriptor
			if utils.TimeWithinThreshold(lastEvent, desc.EventUpdateBuffer, 5000) {
				service.UpdateLastEvent(time.Now())
				p.thingManager.SaveThing(*dev)

				p.thingManager.Handle(dev, service, state)
			}
		}()
	}
}

/**
Matches and handles any RF433 code against a list of 433MHZ Thing Sensors.
*/
func (p *RF433ProtocolHandler) handleCodeMatch(data *RF433Data) {
	ser := data.DecData

	// If serial data gets messed up
	if ser == "" {
		log.Println("Something is wrong. Code from RF 433 service is nil. Skipping..")
		return
	}

	dev, service, ok := p.getThing(ser)

	if ok != nil {
		log.Println("Unknown Thing ", ser)
	} else {
		t, err := p.thingManager.GetThingType(dev.Type)
		if err != nil {
			log.Println(err)
		}

		drv := p.factory.CreateThingAdapter(t.TypeId)
		if drv == nil {
			log.Println("No adapter for thing type " + dev.Type)
			return
		}

		// Sense and Handle Thing Event
		go func() {
			state := drv.OnSense(dev, service, data)

			// We don't want to run rules or fire events too frequently,
			// so check against thing descriptor's Event Update Buffer
			// if we should go ahead
			lastEvent := service.LastEvent
			desc := dev.Descriptor
			if utils.TimeWithinThreshold(lastEvent, desc.EventUpdateBuffer, 5000) {
				service.UpdateLastEvent(time.Now())
				dev.UpdateService(service)
				p.thingManager.SaveThing(*dev)

				p.thingManager.Handle(dev, service, state)
			}
		}()
	}
}

func (p *RF433ProtocolHandler) getThing(ser string) (*api.Thing, *api.ThingService, error) {
	things := p.thingManager.GetThings()

	for i, _ := range things {
		thing := &things[i]
		desc := thing.Descriptor
		if desc.Protocol == "433MHZ" {
			attrs := thing.GetAttributeValues("^code_")

			if len(attrs) > 0 {
				for _, item := range attrs {
					name := item.Name
					code := item.Value

					if code == ser {
						thingType, _ := p.thingManager.GetThingType(thing.Type)
						svc := strings.Replace(name, "code_", "", -1)
						service := thingType.GetService(svc)

						return thing, service, nil
					}
				}
			}
		}
	}
	return new(api.Thing), new(api.ThingService), errors.New("Unknown Thing")
}

func (p *RF433ProtocolHandler) SetFactory(o api.Factory) {
	p.factory = o
}

func (p *RF433ProtocolHandler) SetThingManager(o api.ThingManager) {
	p.thingManager = o
}

func (p *RF433ProtocolHandler) SetEnvironment(o api.Environment) {
	p.environment = o
}

func (p *RF433ProtocolHandler) SetProtocolConfiguration(o api.ProtocolConfiguration) {
	p.config = o
}

func (p *RF433ProtocolHandler) IsEnabled() bool {
	return p.config.Enabled
}

func (p *RF433ProtocolHandler) GetName() string {
	return "433MHZ"
}

func (p *RF433ProtocolHandler) GetLabel() string {
	return "RF 433MHZ"
}

func (p *RF433ProtocolHandler) SetContainer(api.Container) {

}

func (p *RF433ProtocolHandler) ValidateWiring() {

}
