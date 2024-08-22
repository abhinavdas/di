package main

import (
	"di/registry"
	"di/services"
)

func main() {
	reg := registry.MakeRegistry()
	reg.Register(services.ConcreteAB{})
	svc, _ := registry.TypedGet[services.ServiceA](reg, "ab")
	svc.Foo()

}
