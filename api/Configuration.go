// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

// Hub Configuration as loaded from <home>/hub-config.json
const (
	CONFIG_FILE = "hub-config.json"
)

// Configuration for Mail when sending out email (e.g. rule triggers)
type MailConfig struct {
	User string
	Pass string
	Host string
	Port uint16
}

type Configuration struct {
	Name       string
	NodeID     string
	Db         string
	MaxProcs   uint8
	ServerPort uint16
	Channels   []ChannelConfiguration
	Mail       MailConfig
}
