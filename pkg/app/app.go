package app

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type App struct {
	Name       string `json:"name"`
	Port       string `json:"port"`
	IP         string `json:"ip"`
	Ctx        context.Context
	Logger     *zap.Logger
	GrpcServer *grpc.Server
	servers    []Server
}

func NewApp(ctx context.Context, opts ...Option) (*App, error) {
	app := &App{
		Ctx: ctx,
	}
	for _, opt := range opts {
		err := opt.f(app)
		if err != nil {
			return app, err
		}
	}
	return app, nil
}

func (a *App) Start() error {
	listen, err := net.Listen("tcp", ":"+a.Port)
	if err != nil {
		return err
	}
	reflection.Register(a.GrpcServer)
	err = a.GrpcServer.Serve(listen)
	if err != nil {
		return errors.Wrap(err, "app start err")
	}
	go func() {
		<-a.Ctx.Done()
		a.GrpcServer.Stop()
		for _, server := range a.servers {
			server.Stop()
		}
	}()
	for _, server := range a.servers {
		err = server.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

type Option struct {
	f func(app *App) error
}

func Name(name string) Option {
	return Option{
		f: func(app *App) error {
			app.Name = name
			return nil
		},
	}
}

func Port(port string) Option {
	return Option{
		f: func(app *App) error {
			app.Port = port
			return nil
		},
	}
}

func IP(ip string) Option {
	return Option{
		f: func(app *App) error {
			app.IP = ip
			return nil
		},
	}
}

func RegisterServer(server Server) Option {
	return Option{
		f: func(app *App) error {
			app.servers = append(app.servers, server)
			return nil
		},
	}
}

func GrpcServer(server *grpc.Server) Option {
	return Option{
		f: func(app *App) error {
			app.GrpcServer = server
			return nil
		},
	}
}
