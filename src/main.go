package main
import "fmt"
import "os"
import (

	"os/signal"
	"syscall"

	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/go-sql-driver/mysql"
	"github.com/fatih/color"
	"flag"
)

//////////////

func main() {
	// Parse parameters
	configPtr := flag.String("config", "config.yml", "Configuration file")
	parametersPtr := flag.String("parameters", "parameters.yml", "Parameters file")
	flag.Parse()
	color.Green("Starting...")
	// Dependency injection container
	f := []func(k *Kernel){mqttBootstrap, databaseBootstrap}

	kernel := newKernel(*configPtr, *parametersPtr, f)

	fmt.Println(kernel.config.mapping)


	//db2 := container.MustGet("database").(*gorm.DB)
	//db2.Create(&PlayList{Source: "AAA", SourceType: "BBB"})

	//fmt.Println(conf)
	//fmt.Println(conf["mqtt"].(map[interface{}]interface{})["topic"])

	daemonize()
}

func daemonize(){
	// Daemonize
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		color.Yellow("Interrupt \"%v\" received, exiting", sig)

		done <- true
	}()

	<-done

}
