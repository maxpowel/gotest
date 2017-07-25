package main
import (
	wmachinery "github.com/maxpowel/dislet/machinery"
	"github.com/maxpowel/dislet/database/gorm"
	"github.com/maxpowel/dislet/mqtt"
	"github.com/maxpowel/wiphonego/task"
	"github.com/maxpowel/wiphonego/controller"
	"github.com/maxpowel/dislet"
	"github.com/maxpowel/dislet/apirest"
	"github.com/maxpowel/dislet/crypto"
)

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
	//mv := NewFetcher(wiphonego.Credentials{username:"alvaro_gg@hotmail.com"})
	//mv.SetPlainPassword("MBAR4B1")
	//mv := NewFetcher(wiphonego.Credentials{username:"maxpowel@gmail.com", password:"TD2nWhG6"})
	//mv.SetPlainPassword("TD2nWhG6")
	//c ,_ := mv.getInternetConsumption("677077536")
	//fmt.Println(c)
	f := []func(k *dislet.Kernel){
		mqtt.Bootstrap,
		gorm.Bootstrap,
		wmachinery.Bootstrap,
		apirest.Bootstrap,
		task.Bootstrap,
		controller.Bootstrap,
		crypto.Bootstrap,
	}
	dislet.Boot(f)
}



/*

curl -L 'https://yosoymas.masmovil.es/validate/' -H 'Cookie: sid=9ac08fb339ce536a7ba413d0e693aa16; visid_incap_967703=2gznSSTSRAC4pd0nuFJ6PJ8UXlkAAAAAQUIPAAAAAABBD81Yp+ofpzGxuVrdvKYk; incap_ses_504_967703=H1XcZLM41FlNR+3YupH+Bp8UXlkAAAAAKFFt8Zsdu3niR4mFkpxn4g==;'   -H 'Content-Type: application/x-www-form-urlencoded'   --data 'action=login&url=&user=alvaro_gg%40hotmail.com&password=MBAR4B1' --compressed | grep MC171108950
curl -L 'https://yosoymas.masmovil.es/validate/' -H 'Cookie: sid=5a8d29de1ebc29490dc0df10175961eb; visid_incap_967703=45Qk9FfsT3yEVrGIgHZC8IIYXlkAAAAAQUIPAAAAAABeg47sejvt+3KHSOmAV0fd; incap_ses_504_967703=qEH7T85oJBnBu+3YupH+BoIYXlkAAAAAuu9JBq4oSfSXFUTugS2SgQ==;'   -H 'Content-Type: application/x-www-form-urlencoded'   --data 'action=login&url=&user=alvaro_gg%40hotmail.com&password=MBAR4B1' --compressed | grep MC171108950



 */