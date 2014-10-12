package consequences

import "log"

type ExecuteShell struct {
}

func (s ExecuteShell) Execute(config map[string]interface{}) {
	log.Println("Consequence: Execute Shell")
}
