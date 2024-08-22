package services

type ServiceA interface {
	Foo() int
}

type ServiceB interface {
	Bar() int
}

type ConcreteAB struct {
	val int
}

func (c ConcreteAB) Foo() int {
	return c.val
}

func (c ConcreteAB) Bar() int {
	return c.val * c.val
}

func (c ConcreteAB) GetName() string {
	return "ab"
}
