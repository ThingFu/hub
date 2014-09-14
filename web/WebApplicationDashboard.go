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
	"runtime"
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

/*
var templates = template.Must(
	templates.ParseFiles(
		"www/views/layouts/header.html",
		"www/views/layouts/footer.html",
		"www/views/dashboard.html",
		"www/views/rules.html",
		"www/views/settings.html",
		"www/views/events.html",
		"www/views/widget_view.html",
		"www/views/add_device.html",
		"www/views/devices.html",
		"www/views/sysinfo.html",
		"www/views/about.html",
		"www/views/widget_config.html").Funcs(funcMap))
*/

func compileTemplate(name string) *template.Template {
	t := template.New(name)
	t = template.Must(t.Funcs(funcMap).ParseGlob("www/views/layouts/*.html"))

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

type WebApplicationDashboard struct {
	rulesService  api.RulesService
	deviceService api.DeviceService
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
	r.HandleFunc("/widget/{deviceId}/configure", app.handleWidgetConfigure).Methods("GET")
	r.HandleFunc("/widget/{deviceId}/configure", app.handleWidgetUpdateConfiguration).Methods("POST")
	r.HandleFunc("/widget/{deviceId}/view", app.handleWidgetView)
	r.HandleFunc("/device/{deviceType}/resource/icon/128x", app.handleResourceIcon)
	r.HandleFunc("/dashboard", app.handleDashboard)
	r.HandleFunc("/device/add", app.handleDeviceAdd)
	r.HandleFunc("/devices", app.handleDevices)
	r.HandleFunc("/sysinfo", app.handleSysInfo)
	r.HandleFunc("/about", app.handleAbout)
}

func (app *WebApplicationDashboard) handleAbout(w http.ResponseWriter, req *http.Request) {
	w.Write(templateOutput("about", nil))
}

func (app *WebApplicationDashboard) handleDevices(w http.ResponseWriter, req *http.Request) {
	w.Write(templateOutput("devices", nil))
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
	w.Write(templateOutput("rules", nil))
}

func (app *WebApplicationDashboard) handleDeviceAdd(w http.ResponseWriter, req *http.Request) {
	var model = new(webModelDeviceAdd)
	model.Devices = app.deviceService.GetDeviceTypes()

	w.Write(templateOutput("add_device", model))
}

func (app *WebApplicationDashboard) handleDashboard(w http.ResponseWriter, req *http.Request) {
	// TODO Pre-render templates
	var dp = new(webModelDashboard)

	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)
	ramUsed := int(((float64(memStats.Sys) / 1024 / 1024) * 100) / 100)
	dp.RAM_used = fmt.Sprintf("%d", ramUsed)

	dp.Events_count = len(app.dataSource.GetDeviceEvents(0))

	dp.Disk_Free = ""
	dp.Uptime = app.environment.GetUptime()
	devices := app.deviceService.GetDevices()
	dp.Home_devices = devices
	dp.Device_count = len(devices)

	// return to client via t.Execute
	w.Write(templateOutput("dashboard", dp))
}

func (app *WebApplicationDashboard) handleResourceIcon(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	dt := app.deviceService.GetDeviceType(vars["deviceType"])

	w.Header().Set("Content-Type", "image/png")
	path := dt.Path + "/icon_128x.png"

	b, _ := ioutil.ReadFile(path)
	w.Write(b)
}

func (app *WebApplicationDashboard) handleWidgetView(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	dev, ok := app.deviceService.GetDevice(vars["deviceId"])
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
		model.Device = dev

		w.Write(templateOutput("widget_view", model))
	}
}

func (app *WebApplicationDashboard) handleWidgetConfigure(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	dev, ok := app.deviceService.GetDevice(vars["deviceId"])
	if ok {
		w.Write(templateOutput("widget_config", dev))
	}
}

func (app *WebApplicationDashboard) handleWidgetUpdateConfiguration(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	vars := mux.Vars(req)
	dev, ok := app.deviceService.GetDevice(vars["deviceId"])

	for key, value := range req.PostForm {
		val := value[0]

		if strings.HasPrefix(key, "device.") {
			switch key {
			case "device.name":
				dev.Name = val

			case "device.description":
				dev.Description = val
			}
		} else if strings.HasPrefix(key, "attrib.") {
			key = strings.Replace(key, "attrib.", "", -1)
			dev.SaveAttribute(key, val)
		} else {
			log.Println("Unknown form value " + key)
		}
	}

	app.dataSource.SaveDevice(dev)
	app.deviceService.SaveDevice(dev)

	if ok {
		w.Write(templateOutput("widget_config", dev))
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

	/*
		dev, _ := app.deviceService.GetDevice("7820592")
		// facts.Device = &dev

		go app.deviceService.Handle(dev)

		// go rulesManager.Trigger(TRIGGER_DEVICE, facts)
	*/
}

func (app *WebApplicationDashboard) handleEventsView(w http.ResponseWriter, req *http.Request) {
	model := new(webModelEvents)
	model.Events = app.dataSource.GetDeviceEvents(10)
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
	Device  api.Device
}

type webModelDashboard struct {
	RAM_used     string
	Events_count int
	Device_count int
	Uptime       string
	Disk_Free    string
	Home_devices []api.Device
	Events       []api.Event
}

type webModelDeviceAdd struct {
	Devices map[string]api.DeviceType
}

type webModelEvents struct {
	Events []api.Event
}
