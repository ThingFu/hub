// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

type WebApplicationApi struct {
}

func (w *WebApplicationApi) Setup(r *mux.Router) {

	r.HandleFunc("/api/status", handleApiGetStatus)
	r.HandleFunc("/api/devices", handleApiGetDevices)
	r.HandleFunc("/api/device/{deviceId}", handleApiGetDevice)
	r.HandleFunc("/api/device/{deviceId}", handleApiDeleteDevice).Methods("DELETE")
	r.HandleFunc("/api/device/{deviceId}", handleApiUpdateDevice).Methods("POST")
	r.HandleFunc("/api/devices", handleApiAddDevice).Methods("PUT")
	r.HandleFunc("/api/device/{deviceId}/state/{limit}", handleApiGetDeviceState)
	r.HandleFunc("/api/device/{deviceId}/state", handleApiAddDeviceState).Methods("PUT")
	r.HandleFunc("/api/events/{limit}", handleApiGetEvents)
	r.HandleFunc("/api/events/device/{deviceId}/{limit}", handleApiGetDeviceEvents)
	r.HandleFunc("/api/rules/{ruleId}", handleApiInvokeRule).Methods("POST")
	r.HandleFunc("/api/subscribe/{deviceId}", handleApiSubscribeDevice)

	r.HandleFunc("/api/device/{deviceId}/service/{name}", handleDeviceService).Methods("GET", "POST", "DELETE", "PUT")
}

func handleDeviceService(w http.ResponseWriter, req *http.Request)      {}
func handleApiGetStatus(w http.ResponseWriter, req *http.Request)       {}
func handleApiGetDevices(w http.ResponseWriter, req *http.Request)      {}
func handleApiGetDevice(w http.ResponseWriter, req *http.Request)       {}
func handleApiDeleteDevice(w http.ResponseWriter, req *http.Request)    {}
func handleApiUpdateDevice(w http.ResponseWriter, req *http.Request)    {}
func handleApiAddDevice(w http.ResponseWriter, req *http.Request)       {}
func handleApiGetDeviceState(w http.ResponseWriter, req *http.Request)  {}
func handleApiAddDeviceState(w http.ResponseWriter, req *http.Request)  {}
func handleApiGetEvents(w http.ResponseWriter, req *http.Request)       {}
func handleApiGetDeviceEvents(w http.ResponseWriter, req *http.Request) {}
func handleApiInvokeRule(w http.ResponseWriter, req *http.Request)      {}
func handleApiSubscribeDevice(w http.ResponseWriter, req *http.Request) {}
