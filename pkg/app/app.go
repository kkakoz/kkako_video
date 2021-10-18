package app

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"kkako_video/internal/pkg/handle"
	"log"
	"net"
	"net/http"
	"time"
)

type App struct {
	Name       string `json:"name"`
	Port       string `json:"port"`
	IP         string `json:"ip"`
	Ctx        context.Context
	Cancel     context.CancelFunc
	Logger     *zap.Logger
	GrpcServer *grpc.Server
	HttpServer http.Handler
	servers    []Server
}

func NewApp(ctx context.Context, cancel context.CancelFunc, opts ...Option) (*App, error) {
	app := &App{
		Ctx:    ctx,
		Cancel: cancel,
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
	for _, server := range a.servers {
		go func() {
			err := server.Run()
			if err != nil {
				a.Logger.Fatal("server err", zap.Error(err))
				a.Cancel()
			}
		}()
	}
	handler := handle.ServerHandlerFunc(a.GrpcServer, a.HttpServer)
	server := &http.Server{
		Handler: handler,
	}
	// 延时关闭
	go func() {
		<-a.Ctx.Done()
		ctx, _ := context.WithTimeout(context.TODO(), 5 * time.Second)
		err := server.Shutdown(ctx)
		if err != nil {
			log.Fatalln("shutdown err:", err)
		}
		for _, s := range a.servers {
			s.Stop()
		}
	}()
	err = server.Serve(listen)
	if err != nil {
		return errors.Wrap(err, "app start err")
	}
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
