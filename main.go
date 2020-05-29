package main

import (
	"github.com/cfi2017/wings-api/pkg"
)

func InitPlugin() pkg.Plugin {
	return &EventsPlugin{}
}

type EventsPlugin struct {
}

func (p *EventsPlugin) Load(api pkg.Api) error {
	api.Logger().Info("plugin loading")
	api.RegisterHandler("ping", func(api pkg.Api) {
		api.CallHandler("pong")
	})
	return nil
}
