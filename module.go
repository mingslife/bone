package bone

type Model interface {
}

type Repo interface {
}

type Service interface {
}

type Endpoint interface {
}

type Transport interface {
	Register() error
}

type Module interface {
	Component
}
