package main
import "fmt"
import (

	"github.com/fatih/color"
	"flag"
	"github.com/maxpowel/dislet/machinery"
	"github.com/maxpowel/dislet/database/gorm"
	"github.com/maxpowel/dislet/mqtt"
	"github.com/maxpowel/dislet"
)

//////////////





func main() {
	//u := NewUser()
	//PlainPassword(&u, "123456")
// mirar esto https://stackoverflow.com/questions/33646948/go-using-mux-router-how-to-pass-my-db-to-my-handlers

	/*db := NewConnection("mysql", "mqtt:123456@tcp(localhost:3306)/mqtt?charset=utf8&parseTime=true")

	u := User{}
    db.Where(&User{Email:"maxpowel@gmail.com2"}).First(&u)

	fmt.Println(u.ID)*/
	//u2 := NewUser()
	//PlainPassword(&u2, "123456")
	//fmt.Println(checkPassword(&u, "123456"))
	/*fmt.Println(checkPassword(&u2, "some password"))
	fmt.Println(checkPassword(&u2, "123456"))
	db.Create(&u2)*/
	//return
	//mv := NewMasMovilFetcher(Credentials{username:"alvaro_gg@hotmail.com", password:"MBAR4B1"})
	//mv := NewPepephoneFetcher(Credentials{username:"maxpowel@gmail.com", password:"TD2nWhG6"})
	//c ,_ := mv.getInternetConsumption("677077536")
	//fmt.Println(c)



	// The second argument is a consumer tag
	// Ideally, each worker should have a unique tag (worker1, worker2 etc)

	// Parse parameters
	configPtr := flag.String("config", "config.yml", "Configuration file")
	parametersPtr := flag.String("parameters", "parameters.yml", "Parameters file")
	flag.Parse()
	color.Green("Starting...")
	// Dependency injection container
	//f := []func(k *dislet.Kernel){apiRestBootstrap}
	f := []func(k *dislet.Kernel){mqtt.Bootstrap,gorm.Bootstrap, machinery.Bootstrap, apiRestBootstrap}

	kernel := dislet.NewKernel(*configPtr, *parametersPtr, f)

	fmt.Println(kernel.Config.Mapping)


	//db2 := container.MustGet("database").(*gorm.DB)
	//db2.Create(&PlayList{Source: "AAA", SourceType: "BBB"})

	//fmt.Println(conf)
	//fmt.Println(conf["mqtt"].(map[interface{}]interface{})["topic"])

	/*
	time.Sleep(5 * time.Second)
	fmt.Println("EL OTRO HILO")

	// Enviar tarea
	task0 := tasks.Signature{
		Name: "add",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 1,
			},
		},
	}

	fmt.Println("Enviando task...")
	server := kernel.Container.MustGet("machinery").(*machinery.Server)
	asyncResult, err := server.SendTask(&task0)
	if err != nil {
		fmt.Println("Could not send task: %s", err.Error())
	}
	fmt.Println("Task enviada")

	fmt.Println("Esperando resultado ...")
	results, err := asyncResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		fmt.Println("Getting task result failed with error: %s", err.Error())
	}
	fmt.Printf(
		"%v + %v = %v\n",
		asyncResult.Signature.Args[0].Value,
		asyncResult.Signature.Args[1].Value,
		results[0].Interface(),
	)
*/
	dislet.Daemonize()
}



/*

curl -L 'https://yosoymas.masmovil.es/validate/' -H 'Cookie: sid=9ac08fb339ce536a7ba413d0e693aa16; visid_incap_967703=2gznSSTSRAC4pd0nuFJ6PJ8UXlkAAAAAQUIPAAAAAABBD81Yp+ofpzGxuVrdvKYk; incap_ses_504_967703=H1XcZLM41FlNR+3YupH+Bp8UXlkAAAAAKFFt8Zsdu3niR4mFkpxn4g==;'   -H 'Content-Type: application/x-www-form-urlencoded'   --data 'action=login&url=&user=alvaro_gg%40hotmail.com&password=MBAR4B1' --compressed | grep MC171108950
curl -L 'https://yosoymas.masmovil.es/validate/' -H 'Cookie: sid=5a8d29de1ebc29490dc0df10175961eb; visid_incap_967703=45Qk9FfsT3yEVrGIgHZC8IIYXlkAAAAAQUIPAAAAAABeg47sejvt+3KHSOmAV0fd; incap_ses_504_967703=qEH7T85oJBnBu+3YupH+BoIYXlkAAAAAuu9JBq4oSfSXFUTugS2SgQ==;'   -H 'Content-Type: application/x-www-form-urlencoded'   --data 'action=login&url=&user=alvaro_gg%40hotmail.com&password=MBAR4B1' --compressed | grep MC171108950



 */