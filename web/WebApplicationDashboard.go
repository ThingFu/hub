// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package web

import (
	"bytes"
	"fmt"
	"github.com/go-home/hub/api"
	"github.com/go-home/hub/utils"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var funcMap = template.FuncMap{
	"Ago": func(_t time.Time) string {
		gt := new(utils.GoTime)
		gt.SetTime(_t)
		return gt.Ago()
	},
}

type WebApplicationDashboard struct {
	rulesService  api.RulesService
	thingService api.ThingService
	dataSource    api.DataSource
	environment   api.Environment
	factory       api.Factory
	container     api.Container
}

func (app *WebApplicationDashboard) Setup(r *mux.Router) {
	r.HandleFunc("/rules", app.handleManageRules)
	r.HandleFunc("/rule/{id}", app.handleEditRules)
	r.HandleFunc("/settings", app.handleSettingsView)
	r.HandleFunc("/events", app.handleEventsView)
	r.HandleFunc("/sim/event/{protocol}", app.handleSimulationService).Methods("POST")
	// r.HandleFunc("/widget/{thingId}/configure", app.handleWidgetConfigure).Methods("GET")
	r.HandleFunc("/widget/{thingId}/configure", app.handleWidgetUpdateConfiguration).Methods("POST")
	r.HandleFunc("/widget/{thingId}/view", app.handleWidgetView)
	r.HandleFunc("/thing/{thingType}/resource/icon/128x", app.handleResourceIcon)
	r.HandleFunc("/thing/add", app.handleThingAdd)
	r.HandleFunc("/thing/add/{typeId}", app.handleThingAddNew).Methods("POST", "GET")
	r.HandleFunc("/things", app.handleThings)
	r.HandleFunc("/sysinfo", app.handleSysInfo)
	r.HandleFunc("/about", app.handleAbout)
}

func renderContent(path string, model interface{}) template.HTML {
	fileContent, _ := ioutil.ReadFile(path)
	stringContent := string(fileContent)

	tmpl, err := template.New("__tpl_" + path).Parse(stringContent)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}

	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, model)

	htmlContent := buf.String()
	return template.HTML(htmlContent)
}

func (app *WebApplicationDashboard) handleThingAddNew(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	typeId := vars["typeId"]
	if req.Method == "GET" {
		thing := app.thingService.GetThingType(typeId)
		thingType := app.thingService.GetThingType(typeId)

		model := new(webModelThingAddNew)
		model.AddNewContent = renderContent(thing.Path+"/add.html", thingType)

		w.Write(templateOutput("thing_addnew", model))
	} else if req.Method == "POST" {
		//
		body, _ := ioutil.ReadAll(req.Body)
		content := string(body)

		/*
			{
				"_id" : ObjectId("540f33e9ffe79223bcb81706"),
				"uid" : "d3cc6575",
				"c" : "motion",
				"tid" : "433mhz-motion",
				"lbl" : "Motion@Main Door",
				"grp" : "home",
				"prot" : "433MHZ",
				"sub" : [
					{ "n" : "Sensor", "lbl" : "s", "code" : "5592405" }
				]
			}

			{
				"_id" : ObjectId("5410894f11b9eeb306b151fa"),
				"uid" : "b7d51d00",
				"c" : "button",
				"tid" : "433mhz-4buttons",
				"lbl" : "Test Button",
				"grp" : "home",
				"prot" : "433MHZ",
				"sub" : [
					{ "n" : "button_a", "lbl" : "Button A", "code" : "5592512" },
					{ "n" : "button_b", "lbl" : "Button B", "code" : "5592368" },
					{ "n" : "button_c", "lbl" : "Button C", "code" : "5592332" },
					{ "n" : "button_d", "lbl" : "Button D", "code" : "5592323" }
				]
			}
		*/

		fmt.Println(content)
		fmt.Println("POST!!")
	}

}

func (app *WebApplicationDashboard) handleAbout(w http.ResponseWriter, req *http.Request) {
	w.Write(templateOutput("about", nil))
}

func (app *WebApplicationDashboard) handleThings(w http.ResponseWriter, req *http.Request) {
	w.Write(templateOutput("things", nil))
}

func (app *WebApplicationDashboard) handleSysInfo(w http.ResponseWriter, req *http.Request) {
	w.Write(templateOutput("sysinfo", nil))
}

func (app *WebApplicationDashboard) handleManageRules(w http.ResponseWriter, req *http.Request) {
	model := new(rulesEditorModel)
	model.Rules = app.rulesService.GetRules()

	w.Write(templateOutput("rules", model))
}

func (app *WebApplicationDashboard) handleEditRules(w http.ResponseWriter, req *http.Request) {
	w.Write(templateOutput("rule_edit", nil))
}

func (app *WebApplicationDashboard) handleThingAdd(w http.ResponseWriter, req *http.Request) {
	var model = new(webModelThingAdd)
	model.Things = app.thingService.GetThingTypes()

	w.Write(templateOutput("thing_add", model))
}

func (app *WebApplicationDashboard) handleResourceIcon(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	dt := app.thingService.GetThingType(vars["thingType"])

	w.Header().Set("Content-Type", "image/png")
	path := dt.Path + "/icon_128x.png"

	b, _ := ioutil.ReadFile(path)
	w.Write(b)
}

func (app *WebApplicationDashboard) handleWidgetView(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	dev, ok := app.thingService.GetThing(vars["thingId"])
	if ok {
		dt := dev.Descriptor
		path := dt.Path + "/view.html"

		fileContent, _ := ioutil.ReadFile(path)
		stringContent := string(fileContent)
		tmpl, _ := template.New("widgetview_" + dt.Name).Parse(stringContent)

		buf := bytes.NewBufferString("")
		err := tmpl.Execute(buf, dev)
		if err != nil {
			log.Fatalf("execution failed: %s", err)
		}

		model := new(webModelWidgetView)
		model.Content = template.HTML(buf.String())
		model.Thing = dev

		w.Write(templateOutput("thing_view", model))
	}
}

func (app *WebApplicationDashboard) handleWidgetConfigure(w http.ResponseWriter, req *http.Request) {

}

func (app *WebApplicationDashboard) handleWidgetUpdateConfiguration(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	vars := mux.Vars(req)
	dev, ok := app.thingService.GetThing(vars["thingId"])

	for key, value := range req.PostForm {
		val := value[0]

		if strings.HasPrefix(key, "thing.") {
			switch key {
			case "thing.name":
				dev.Name = val

			case "thing.description":
				dev.Description = val
			}
		} else if strings.HasPrefix(key, "attrib.") {
			key = strings.Replace(key, "attrib.", "", -1)
			dev.SaveAttribute(key, val)
		} else {
			log.Println("Unknown form value " + key)
		}
	}

	app.thingService.SaveThing(dev)

	if ok {
		w.Write(templateOutput("thing_config", dev))
	}
}

func (app *WebApplicationDashboard) handleSimulationService(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	protocol := vars["protocol"]
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	content := string(body)
	handler := app.container.ProtocolHandler(protocol)
	handler.Handle(content)
}

func (app *WebApplicationDashboard) handleEventsView(w http.ResponseWriter, req *http.Request) {
	model := new(webModelEvents)
	model.Events = app.dataSource.GetThingEvents(10)
	w.Write(templateOutput("events", model))
}

func (app *WebApplicationDashboard) handleSettingsView(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		config := app.environment.GetConfig()
		w.Write(templateOutput("settings", config))
	} else if req.Method == "POST" {
		fmt.Println(req.Body)
	} else {
		fmt.Println("Unknown Method")
	}
}

// WEB MODELS
type rulesEditorModel struct {
	Rules map[string]api.Rule
}

type webModelWidgetView struct {
	Content template.HTML
	Thing  api.Thing
}

type webModelDashboard struct {
	RAM_used     string
	Events_count int
	Thing_count int
	Uptime       string
	Disk_Free    string
	Home_things []api.Thing
	Events       []api.Event
}

type webModelThingAdd struct {
	Things map[string]api.ThingType
}

type webModelEvents struct {
	Events []api.Event
}

type webModelThingAddNew struct {
	AddNewContent template.HTML
}
