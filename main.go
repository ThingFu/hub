// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"flag"
	"github.com/go-home/hub/api"
	"github.com/go-home/hub/container"
	"github.com/go-home/hub/setup"
	"io/ioutil"
	"log"
)

func main() {
	homeFlag := flag.String("home", "./home", "Home directory of node")
	flag.Parse()
	home := *homeFlag

	log.Println("[INFO] Starting GoHome Node..")
	log.Println("[INFO] Home: " + home)
	log.Println("[INFO] Learning Kung-Fu..")

	content, err := ioutil.ReadFile(string(home + "/node_config.json"))
	if err != nil {
		log.Fatal(err)
	}

	var config api.Configuration
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal(err)
	}
	validateConfig(config)
	_, env := container.Initialize(home, config)

	setup.Setup(env)
}

func validateConfig(cfg api.Configuration) {

}
