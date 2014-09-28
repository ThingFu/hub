// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"flag"
	"github.com/thingfu/hub/api"
	"github.com/thingfu/hub/container"
	"github.com/thingfu/hub/setup"
	"io/ioutil"
	"log"
	_ "net/http/pprof"
)

func main() {
	// Read -home parameter and, if specified, use that as home directory
	// for reading thing definitions, rules and config
	homeFlag := flag.String("home", "./home", "Home directory of node")
	flag.Parse()
	home := *homeFlag

	log.Println("[INFO] Starting GoHome Node..")
	log.Println("[INFO] Home: " + home)
	log.Println("[INFO] Learning Kung-Fu..")

	// Read config file and parse
	content, err := ioutil.ReadFile(string(home + "/" + api.CONFIG_FILE))
	if err != nil {
		log.Fatal(err)
	}

	var config api.Configuration
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal(err)
	}

	// Do any validation on config, set sensible defaults or die if any
	// faults or missing config found
	validateConfig(&config)

	// Initialize IoC container
	_, env := container.Initialize(home, config)

	// Do setup process
	setup.Setup(env)
}

func validateConfig(cfg *api.Configuration) {
	if cfg.Db == "" {
		log.Fatal("Database Parameter Missing")
	}

	if cfg.ServerPort == 0 {
		log.Println("Server port not defined. Defaulting to 8181")
		cfg.ServerPort = 8181
	}

	if len(cfg.Protocols) == 0 {
		log.Println("No protocols defined")
	}
}
