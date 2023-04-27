package iprose

import (
	"github.com/go-chi/chi/v5"
	"github.com/vpnhouse/iprose-go/pkg/server"
)

type Instance struct {
	iprose *server.IPRoseServer
}

func New() (*Instance, error) {
	iprose, err := server.New(
		"iprose0",
		"10.123.76.1/24",
		"",
		[]string{"0.0.0.0/0"},
		128,
	)
	if err != nil {
		return nil, err
	}
	return &Instance{
		iprose: iprose,
	}, nil
}

func (instance *Instance) RegisterHandlers(r chi.Router) {
	for _, hndlr := range instance.iprose.Handlers() {
		r.HandleFunc(hndlr.Pattern, hndlr.Func)
	}
}

func (instance *Instance) Shutdown() error {
	instance.iprose.Shutdown()
	return nil
}

func (instance *Instance) Running() bool {
	return instance.iprose.Running()
}
