// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/go-home/hub/container"
	"log"
)

type WebApplication struct {
	port int
}

func NewWebApplication(port int) {
	w := new(WebApplication)
	w.port = port

	r := w.initializeRoutes()
	portStr := strconv.Itoa(w.port)

	http.Handle("/", r)
	log.Println("[INFO] Start Node WebServer @ " + portStr)
	err := http.ListenAndServe(":" + portStr, nil)
	if err != nil {
		log.Print("Error starting GoHome: ", err )
	}
}

func (w WebApplication) initializeRoutes() *mux.Router {
	r := mux.NewRouter()

	c := container.Instance()

	dashboardSetup := new(WebApplicationDashboard)
	dashboardSetup.rulesService = c.RulesService()
	dashboardSetup.dataSource = c.DataSource()
	dashboardSetup.deviceService = c.DeviceService()
	dashboardSetup.environment = c.Env()
	dashboardSetup.factory = c.Factory()
	dashboardSetup.container = c;
	dashboardSetup.Setup(r)

	apiSetup := new(WebApplicationApi)
	apiSetup.Setup(r)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./www/static/"))))

	return r
}

