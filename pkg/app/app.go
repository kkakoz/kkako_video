package app

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
)

type App struct {
	Name       string `json:"name"`
	Port       string `json:"port"`
	IP         string `json:"ip"`
	HttpPort   string `json:"http_port"`
	Ctx        context.Context
	Cancel     context.CancelFunc
	Logger     *zap.Logger
	GrpcServer *grpc.Server
	HttpServer http.Handler
	servers    []Server
}

func NewApp(ctx context.Context, cancel context.CancelFunc, opts ...Option) (*App, error) {
	app := &App{
		Ctx:      ctx,
		Cancel:   cancel,
		Port:     "9001",
		HttpPort: "10001",
	}
	mode := viper.GetViper().GetString("app.mode")
	if mode == "test" || mode == "pro" {
		name, b := os.LookupEnv("MY_POD_NAME")
		if !b {
			return nil, errors.New("get name err")
		}
		app.Name = name
		ip, b := os.LookupEnv("MY_POD_IP")
		if !b {
			return nil, errors.New("get ip err")
		}
		app.IP = ip
	}
	app.Port = viper.GetString("app.port")
	for _, opt := range opts {
		err := opt.f(app)
		if err != nil {
			return app, err
		}
	}
	return app, nil
}

func (a *App) Start() error {
	grpcConn, err := net.Listen("tcp", ":"+a.Port)
	if err != nil {
		return err
	}

	reflection.Register(a.GrpcServer)
	for _, server := range a.servers {
		go func() {
			err := server.Run()
			if err != nil {
				a.Logger.Fatal("server err", zap.Error(err))
				a.Cancel()
			}
		}()
	}
	go func() {
		<-a.Ctx.Done()
		a.GrpcServer.Stop()
		for _, s := range a.servers {
			s.Stop()
		}
	}()
	go a.ServerHttp()
	a.ServerGrpc(grpcConn)
	//err = a.GrpcServer.Serve(listen)
	//if err != nil {
	//	return errors.Wrap(err, "app start err")
	//}

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

func HttpServer(server http.Handler) Option {
	return Option{
		f: func(app *App) error {
			app.HttpServer = server
			return nil
		},
	}
}

func (a *App) ServerGrpc(conn net.Listener) {
	if a.GrpcServer == nil {
		a.Logger.Fatal("grpc server is nil")
	}
	err := a.GrpcServer.Serve(conn)
	if err != nil {
		a.Logger.Fatal("start grpc server err:", zap.Error(err))
		a.Cancel()
	}
}

func (a *App) ServerHttp() {
	if a.HttpServer == nil {
		return
	}
	conn, err := net.Listen("tcp", ":"+a.HttpPort)
	if err != nil {
		a.Logger.Fatal("app http server port is nil")
	}
	err = http.Serve(conn, a.HttpServer)
	if err != nil {
		a.Logger.Fatal("start http server err:", zap.Error(err))
		a.Cancel()
	}
}
