package factory

import (
	"github.com/go-home/hub/api"
	_ "github.com/go-home/hub/rules/conditions"
	"github.com/go-home/hub/tests"
	"testing"
)

func TestFactoryConditions(t *testing.T) {
	f := new(DefaultFactory)

	if f.CreateCondition("xoxo") != nil {
		t.Errorf("Condition not Expected")
	}

	if f.CreateCondition("sense") == nil {
		t.Errorf("'Sense' Condition Expected")
	}

	if f.CreateCondition("hourly") == nil {
		t.Errorf("'Hourly' Condition Expected")
	}
}

func TestFactoryConsequences(t *testing.T) {
	f := new(DefaultFactory)

	if f.CreateConsequence("xoxo") != nil {
		t.Errorf("Consequence not Expected")
	}

	if f.CreateConsequence("sendmail") == nil {
		t.Errorf("'SendMail' Consequence Expected")
	}

	if f.CreateConsequence("logwrite") == nil {
		t.Errorf("'LogWrite' Consequence Expected")
	}
}

func TestFactoryProtocolHandlers(t *testing.T) {
	f := new(DefaultFactory)
	f.SetContainer(tests.NewMockContainer())
	cfg := new(api.ProtocolConfiguration)

	if f.CreateProtocolHandler("xoxo", *cfg) != nil {
		t.Errorf("Protocol Handler not Expected")
	}

	if f.CreateConsequence("RF433") == nil {
		t.Errorf("'RF433' Protocol Handler Expected")
	}

	if f.CreateConsequence("http") == nil {
		t.Errorf("'HTTP' Protocol Handler Expected")
	}

	if f.CreateConsequence("zigbee") == nil {
		t.Errorf("'Zigbee' Protocol Handler Expected")
	}

	if f.CreateConsequence("sim") == nil {
		t.Errorf("'Simulating' Protocol Handler Expected")
	}
}

func TestFactoryThingAdapters(t *testing.T) {
	f := new(DefaultFactory)

	if f.CreateThingAdapter("xoxo") != nil {
		t.Errorf("Adapter not Expected")
	}

	if f.CreateThingAdapter("lgtv-47ls5700") != nil {
		t.Errorf("'LGTV 47LS5700' Adapter Expected")
	}

	if f.CreateThingAdapter("433mhz-contact") != nil {
		t.Errorf("'433Mhz Contact' Adapter Expected")
	}

	if f.CreateThingAdapter("433mhz-1button") != nil {
		t.Errorf("'433Mhz 1 Button' Adapter Expected")
	}

	if f.CreateThingAdapter("433mhz-wt450") != nil {
		t.Errorf("'433Mhz DHT' Adapter Expected")
	}

	if f.CreateThingAdapter("433mhz-knock") != nil {
		t.Errorf("'433Mhz Knock' Adapter Expected")
	}

	if f.CreateThingAdapter("433mhz-motion") != nil {
		t.Errorf("'433Mhz Motion' Adapter Expected")
	}

	if f.CreateThingAdapter("dlink-dcs930l") != nil {
		t.Errorf("'DCS930L IP Cam' Adapter Expected")
	}

	if f.CreateThingAdapter("433mhz-4buttons") != nil {
		t.Errorf("'433Mhz 4-Buttons' Adapter Expected")
	}

	if f.CreateThingAdapter("weather") != nil {

	}
}

/*
func (s *DefaultFactory) CreateThingAdapter(t string) api.ThingAdapter {
	var adapter api.ThingAdapter
	switch t {
	case "lgtv-47ls5700":
		adapter = new(adapters.AdapterLGTV47LS5700)

	case "433mhz-contact":
		adapter = new(adapters.AdapterContactSensor433)

	case "433mhz-1button":
		adapter = new(adapters.Adapter1ButtonWireless433)

	case "433mhz-wt450":
		adapter = new(adapters.AdapterDigitalHumidityTemperature433)

	case "433mhz-knock":
		adapter = new(adapters.AdapterKnockSensor433)

	case "433mhz-motion":
		adapter = new(adapters.AdapterMotionSensor433)

	case "dlink-dcs930l":
		adapter = new(adapters.AdapterIPCamera)

	case "433mhz-4buttons":
		adapter = new(adapters.Adapter4ButtonWireless433)

	case "weather":
		adapter = new(adapters.AdapterWeather)

	default:
		return nil
	}
	return adapter
}

*/
