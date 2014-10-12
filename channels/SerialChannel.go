package channels

import (
	"github.com/thingfu/hub/api"
	serial "github.com/tarm/goserial"
	"log"
	"strings"
)

type SerialChannel struct {
	api.BaseCommunicationChannel
}

func NewSerialChannel() (*SerialChannel) {
	s := new (SerialChannel)
	s.BaseCommunicationChannel = api.NewBaseCommunicationChannel()

	return s
}

func (ser SerialChannel) Start() error     {
	log.Println("[INFO] Starting Serial")

	cfg := ser.GetConfiguration()
	port := cfg.Port
	baud := cfg.Baud

	c := &serial.Config{ Name: port, Baud: int(baud) }
	s, err := serial.OpenPort(c)
	if err == nil {
		buf := ""

		for {
			serialContent := make([]byte, 256)
			_, err := s.Read(serialContent)
			if err != nil {
				log.Fatal(err)
			}

			c := string(serialContent)
			if strings.Index(c, "\n") != -1 {
				str := strings.Split(c, "\n")
				buf += str[0]
				buf = strings.Replace(buf, "\x00", "", -1)

				// Iterate all protocols and call Handle for Buf
				payload := api.NewReadRequest(buf)
				for _, prot := range ser.GetProtocols() {
					go prot.OnRead(payload)
				}

				buf = str[1]
			} else {
				buf += c
			}
		}
	} else {
		log.Print("[ERROR] Unable to start Serial -- ", err)
	}
	return nil
}

func (s SerialChannel) GetName() string  {
	return "Serial"
}

func (s SerialChannel) GetLabel() string {
	return "Serial Channel"
}

func (p *SerialChannel) ValidateWiring() {

}

