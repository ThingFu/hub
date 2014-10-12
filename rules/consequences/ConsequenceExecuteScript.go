package consequences

import (
	"log"
)

type ExecuteScript struct {
}

func (s ExecuteScript) Execute(config map[string]interface{}) {
	log.Println("Consequence: Execute Script")
}
