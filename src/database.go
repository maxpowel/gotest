package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/fatih/color"
)

type DatabaseConfig struct {
	Dialect string `default:"valor"`
	Uri string `default:"PUERTOOO"`
}


func NewConnection(dialect string, uri string) *gorm.DB {
	db, _ := gorm.Open(dialect, uri)
	fmt.Println("CREANDO CONEXION")
	return db
}


func databaseBootstrap(k *Kernel) {
	//fmt.Println("DATABASE BOOT")
	mapping := k.config.mapping
	mapping["database"] = &DatabaseConfig{}

	var baz OnKernelReady = func(k *Kernel){
		color.Green("Evento en database")
		conf := k.config.mapping["database"].(*DatabaseConfig)
		k.container.RegisterType("database", NewConnection, conf.Dialect, conf.Uri)
	}
	k.subscribe(baz)




}
