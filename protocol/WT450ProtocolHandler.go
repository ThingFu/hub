package protocol

import (
	"github.com/thingfu/hub/api"
	"github.com/thingfu/hub/utils"
	"log"
	"strconv"
	"time"
	"errors"
)

type WT450ProtocolHandler struct {
	api.BaseProtocolHandler
}

func (h *WT450ProtocolHandler) OnStart() {

}

func (h *WT450ProtocolHandler) OnStop()  {

}

func (h *WT450ProtocolHandler) GetName() string {
	return "WT450"
}

func (h *WT450ProtocolHandler) GetLabel() string {
	return "WT450 Digital Humidity and Temperature"
}

func (p *WT450ProtocolHandler) Write(t *api.Thing, req api.WriteRequest) {

}

func (h *WT450ProtocolHandler) OnRead(data api.ReadRequest) {
	parsed, err := utils.ParseThingFuSerialData(data.GetPayload().(string))
	if err != nil {
		log.Println(err)
		return
	}

	val, _ := strconv.Atoi(parsed["Data"].(string))
	ch := val >> 26
	thing, service, err := h.getThing(strconv.Itoa(ch))

	if err == nil {
		factory := h.GetFactory()
		thingManager := h.GetThingManager()

		drv := factory.CreateThingAdapter("433mhz-wt450")

		go func() {
			data.Put("channel", ch)
			data.Put("dhtdata", val)

			handler := thingManager.GetProtocolHandlerForThing(thing)
			state := drv.OnRead(thing, service, data, handler)

			lastEvent := service.LastEvent
			desc := thing.Descriptor
			if utils.TimeWithinThreshold(lastEvent, desc.EventUpdateBuffer, 5000) {
				service.UpdateLastEvent(time.Now())
				thingManager.SaveThing(*thing)

				thingManager.Handle(thing, service, state)
			}
		}()
	}
}

func (p *WT450ProtocolHandler) getThing(ch string) (*api.Thing, *api.ThingService, error) {
	thingManager := p.GetThingManager()
	things := thingManager.GetThings()

	for i, _ := range things {
		thing := &things[i]
		desc := thing.Descriptor
		if desc.Protocol == "WT450" {
			attr := thing.GetAttributeValue("channel")
			val := attr.Value

			if val == ch {
				thingType, _ := thingManager.GetThingType(thing.Type)
				svc := thingType.GetService("DHT")

				return thing, svc, nil
			}
		}
	}
	return new(api.Thing), new(api.ThingService), errors.New("Unknown Thing")
}

