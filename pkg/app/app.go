package app

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type App struct {
	Name       string
	Port       string
	IP         string
	logger     *zap.Logger
	grpcServer *grpc.Server
}

func NewApp(opts ...Option) (*App, error) {
	app := &App{}
	for _, opt := range opts {
		err := opt.f(app)
		if err != nil {
			return app, err
		}
	}
	//app := &App{}
	//err := viper.UnmarshalKey("app", app)
	//if err != nil {
	//	return nil, errors.Wrap(err, "viper unmarshal失败")
	//}
	//app.grpcServer = grpcServer
	return app, nil
}

func (a *App) Start() error {
	listen, err := net.Listen("tcp", ":"+a.Port)
	if err != nil {
		return err
	}
	reflection.Register(a.grpcServer)
	err = a.grpcServer.Serve(listen)
	if err != nil {
		return errors.Wrap(err, "app start err")
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

func GrpcServer(server *grpc.Server) Option {
	return Option{
		f: func(app *App) error {
			app.grpcServer = server
			return nil
		},
	}
}
