package api

type CommChannelManager interface {
	GetChannelForProtocol(string) CommunicationChannel
	InitChannels([]ChannelConfiguration)
	Register(string, CommunicationChannel)

	ContainerAware
	SetFactory(Factory)
}


