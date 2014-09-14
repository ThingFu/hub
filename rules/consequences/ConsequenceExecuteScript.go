package consequences

import "fmt"

type ExecuteScript struct {
}

func (s ExecuteScript) Execute(config map[string]interface{}) {
	fmt.Println("Consequence: Execute Script")
}
