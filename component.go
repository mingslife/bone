package bone

type Component interface {
	Name() string
	Init() error
	Register() error
	Unregister() error
}
