package channels

import (
	"github.com/thingfu/hub/api"
)

type DefaultCommChannelManager struct {
	container	api.Container
	chanMap		map[string] api.CommunicationChannel
	factory		api.Factory
}

func NewCommChannelManager() api.CommChannelManager {
	c := new (DefaultCommChannelManager)
	c.chanMap = make(map[string] api.CommunicationChannel)

	return c
}

func (d *DefaultCommChannelManager) InitChannels(chs []api.ChannelConfiguration) {
	// Register Channels
	for _, channel := range chs {
		c := d.factory.CreateChannelHandler(channel)
		c.SetChannelConfiguration(channel)

		for _, protocol := range channel.Protocols {
			p := d.factory.CreateProtocolHandler(protocol)

			c.AddProtocol(p)
			p.SetChannel(c)

			d.Register(protocol, c)
		}

		if c.IsEnabled() {
			go c.Start()
		}
	}
}

func (d *DefaultCommChannelManager) Register(protocol string, channel api.CommunicationChannel) {
	d.chanMap[protocol] = channel
}

func (d *DefaultCommChannelManager) GetChannelForProtocol(p string) api.CommunicationChannel {
	return d.chanMap[p]
}

func (d *DefaultCommChannelManager) SetFactory(o api.Factory) {
	d.factory = o
}

func (d *DefaultCommChannelManager) SetContainer(c api.Container) {
	d.container = c
}

func (s *DefaultCommChannelManager) ValidateWiring() {

}
