package task

import "fmt"
import "github.com/maxpowel/wiphonego/fetcher/masmovil"
import (
	"github.com/maxpowel/wiphonego"
	"github.com/maxpowel/dislet"
	"github.com/RichardKnop/machinery/v1"
	"log"
	"github.com/jinzhu/gorm"
)

var kernel *dislet.Kernel

func GetConsumptionTask(username string, password string, operator string) (wiphonego.UserDeviceConsumption, error){
	if operator == "masmovil" {
		mv := masmovil.NewFetcher(&wiphonego.Credentials{Username: username, Password: password})
		//mv := NewFetcher(Credentials{username:"maxpowel@gmail.com", password:"TD2nWhG6"})
		c, err := mv.GetInternetConsumption("677077536")
		return c, err
	}

	return wiphonego.UserDeviceConsumption{}, fmt.Errorf("Operator \"%v\" not available", operator)

}

func GetAnonymousConsumptionTask(username string, password string, operator string, deviceId string) (string, map[string]string, error){

	db := kernel.Container.MustGet("database").(*gorm.DB)


	if operator == "masmovil" {
		mv := masmovil.NewFetcher(&wiphonego.Credentials{Username: username, Password: password})
		//mv := NewFetcher(Credentials{username:"maxpowel@gmail.com", password:"TD2nWhG6"})
		c, err := mv.GetInternetConsumption("677077536")
		if err == nil {
			device := wiphonego.UserDevice{}
			db.Where("Uuid = ?", deviceId).FirstOrCreate(&device, wiphonego.UserDevice{Uuid: deviceId})
			//device.Uuid = device
			c.Device = device
			db.Create(&c)
		}
		return "", nil, err
	}
	params := make(map[string]string)
	params["operator"] = operator
	return "Operator {operator} is not available", params, nil

}
//fmt.Errorf("Operator \"%v\" is not available", operator)

func Bootstrap(k *dislet.Kernel) {
	kernel = k
	var baz dislet.OnKernelReady = func(k *dislet.Kernel){
		server := k.Container.MustGet("machinery").(*machinery.Server)
		// Register tasks

		err := server.RegisterTasks(map[string]interface{}{
			"consumption": GetConsumptionTask,
			"anonymousConsumption": GetAnonymousConsumptionTask,
		})

		if err != nil {
			log.Fatal(err)
		}
	}

	k.Subscribe(baz)

}
