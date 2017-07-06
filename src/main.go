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
	cosa()
	return
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


/*

curl -L 'https://yosoymas.masmovil.es/validate/' -H 'Cookie: sid=9ac08fb339ce536a7ba413d0e693aa16; visid_incap_967703=2gznSSTSRAC4pd0nuFJ6PJ8UXlkAAAAAQUIPAAAAAABBD81Yp+ofpzGxuVrdvKYk; incap_ses_504_967703=H1XcZLM41FlNR+3YupH+Bp8UXlkAAAAAKFFt8Zsdu3niR4mFkpxn4g==;'   -H 'Content-Type: application/x-www-form-urlencoded'   --data 'action=login&url=&user=alvaro_gg%40hotmail.com&password=MBAR4B1' --compressed | grep MC171108950
curl -L 'https://yosoymas.masmovil.es/validate/' -H 'Cookie: sid=5a8d29de1ebc29490dc0df10175961eb; visid_incap_967703=45Qk9FfsT3yEVrGIgHZC8IIYXlkAAAAAQUIPAAAAAABeg47sejvt+3KHSOmAV0fd; incap_ses_504_967703=qEH7T85oJBnBu+3YupH+BoIYXlkAAAAAuu9JBq4oSfSXFUTugS2SgQ==;'   -H 'Content-Type: application/x-www-form-urlencoded'   --data 'action=login&url=&user=alvaro_gg%40hotmail.com&password=MBAR4B1' --compressed | grep MC171108950



 */