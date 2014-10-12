// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package setup

import (
	"encoding/json"
	"github.com/thingfu/hub/api"
	"github.com/thingfu/hub/container"
	"github.com/thingfu/hub/utils"
	"github.com/thingfu/hub/web"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Setup(env api.Environment) {
	// setupChan := make(chan bool, 6)

	// Load Thing Definitions
	go loadThingTypes(env)

	// Load Thing Instances
	loadThings(env)

	// Load Rule Defintions
	loadRules(env)

	// Register protocols. Has to be goroutines since these buggers
	// loops forever
	// go setupProtocols()

	go setupChannels()

	// Start scheduler for stuff like the peroidic rule invoker
	startScheduler()

	// Setup the WebApplication (UI and REST)
	setupWebApplication(env)

	// Signal that everything's completed and we can start
	// accepting/serving requests
	setupCompleted()
}

func startScheduler() {
	log.Println("[INFO] Whipping Hamsters into a Frenzy..")
	scheduleService := container.Instance().ScheduleService()
	scheduleService.Start()
}

func loadRules(env api.Environment) {
	log.Println("[INFO] Dividing by Zero")
	rulesService := container.Instance().RulesManager()
	filepath.Walk(env.GetHome()+"/rules", func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			var rule api.Rule
			err = json.Unmarshal(content, &rule)

			if err != nil {
				log.Printf("[ERROR] Parsing Rule %s\n", path)
				return nil
			}
			rule.Id = utils.RandomString(7)

			rule.Path = path[strings.LastIndex(path, "/")+1:]

			log.Printf("[INFO] Registering rule file: %s\n", rule.Name)

			rulesService.RegisterRule(rule)
		}
		return nil
	})
}

func setupCompleted() {
	log.Println("[INFO] Node Started")
}

// Loads thing definitions from a DataSource
// and registers them
//
func loadThings(env api.Environment) {
	log.Println("[INFO] Finding Waldo..")

	thingManager := container.Instance().ThingManager()
	thingManager.LoadThings()
}

// Scans for defintions under home's /things subdirectory
// and it walks through the path.
// if the descriptor.json file is found, it is assumed that the
// directory it resides in contains a thing defintion
//
func loadThingTypes(env api.Environment) {
	log.Println("[INFO] Pressing the Any Key...")
	thingManager := container.Instance().ThingManager()
	home := env.GetHome()
	root := home + "/things"

	filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, "descriptor.json") {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			var thingType api.ThingType
			err = json.Unmarshal(content, &thingType)

			if err != nil {
				log.Println("error: ", err)
			}
			thingType.Path = strings.Replace(path, "/descriptor.json", "", -1)
			thingManager.RegisterThingType(thingType)
		}
		return nil
	})
}

func setupWebApplication(env api.Environment) {
	web.NewWebApplication(env.GetConfig().ServerPort)
}

// Call each Protocol Handler's Start.
// Each protocol handler would need its own dedicated
// goroutine/thread to run in
//
func setupChannels() {
	log.Println("[INFO] Setting Up Channels")
	c := container.Instance()

	chs := c.Channels()
	for _, ch := range chs {
		if ch.IsEnabled() {
			ch.Start()
		}
	}
}
