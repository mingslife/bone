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

// ApplicationOptions represents the options to create Bone Application
// instance.
type ApplicationOptions struct {
	Host string
	Port int
}

// DefaultApplicationOptions returns the default options to create Bone
// Application instance.
func DefaultApplicationOptions() *ApplicationOptions {
	return &ApplicationOptions{
		Host: "127.0.0.1",
		Port: 8080,
	}
}

// Application defines the struct of Bone Application.
type Application struct {
	o *ApplicationOptions
	g *inject.Graph
	r *Router
}

// NewApplication returns the instance of Bone Application.
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
		defer log.Println("Application stopped")

		for _, component := range components {
			defer app.unregister(component)
		}
	}()
}

// Run starts up the Bone Application.
func (app *Application) Run() {
	addr := fmt.Sprintf("%s:%d", app.o.Host, app.o.Port)
	s := &http.Server{
		Addr:    addr,
		Handler: app.r,
	}
	log.Println("Listen on " + addr)
	s.ListenAndServe()
}
