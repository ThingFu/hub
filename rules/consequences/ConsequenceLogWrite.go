package consequences

import (
	"github.com/thingfu/hub/api"
	"log"
)

type LogWrite struct {
}

func (s LogWrite) Execute(config map[string]interface{}, container api.Container) {
	log.Println("%v", config)
}
