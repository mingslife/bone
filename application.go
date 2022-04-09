package bone

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/facebookgo/inject"
)

type ApplicationOptions struct {
	Host string
	Port int
}

func DefaultApplicationOptions() *ApplicationOptions {
	return &ApplicationOptions{
		Host: "127.0.0.1",
		Port: 8080,
	}
}

type Application struct {
	o *ApplicationOptions
	g *inject.Graph
	r *Router
}

func NewApplication(options *ApplicationOptions) *Application {
	app := &Application{
		o: options,
		g: &inject.Graph{},
		r: NewRouter(),
	}

	app.g.Provide(&inject.Object{Name: "application.router", Value: app.r})

	return app
}

func (*Application) register(component Component) {
	log.Println("Registering Component: " + component.Name())
	if err := component.Register(); err != nil {
		panic(err)
	}
}

func (*Application) unregister(component Component) {
	log.Println("Unregistering Component: " + component.Name())
	component.Unregister()
}

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
		defer log.Println("Application stopped")

		for _, component := range components {
			defer app.unregister(component)
		}
	}()
}

func (app *Application) Run() {
	addr := fmt.Sprintf("%s:%d", app.o.Host, app.o.Port)
	s := &http.Server{
		Addr:    addr,
		Handler: app.r,
	}
	log.Println("Listen on " + addr)
	s.ListenAndServe()
}
