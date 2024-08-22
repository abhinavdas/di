package services

type ServiceA interface {
	foo() int
}

type ServiceB interface {
	bar() int
}

type ConcreteAB struct {
	val int
}

func (c ConcreteAB) foo() int {
	return c.val
}

func (c ConcreteAB) bar() int {
	return c.val * c.val
}

func (c ConcreteAB) GetName() string {
	return "ab"
}
