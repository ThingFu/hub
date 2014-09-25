// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/go-home/hub/container"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"runtime"
	"encoding/json"
	"fmt"
	"github.com/go-home/hub/api"
	"io/ioutil"
	"html/template"
	"bytes"
)

type WebApplication struct {
	port int

	rulesService  api.RulesService
	deviceService api.DeviceService
	dataSource    api.DataSource
	environment   api.Environment
	factory       api.Factory
	container     api.Container
}

func NewWebApplication(port int) {
	w := new(WebApplication)
	w.port = port

	c := container.Instance()
	w.container = c
	w.rulesService = c.RulesService()
	w.dataSource = c.DataSource()
	w.deviceService = c.DeviceService()
	w.environment = c.Env()
	w.factory = c.Factory()

	r := w.initializeRoutes()
	portStr := strconv.Itoa(w.port)

	http.Handle("/", r)
	log.Println("[INFO] Start Node WebServer @ " + portStr)
	err := http.ListenAndServe(":"+portStr, nil)
	if err != nil {
		log.Print("Error starting GoHome: ", err)
	}
}

func (w WebApplication) initializeRoutes() *mux.Router {
	r := mux.NewRouter()

	// START: NEW
	r.HandleFunc("/api/ui/dashboard", w.apiUiDashboard)
	// END: NEW

	c := container.Instance()

	dashboardSetup := new(WebApplicationDashboard)
	dashboardSetup.rulesService = c.RulesService()
	dashboardSetup.dataSource = c.DataSource()
	dashboardSetup.deviceService = c.DeviceService()
	dashboardSetup.environment = c.Env()
	dashboardSetup.factory = c.Factory()
	dashboardSetup.container = c
	dashboardSetup.Setup(r)

	apiSetup := new(WebApplicationApi)
	apiSetup.Setup(r)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./www/static/"))))

	return r
}

func renderStringContent(path string, model interface {}) string {
	fileContent, _ := ioutil.ReadFile(path)
	stringContent := string(fileContent)

	tmpl, err := template.New("__tpl_" + path).Parse(stringContent)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}

	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, model)

	htmlContent := buf.String()
	return htmlContent
}

/*
func (d *Device) RenderWidget() template.HTML {
	path := d.Descriptor.Path

	fileContent, _ := ioutil.ReadFile(path + "/widget.html")
	stringContent := string(fileContent)

	tmpl, err := template.New("widget_" + d.Descriptor.Name).Parse(stringContent)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}

	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, d)

	htmlContent := buf.String()
	return template.HTML(htmlContent)
}
 */

// /api/ui/dashboard
func (app *WebApplication) apiUiDashboard(w http.ResponseWriter, req *http.Request) {
	model := make(map[string]interface {})

	// RAM Used
	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)
	ramUsed := int(((float64(memStats.Sys) / 1024 / 1024) * 100) / 100)
	model["RAMUsed"] = fmt.Sprintf("%d MB", ramUsed)
	model["EventsProcessed"] = len(app.dataSource.GetDeviceEvents(0))
	model["Uptime"] = app.environment.GetUptime()

	devices := app.deviceService.GetDevices()

	device_models := make([]*api.Device, len(devices))

	for i:=0; i < len(devices); i++ {
		dev := &devices[i]
		dev.Content = renderStringContent(dev.Descriptor.Path + "/widget.html", dev)
		device_models = append(device_models, dev)
	}

	model["DeviceCount"] = len(devices)
	model["Devices"] = devices

	out, _ := json.Marshal(model)
	w.Write(out)
}

// api/events/{limit}

// api/device/types



