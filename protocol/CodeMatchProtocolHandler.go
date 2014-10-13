package protocol

import (
	"github.com/thingfu/hub/api"
	"log"
	"github.com/thingfu/hub/utils"
	"time"
	"strings"
	"errors"
)

type CodeMatchProtocolHandler struct {
	api.BaseProtocolHandler
}

func (cm *CodeMatchProtocolHandler) OnStart() {

}

func (cm *CodeMatchProtocolHandler) OnStop()  {

}

func (cm *CodeMatchProtocolHandler) GetName() string {
	return "CodeMatch"
}

func (cm *CodeMatchProtocolHandler) GetLabel() string {
	return "Code Match"
}

func (cm *CodeMatchProtocolHandler) OnRead(data api.ReadRequest) {
	parsed, err := utils.ParseThingFuSerialData(data.GetPayload().(string))

	// Probably not the right thing we're looking for
	if parsed["Protocol"] != "1" {
		return
	}

	if err != nil {
		log.Println(err)
		return
	}
	ser := parsed["Data"].(string)
	if ser == "" {
		log.Println("Something is wrong. Code is nil. Skipping..")
		return
	}

	thing, service, ok := cm.getThing(ser)

	if ok != nil {
		log.Println("Unknown Thing ", ser)
	} else {
		thingManager := cm.GetThingManager()
		t, err := thingManager.GetThingType(thing.Type)
		if err != nil {
			log.Println(err)
		}

		drv := cm.GetFactory().CreateThingAdapter(t.TypeId)
		if drv == nil {
			log.Println("No adapter for thing type " + thing.Type)
			return
		}

		// Sense and Handle Thing Event
		go func() {
			handler := thingManager.GetProtocolHandlerForThing(thing)
			state := drv.OnRead(thing, service, data, handler)

			// We don't want to run rules or fire events too frequently,
			// so check against thing descriptor's Event Update Buffer
			// if we should go ahead
			lastEvent := service.LastEvent
			desc := thing.Descriptor
			if utils.TimeWithinThreshold(lastEvent, desc.EventUpdateBuffer, 5000) {
				service.UpdateLastEvent(time.Now())
				thing.UpdateService(service)
				thingManager.SaveThing(*thing)

				thingManager.Handle(thing, service, state)
			}
		}()
	}
}

func (p *CodeMatchProtocolHandler) Write(t *api.Thing, req api.WriteRequest) {

}

func (p *CodeMatchProtocolHandler) getThing(ser string) (*api.Thing, *api.ThingService, error) {
	thingManager := p.GetThingManager()
	things := thingManager.GetThings()

	for i, _ := range things {
		thing := &things[i]
		desc := thing.Descriptor
		if desc.Protocol == "CodeMatch" {
			attrs := thing.GetAttributeValues("^code_")

			if len(attrs) > 0 {
				for _, item := range attrs {
					name := item.Name
					code := item.Value

					if code == ser {
						thingType, _ := thingManager.GetThingType(thing.Type)
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
