// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package consequences

import (
	"fmt"
	"github.com/thingfu/hub/api"
	"github.com/thingfu/hub/utils"
	"log"
)

type SendMail struct {
}

func (s SendMail) Execute(config map[string]interface{}, container api.Container) {
	mailConfig := container.Env().GetConfig().Mail

	subject := config["subject"].(string)
	to := config["to"].(string)
	content := config["content"].(string)

	host := mailConfig.Host
	port := mailConfig.Port
	user := mailConfig.User
	pass := mailConfig.Pass

	err := utils.SendEmail(host, port, user, pass, []string{to}, subject, content)
	if err != nil {
		log.Print("Error sending email ", err)
	}

	fmt.Println("%v", config)
}
