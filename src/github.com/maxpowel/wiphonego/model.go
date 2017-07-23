package wiphonego

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserDeviceConsumption struct {
	gorm.Model
	InternetConsumed int64
	InternetTotal int64
	CallConsumed int
	CallTotal int
	RenewTime time.Time
	Device UserDevice
	DeviceId int
}



type UserDevice struct {
	gorm.Model
	Uuid string
	Consumptions []UserDeviceConsumption

}




type Operator struct{
	gorm.Model
	Name string
}

type Credentials struct {

	Operator Operator
	Device string
	Username string
	Password string
}