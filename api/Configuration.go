// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

// Configuration as loaded from <home>/hub-config.json

const (
	CONFIG_FILE = "hub-config.json"
)

type MailConfig struct {
	User string
	Pass string
	Host string
	Port int
}

type Configuration struct {
	Name       string
	NodeID     string
	Db         string
	MaxProcs   int
	ServerPort int
	Protocols  map[string]ProtocolConfiguration
	Mail       MailConfig
}

func (c Configuration) GetProtocolConfiguration(protocol string) ProtocolConfiguration {
	return c.Protocols[protocol]
}
