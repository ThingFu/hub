// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"net/smtp"
	"strings"
	"text/template"
	"time"
	"errors"
)

func RandomString(str_size int) string {
	alphanum := "012345678S9ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func TimeWithinThreshold(last time.Time, threshold int, defaultThreshold int) bool {
	buf := int64(threshold * 1000000)
	if buf == 0 {
		buf = int64(defaultThreshold * 1000000)
	}

	now := time.Now()
	ns := now.Sub(last).Nanoseconds()
	if ns < buf {
		return false
	}
	return true
}

func SendEmail(host string, port uint16, userName string, password string, to []string, subject string, message string) (err error) {
	parameters := struct {
		From    string
		To      string
		Subject string
		Message string
	}{
		userName,
		strings.Join([]string(to), ","),
		subject,
		message,
	}

	buffer := new(bytes.Buffer)

	tpl := template.Must(template.New("emailTemplate").Parse(emailScript()))
	tpl.Execute(buffer, &parameters)

	auth := smtp.PlainAuth("", userName, password, host)

	err = smtp.SendMail(
		fmt.Sprintf("%s:%d", host, port),
		auth,
		userName,
		to,
		buffer.Bytes())

	return err
}

func emailScript() (script string) {
	return `From: {{.From}}
			To: {{.To}}
			Subject: {{.Subject}}
			MIME-version: 1.0
			Content-Type: text/html; charset="UTF-8"

			{{.Message}}`
}

func Now() *GoTime {
	t := new(GoTime)
	t.SetTime(time.Now())

	return t
}

func NewGoTime(o time.Time) *GoTime {
	t := new(GoTime)
	t.SetTime(o)

	return t
}

func ParseThingFuSerialData(data string) (p map[string]interface {}, e error) {
	parts := strings.Split(data, "|")

	if len(parts) < 4 {
		e = errors.New("Malformed ThingFu Serial Data: Invalid parameter length :" +  data)
		return
	}

	if parts[0] != "TF" {
		e = errors.New("Malformed ThingFu Serial Data: Does not start with TF :" +  data)
		return
	}

	/*
	if parts[3] != "E0" {
		e = errors.New("Malformed ThingFu Serial Data: Does not end with E0.")
		return
	}
	*/
	p = make(map[string]interface {})
	p["Protocol"] = parts[1]
	p["Data"] = parts[2]

	return

	// TF|1|13981013|E0
}
