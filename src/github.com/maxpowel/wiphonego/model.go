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
/*
func(c *Credentials) SetPlainPassword(plainPassword string) (error) {
	key := []byte("example key 1234")
	plaintext := []byte(plainPassword)

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	c.Password = base64.URLEncoding.EncodeToString(ciphertext)
	return nil
}

func(c *Credentials) GetPlainPassword() (string, error){
	out := c.Password
	key := []byte("example key 1234")
	ciphertext, err := base64.URLEncoding.DecodeString(out)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream2 := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream2.XORKeyStream(ciphertext, ciphertext)
	//fmt.Printf("%s", []byte(ciphertext2))
	return string(ciphertext), nil
}*/