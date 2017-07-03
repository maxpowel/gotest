package main

import (
	"github.com/fgrosse/goldi"
	"github.com/fgrosse/goldi/validation"
)

type Kernel struct {
	event chan int
	config *Config
	registry goldi.TypeRegistry
	container *goldi.Container
}


func newKernel(configPath, paramsPath string, bootstrapModules []func(k *Kernel)) *Kernel {
	k := Kernel{}
	k.event = make(chan int)
	k.registry = goldi.NewTypeRegistry()
	k.container = goldi.NewContainer(k.registry, nil)
	k.container.RegisterType("config", NewConfig, configPath, paramsPath)

	//TODO hacer defer db.Close()
	// Config must be created before module bootstraping
	k.config = k.container.MustGet("config").(*Config)
	for _, f := range bootstrapModules {
		go f(&k)
	}

	// Check that container is ok
	validator := validation.NewContainerValidator()
	validator.MustValidate(k.container)
	// Load configuration
	k.config.load()
	k.container.Config = k.config.mapping

	k.event <- 1
	k.event <- 1
	return &k
}