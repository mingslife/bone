package bone

// Component defines the common methods
type Component interface {
	Name() string
	Init() error
	Register() error
	Unregister() error
}
