package consequences

import (
	"fmt"
	"github.com/go-home/hub/api"
)

type LogWrite struct {
}

func (s LogWrite) Execute(config map[string]interface{}, container api.Container) {
	fmt.Println("Executing Log Write")

	fmt.Println("%v", config)
}
