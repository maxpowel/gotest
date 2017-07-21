package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/fatih/color"
	"github.com/maxpowel/dislet"
)

type Config struct {
	Dialect string `default:"valor"`
	Uri string `default:"PUERTOOO"`
}


func NewConnection(dialect string, uri string) *gorm.DB {
	db, _ := gorm.Open(dialect, uri)
	fmt.Println(uri)
	//db.LogMode(true)
	fmt.Println("CREANDO CONEXION")
	//db.AutoMigrate(&User{})
	/*u := NewUser()
	PlainPassword(&u, "123456")
	db.Create(&u)
	db.Commit()*/
	return db
}


func Bootstrap(k *dislet.Kernel) {
	//fmt.Println("DATABASE BOOT")
	mapping := k.Config.Mapping
	mapping["database"] = &Config{}

	var baz dislet.OnKernelReady = func(k *dislet.Kernel){
		color.Green("Evento en database")
		conf := k.Config.Mapping["database"].(*Config)
		k.Container.RegisterType("database", NewConnection, conf.Dialect, conf.Uri)
	}
	k.Subscribe(baz)




}
