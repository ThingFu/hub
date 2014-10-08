package api

type Channel interface {
	Start() error
	IsEnabled(bool)
	GetName() string
	GetLabel() string
}
