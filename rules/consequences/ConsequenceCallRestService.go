package consequences

import "log"

type CallRestService struct {
}

func (s CallRestService) Execute(config map[string]interface{}) {
	log.Println("Consequence: Call REST Service")
}
