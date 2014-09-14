package consequences

import "fmt"

type ExecuteShell struct {
}

func (s ExecuteShell) Execute(config map[string]interface{}) {
	fmt.Println("Consequence: Execute Shell")
}
