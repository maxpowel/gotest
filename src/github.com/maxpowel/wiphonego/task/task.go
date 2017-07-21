package task

import "fmt"
import "github.com/maxpowel/wiphonego/fetcher/masmovil"
import (
	"github.com/maxpowel/wiphonego"
	"github.com/maxpowel/dislet"
	"github.com/RichardKnop/machinery/v1"
	"log"
)

func GetConsumptionTask(username string, password string, operator string) (int64, error){
	if operator == "masmovil" {
		mv := masmovil.NewFetcher(&wiphonego.Credentials{Username: username, Password: password})
		//mv := NewFetcher(Credentials{username:"maxpowel@gmail.com", password:"TD2nWhG6"})
		c, err := mv.GetInternetConsumption("677077536")
		return c.Consumed, err
	}

	return 0, fmt.Errorf("Operator \"%v\" not available", operator)

}

func GetAnonymousConsumptionTask(username string, password string, operator string, deviceId string) (int64, error){
	if operator == "masmovil" {
		mv := masmovil.NewFetcher(&wiphonego.Credentials{Username: username, Password: password})
		//mv := NewFetcher(Credentials{username:"maxpowel@gmail.com", password:"TD2nWhG6"})
		c, err := mv.GetInternetConsumption("677077536")
		return c.Consumed, err
	}

	return 0, fmt.Errorf("Operator \"%v\" not available", operator)

}


func Bootstrap(k *dislet.Kernel) {
	var baz dislet.OnKernelReady = func(k *dislet.Kernel){
		server := k.Container.MustGet("machinery").(*machinery.Server)
		// Register tasks
		fmt.Println("Registrar tasks")

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
