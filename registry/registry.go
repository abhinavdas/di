package registry

// Registrable Define a Registrable interface.
type Registrable interface {
	GetName() string
}

type Registry struct {
	registry map[string]Registrable
}

// MakeRegistry Registry to hold registrable entities
func MakeRegistry() *Registry {
	return &Registry{registry: make(map[string]Registrable)}
}

// Register function to add an entity to the registry
func (r *Registry) Register(entity Registrable) {
	name := entity.GetName()
	r.registry[name] = entity
}

// Get function to retrieve an entity by name
func (r *Registry) Get(name string) (any, bool) {
	entity, found := r.registry[name]
	return entity, found
}
