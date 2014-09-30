// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thingfu/hub/api"
	"github.com/thingfu/hub/container"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
)

type WebApplication struct {
	port uint16

	rulesService api.RulesManager
	thingManager api.ThingManager
	dataSource   api.DataSource
	environment  api.Environment
	factory      api.Factory
	container    api.Container
}

func NewWebApplication(port uint16) {
	w := new(WebApplication)
	w.port = port

	c := container.Instance()
	w.container = c
	w.rulesService = c.RulesManager()
	w.dataSource = c.DataSource()
	w.thingManager = c.ThingManager()
	w.environment = c.Env()
	w.factory = c.Factory()

	r := w.initializeRoutes()
	portStr := fmt.Sprintf("%d", w.port)
	// portStr := strconv.Itoa(w.port)

	http.Handle("/", r)
	log.Println("[INFO] Start Node WebServer @ " + portStr)
	err := http.ListenAndServe(":"+portStr, nil)
	if err != nil {
		log.Print("Error starting GoHome: ", err)
	}
}

func (w WebApplication) initializeRoutes() *mux.Router {
	r := mux.NewRouter()

	// PAGES
	reg(r, "/dashboard", w.showDashboard)
	reg(r, "/rules", w.showRules)
	reg(r, "/rule/{id}", w.showRule)
	reg(r, "/settings", w.showSettings)
	reg(r, "/events", w.showEvents)
	reg(r, "/thing/{id}/view", w.showThingView)
	reg(r, "/thing/{id}/configure", w.showThingConfigure)
	reg(r, "/thing/{type}/resource/img/{img}", w.showImage)
	reg(r, "/thing/add/{type}", w.showAddThing)
	reg(r, "/things/add", w.showThingsToAdd)
	reg(r, "/sysinfo", w.showSysInfo)
	reg(r, "/about", w.showAbout)

	// UI API
	reg(r, "/api/ui/dashboard", w.getDashboardState)
	reg(r, "/api/ui/thing/{id}/view", w.viewThing)

	// API
	reg(r, "/api/things", w.getThings)
	reg(r, "/api/things/types", w.getThingTypes)
	reg(r, "/api/thing/{id}", w.getThing)
	reg(r, "/api/thing/{id}/event/{svc}", w.triggerEventForThing).Methods("POST")
	reg(r, "/api/thing/{id}/events/{limit}", w.getEventsForThing).Methods("POST")
	reg(r, "/api/thing/{id}/op", w.invokeThingOperation).Methods("POST")
	reg(r, "/api/settings", w.getSettings)

	// Events
	reg(r, "/api/events/{limit}", w.getEvents)
	reg(r, "/api/event", w.addEvent).Methods("PUT")

	// Rules
	reg(r, "/api/rules", w.getRules)
	reg(r, "/api/rule/{id}", w.getRule).Methods("GET")
	reg(r, "/api/rule/{id}", w.saveRule).Methods("POST")
	reg(r, "/api/rule/{id}", w.deleteRule).Methods("DELETE")
	reg(r, "/api/rule/{id}", w.addRule).Methods("PUT")

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./www/static/"))))

	return r
}

// GET http://localhost:8181/api/ui/dashboard
func (app *WebApplication) getDashboardState(w http.ResponseWriter, req *http.Request) {
	model := make(map[string]interface{})

	// RAM Used
	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)
	ramUsed := int(((float64(memStats.Sys) / 1024 / 1024) * 100) / 100)
	model["RAMUsed"] = fmt.Sprintf("%d MB", ramUsed)
	model["EventsProcessed"] = app.dataSource.GetEventsCount()
	model["Uptime"] = app.environment.GetUptime()

	things := app.thingManager.GetThings()

	thing_models := make([]*api.Thing, len(things))

	for i := 0; i < len(things); i++ {
		dev := &things[i]
		content := renderStringContent(dev.Descriptor.Path+"/widget.html", dev)
		dev.Content = content
		thing_models = append(thing_models, dev)
	}

	model["ThingCount"] = len(things)
	model["Things"] = things

	writeJsonModel(w, model)
}

// GET /api/events/{limit}
func (app *WebApplication) getEvents(w http.ResponseWriter, req *http.Request) {
	model := make(map[string]interface{})
	vars := mux.Vars(req)
	limit, _ := strconv.Atoi(vars["limit"])

	events := app.dataSource.GetEvents(limit)
	model["total"] = len(events)
	model["events"] = events

	writeJsonModel(w, model)
}

// PUT /api/event
func (app *WebApplication) addEvent(w http.ResponseWriter, req *http.Request) {
	model := make(map[string]interface{})

	writeJsonModel(w, model)
}

// GET http://localhost:8181/api/ui/thing/{id}/view
func (app *WebApplication) viewThing(w http.ResponseWriter, req *http.Request) {
	model := make(map[string]interface{})
	vars := mux.Vars(req)
	dev, ok := app.thingManager.GetThing(vars["thingId"])
	if ok {
		dev.Content = renderStringContent(dev.Descriptor.Path+"/view.html", dev)
		model["Thing"] = dev
	}

	writeJsonModel(w, model)
}

// GET http://localhost:8181/api/rules
func (app *WebApplication) getRules(w http.ResponseWriter, req *http.Request) {
	writeJsonModel(w, app.rulesService.GetRules())
}

// PUT http://localhost:8181/api/rule/{id}
func (app *WebApplication) addRule(w http.ResponseWriter, req *http.Request) {
	model := make(map[string]interface{})

	writeJsonModel(w, model)
}

// POST http://localhost:8181/api/rule/{id}
func (app *WebApplication) saveRule(w http.ResponseWriter, req *http.Request) {
	model := make(map[string]interface{})

	writeJsonModel(w, model)
}

// DELETE http://localhost:8181/api/rule/{id}
func (app *WebApplication) deleteRule(w http.ResponseWriter, req *http.Request) {
	model := make(map[string]interface{})

	writeJsonModel(w, model)
}

// GET http://localhost:8181/api/rule/{id}
func (app *WebApplication) getRule(w http.ResponseWriter, req *http.Request) {
	model := make(map[string]interface{})
	vars := mux.Vars(req)
	ruleId := vars["ruleId"]

	rule := app.rulesService.GetRule(ruleId)
	model["id"] = rule.Id

	fileContent, _ := ioutil.ReadFile(rule.Path)
	stringContent := string(fileContent)
	model["content"] = stringContent
	model["name"] = rule.Name

	writeJsonModel(w, model)
}

// POST http://localhost:8181/api/thing/{id}/op/{op}
func (app *WebApplication) invokeThingOperation(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	op := vars["op"]

	thing, ok := app.thingManager.GetThing(id)
	if !ok {
		http.Error(w, "Not Found", 404)
	}

	params := make(map[string]interface{})
	app.thingManager.Actuate(&thing, op, params)

	model := make(map[string]interface{})

	writeJsonModel(w, model)
}

// GET http://localhost:8181/api/things
func (app *WebApplication) getThings(w http.ResponseWriter, req *http.Request) {
	writeJsonModel(w, app.thingManager.GetThings())
}

// GET http://localhost:8181/api/things/types
func (app *WebApplication) getThingTypes(w http.ResponseWriter, req *http.Request) {
	writeJsonModel(w, app.thingManager.GetThingTypes())
}

// GET http://localhost:8181/api/thing/{id}
func (app *WebApplication) getThing(w http.ResponseWriter, req *http.Request) {
	model := make(map[string]interface{})
	vars := mux.Vars(req)
	id := vars["id"]

	dev, ok := app.thingManager.GetThing(id)
	if ok {
		model["thing"] = dev
	}

	writeJsonModel(w, model)
}

// GET http://localhost:8181/api/thing/{id}/events/{limit}
func (app *WebApplication) getEventsForThing(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	limit, _ := strconv.Atoi(vars["limit"])

	events := app.dataSource.GetThingEvents(limit, id)

	writeJsonModel(w, events)
}

// POST http://localhost:8181/api/thing/{id}/event/{svc}
func (app *WebApplication) triggerEventForThing(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	svc := vars["svc"]
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	// content := string(body)
	thing, _ := app.thingManager.GetThing(id)
	service := thing.GetService(svc)
	var state map[string] interface {}
	json.Unmarshal(body, &state)

	app.thingManager.Handle(&thing, service, state)
}

// GET http://localhost:8181/api/settings
func (app *WebApplication) getSettings(w http.ResponseWriter, req *http.Request) {
	writeJsonModel(w, app.environment.GetConfig())
}

// VIEWS
// GET http://localhost:8181/dashboard
func (app *WebApplication) showDashboard(w http.ResponseWriter, req *http.Request) {
	w.Write(templateOutput("dashboard", nil))
}

// GET http://localhost:8181/rules
func (app *WebApplication) showRules(w http.ResponseWriter, req *http.Request) {
	w.Write(templateOutput("rules", nil))
}

// GET http://localhost:8181/rule/{id}
func (app *WebApplication) showRule(w http.ResponseWriter, req *http.Request) {
	w.Write(templateOutput("rule_edit", nil))
}

// GET http://localhost:8181/settings
func (app *WebApplication) showSettings(w http.ResponseWriter, req *http.Request) {
	w.Write(templateOutput("settings", nil))
}

// GET http://localhost:8181/events
func (app *WebApplication) showEvents(w http.ResponseWriter, req *http.Request) {
	w.Write(templateOutput("events", nil))
}

// GET http://localhost:8181/thing/{id}/view
func (app *WebApplication) showThingView(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	dev, ok := app.thingManager.GetThing(vars["id"])
	if ok {
		model := new(webModelWidgetView)
		model.Content = template.HTML(renderStringContent(dev.Descriptor.Path+"/view.html", dev))
		model.Thing = dev

		w.Write(templateOutput("thing_view", model))
	}
}

// GET http://localhost:8181/thing/{id}/configure
func (app *WebApplication) showThingConfigure(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	dev, ok := app.thingManager.GetThing(vars["id"])
	if ok {
		w.Write(templateOutput("thing_config", dev))
	}
}

// GET http://localhost:8181/thing/add/{type}
func (app *WebApplication) showAddThing(w http.ResponseWriter, req *http.Request) {
	w.Write(templateOutput("thing_addnew", nil))
}

// GET http://localhost:8181/things/add
func (app *WebApplication) showThingsToAdd(w http.ResponseWriter, req *http.Request) {
	w.Write(templateOutput("thing_add", nil))
}

// GET http://localhost:8181/sysinfo
func (app *WebApplication) showSysInfo(w http.ResponseWriter, req *http.Request) {
	w.Write(templateOutput("sysinfo", nil))
}

// GET http://localhost:8181/about
func (app *WebApplication) showAbout(w http.ResponseWriter, req *http.Request) {
	w.Write(templateOutput("about", nil))
}

// GET http://localhost:8181/thing/{type}/resource/img/{img}}
func (app *WebApplication) showImage(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	thingType := vars["type"]
	img := vars["img"]

	dt := app.thingManager.GetThingType(thingType)

	w.Header().Set("Content-Type", "image")
	path := dt.Path + "/" + img

	b, _ := ioutil.ReadFile(path)
	w.Write(b)
}

func reg(r *mux.Router, url string, fn func(http.ResponseWriter, *http.Request)) *mux.Router {
	r.HandleFunc(url, fn)

	return r
}

// PACKAGE FUNCTIONS
func writeJsonModel(w http.ResponseWriter, model interface{}) {
	out, err := json.Marshal(model)
	if err != nil {
		log.Println(err)
	}
	w.Write(out)
}

func compileTemplate(name string) *template.Template {
	t := template.New(name)
	t = template.Must(t.Funcs(funcMap).ParseGlob("www/views/layouts/*.html"))
	t.Delims("#{", "}#")

	return template.Must(t.ParseFiles("www/views/" + name + ".html"))
}

func templateOutput(name string, model interface{}) []byte {
	tpl := compileTemplate(name)

	var buf bytes.Buffer
	err := tpl.Execute(&buf, model)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
	return buf.Bytes()
}

func renderStringContent(path string, model interface{}) string {
	fileContent, _ := ioutil.ReadFile(path)
	stringContent := string(fileContent)

	t, err := template.New("__tpl_"+path).Delims("#{", "}#").Parse(stringContent)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}

	buf := bytes.NewBufferString("")
	err = t.Execute(buf, model)

	htmlContent := buf.String()
	return htmlContent
}
