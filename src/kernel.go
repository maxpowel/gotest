package main

import (
	"github.com/fgrosse/goldi"
	"github.com/fgrosse/goldi/validation"
	"fmt"
)

type Kernel struct {
	event []OnKernelReady
	config *Config
	registry goldi.TypeRegistry
	container *goldi.Container

}

type OnKernelReady func(k *Kernel)

func (k *Kernel) subscribe(ready OnKernelReady) {
	k.event = append(k.event, ready)
	fmt.Printf("NOTIFICANDO %v", len(k.event))
}


func newKernel(configPath, paramsPath string, bootstrapModules []func(k *Kernel)) *Kernel {
	k := Kernel{}

	//k.event = make(chan int)
	k.registry = goldi.NewTypeRegistry()
	k.container = goldi.NewContainer(k.registry, nil)
	k.container.RegisterType("config", NewConfig, configPath, paramsPath)

	//TODO hacer defer db.Close()
	// Config must be created before module bootstraping
	k.config = k.container.MustGet("config").(*Config)
	for _, f := range bootstrapModules {
		f(&k)
	}

	// Check that container is ok
	validator := validation.NewContainerValidator()
	validator.MustValidate(k.container)
	// Load configuration
	k.config.load()
	k.container.Config = k.config.mapping

	// On kernel ready event

	for _, f := range k.event {
		fmt.Println("AAAAAA")
		f(&k)
	}


	return &k
}