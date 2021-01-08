package httplistener

import (
	"net/http"
	"net/rpc"
)

type Listener struct {
	addr   string
	server *rpc.Server
}

func New(addr string) *Listener {
	return &Listener{
		addr:   addr,
		server: rpc.NewServer(),
	}
}

func (l *Listener) RegisterReceiver(name string, receiver interface{}) error {
	return l.server.RegisterName(name, receiver)
}

func (l *Listener) Handle() error {
	return http.ListenAndServe(l.addr, nil)
}
