package main

import (
	"github.com/fatih/color"
	"github.com/RichardKnop/machinery/v1"
	"fmt"
	"github.com/RichardKnop/machinery/v1/config"
)

type MachineryConfig struct {
	Broker string `default:"valor"`
	ResultBackend string `default:"PUERTOOO"`
	DefaultQueue string
}

/*func NewConnection(dialect string, uri string) *gorm.DB {
	db, _ := gorm.Open(dialect, uri)
	fmt.Println("CREANDO CONEXION")
	return db
}*/


func Add(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}

func GetConsumptionTask(username string, password string, operator string) (int64, error){
	if operator == "masmovil" {
		mv := NewMasMovilFetcher(Credentials{username: username, password: password})
		//mv := NewPepephoneFetcher(Credentials{username:"maxpowel@gmail.com", password:"TD2nWhG6"})
		c, err := mv.getInternetConsumption("677077536")
		return c.consumed, err
	}

	return 0, fmt.Errorf("Operator \"%v\" not available", operator)

}

func GetAnonymousConsumptionTask(username string, password string, operator string, deviceId string) (int64, error){
	if operator == "masmovil" {
		mv := NewMasMovilFetcher(Credentials{username: username, password: password})
		//mv := NewPepephoneFetcher(Credentials{username:"maxpowel@gmail.com", password:"TD2nWhG6"})
		c, err := mv.getInternetConsumption("677077536")
		return c.consumed, err
	}

	return 0, fmt.Errorf("Operator \"%v\" not available", operator)

}

func machineryBootstrap(k *Kernel) {
	//fmt.Println("DATABASE BOOT")
	mapping := k.config.mapping
	mapping["machinery"] = &MachineryConfig{}
	var baz OnKernelReady = func(k *Kernel){
		color.Green("Evento en machinery")
		machineryConfig := k.config.mapping["machinery"].(*MachineryConfig)
		fmt.Println(machineryConfig)

		var cnf = config.Config{
			Broker: machineryConfig.Broker,
			ResultBackend: machineryConfig.ResultBackend,
			DefaultQueue: machineryConfig.DefaultQueue,
			//Broker : "redis://localhost:6379/0",
			//ResultBackend: "redis://localhost:6379/0",
			//Broker:             "amqp://guest:guest@localhost:5672/",
			//ResultBackend:      "amqp://guest:guest@localhost:5672/",
			//DefaultQueue: "machinery_tasks",
			/*AMQP:               &config.AMQPConfig{
				Exchange:     "machinery_exchange",
				ExchangeType: "direct",
				BindingKey:   "machinery_task",
			},*/
		}
		fmt.Println(cnf)

		server, err := machinery.NewServer(&cnf)
		if err != nil {
			panic(err)
		}

		// Register tasks
		fmt.Println("Registrar tasks")
		taskList := map[string]interface{}{
			"add":        Add,
			"consumption": GetConsumptionTask,
			"anonymousConsumption": GetAnonymousConsumption,
		}
		err = server.RegisterTasks(taskList)

		runWorker := func (server *machinery.Server) {
			worker := server.NewWorker("machinery_worker")
			if err := worker.Launch(); err != nil {
				fmt.Println("Error worker")
				fmt.Println(err)
			}
		}

		go runWorker(server)

		//k.container.RegisterType("database", NewConnection, conf.Dialect, conf.Uri)
		iny := func() *machinery.Server{
			return server
		}
		k.container.RegisterType("machinery", iny)


	}
	k.subscribe(baz)




}