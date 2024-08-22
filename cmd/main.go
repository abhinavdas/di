package main

import (
	"di/registry"
	"di/services"
)

func main() {
	reg := registry.MakeRegistry()
	reg.Register(services.ConcreteAB{})
	svc, _ := reg.Get("ab")
	sA := svc.(services.ServiceA)
	sA.Foo()

}
