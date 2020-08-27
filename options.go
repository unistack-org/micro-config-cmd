package cmd

import (
	"context"

	"github.com/unistack-org/micro/v3/auth"
	"github.com/unistack-org/micro/v3/broker"
	"github.com/unistack-org/micro/v3/client"
	"github.com/unistack-org/micro/v3/config"
	"github.com/unistack-org/micro/v3/debug/profile"
	"github.com/unistack-org/micro/v3/debug/trace"
	"github.com/unistack-org/micro/v3/registry"
	"github.com/unistack-org/micro/v3/runtime"
	"github.com/unistack-org/micro/v3/server"
	"github.com/unistack-org/micro/v3/store"
	"github.com/unistack-org/micro/v3/transport"
)

type Options struct {
	// For the Command Line itself
	Name        string
	Description string
	Version     string

	// We need pointers to things so we can swap them out if needed.
	Broker    *broker.Broker
	Registry  *registry.Registry
	Transport *transport.Transport
	Config    *config.Config
	Client    *client.Client
	Server    *server.Server
	Runtime   *runtime.Runtime
	Store     *store.Store
	Tracer    *trace.Tracer
	Auth      *auth.Auth
	Profile   *profile.Profile
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

// Command line Name
func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

// Command line Description
func Description(d string) Option {
	return func(o *Options) {
		o.Description = d
	}
}

// Command line Version
func Version(v string) Option {
	return func(o *Options) {
		o.Version = v
	}
}

func Broker(b *broker.Broker) Option {
	return func(o *Options) {
		o.Broker = b
	}
}

func Config(c *config.Config) Option {
	return func(o *Options) {
		o.Config = c
	}
}

func Registry(r *registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}

func Runtime(r *runtime.Runtime) Option {
	return func(o *Options) {
		o.Runtime = r
	}
}

func Transport(t *transport.Transport) Option {
	return func(o *Options) {
		o.Transport = t
	}
}

func Client(c *client.Client) Option {
	return func(o *Options) {
		o.Client = c
	}
}

func Server(s *server.Server) Option {
	return func(o *Options) {
		o.Server = s
	}
}

func Store(s *store.Store) Option {
	return func(o *Options) {
		o.Store = s
	}
}

func Tracer(t *trace.Tracer) Option {
	return func(o *Options) {
		o.Tracer = t
	}
}

func Auth(a *auth.Auth) Option {
	return func(o *Options) {
		o.Auth = a
	}
}

func Profile(p *profile.Profile) Option {
	return func(o *Options) {
		o.Profile = p
	}
}
