package consequences

import "fmt"

type CallRestService struct {
}

func (s CallRestService) Execute(config map[string]interface{}) {
	fmt.Println("Consequence: Call REST Service")
}
