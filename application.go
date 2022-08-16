package bone

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/facebookgo/inject"
	"google.golang.org/grpc"
)

// ApplicationOptions represents the options to create Bone Application
// instance.
type ApplicationOptions struct {
	// Basic part
	Debug bool

	// HTTP part
	HttpEnable bool
	HttpHost   string
	HttpPort   int

	// gRPC part
	GrpcEnable bool
	GrpcHost   string
	GrpcPort   int
}

// DefaultApplicationOptions returns the default options to create Bone
// Application instance. Enable HTTP Server and disable gRPC Server by default.
func DefaultApplicationOptions() *ApplicationOptions {
	return &ApplicationOptions{
		Debug: true,

		HttpEnable: true,
		HttpHost:   "127.0.0.1",
		HttpPort:   8080,

		GrpcEnable: false, // Disable gRPC by default
		GrpcHost:   "127.0.0.1",
		GrpcPort:   50051,
	}
}

// Application defines the struct of Bone Application.
type Application struct {
	o          *ApplicationOptions
	g          *inject.Graph
	httpRouter *Router
	grpcServer *grpc.Server
}

// NewApplication returns the instance of Bone Application.
func NewApplication(options *ApplicationOptions) *Application {
	app := &Application{
		o:          options,
		g:          &inject.Graph{},
		httpRouter: NewRouter(),
		grpcServer: grpc.NewServer(),
	}

	if app.o.Debug {
		log.Println("WARNING: Running in debug mode.")
	}

	// Inject core components
	app.Inject(&inject.Object{Name: "application.router", Value: app.httpRouter})
	app.Inject(&inject.Object{Name: "application.grpc", Value: app.grpcServer})

	return app
}

// Inject is a low-level API to inject *inject.Object to bone Application.
// Do NOT use it except must to.
func (app *Application) Inject(obj *inject.Object) error {
	return app.g.Provide(obj)
}

func (app *Application) register(component Component) {
	if app.o.Debug {
		log.Println("Registering Component: " + component.Name())
	}
	if err := component.Register(); err != nil {
		panic(err)
	}
}

func (app *Application) unregister(component Component) {
	if app.o.Debug {
		log.Println("Unregistering Component: " + component.Name())
	}
	component.Unregister()
}

// Use injects the components to the Bone Application, combines other
// component into a interconnected system.
func (app *Application) Use(components ...Component) {
	for _, component := range components {
		err := component.Init()
		if err != nil {
			panic(err)
		}

		app.g.Provide(&inject.Object{Name: component.Name(), Value: component})
	}

	if err := app.g.Populate(); err != nil {
		panic(err)
	}

	for _, component := range components {
		app.register(component)
	}

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		<-ch
		defer os.Exit(0)
		if app.o.Debug {
			defer log.Println("Application stopped")
		}

		for _, component := range components {
			defer app.unregister(component)
		}
	}()
}

// Run starts up the Bone Application.
func (app *Application) Run() {
	if app.o.HttpEnable {
		go func() {
			addr := fmt.Sprintf("%s:%d", app.o.HttpHost, app.o.HttpPort)
			s := &http.Server{
				Addr:    addr,
				Handler: app.httpRouter,
			}
			if app.o.Debug {
				log.Println("HTTP Server listens on " + addr)
			}
			s.ListenAndServe()
		}()
	}

	if app.o.GrpcEnable {
		go func() {
			addr := fmt.Sprintf("%s:%d", app.o.GrpcHost, app.o.GrpcPort)
			lis, err := net.Listen("tcp", addr)
			if err != nil {
				log.Fatalf("Failed to listen: %v", err)
			}
			if app.o.Debug {
				log.Println("GRPC Server listens on " + addr)
			}
			app.grpcServer.Serve(lis)
		}()
	}

	select {
	// Dothing to do, just block here, in order to keep this application on.
	}
}
