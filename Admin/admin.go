package admin

import (
	"log"

	"auraluvsu.com/Utils"
)

type Admin struct {
	Key string
}

func (a *Admin) CreateAdminKey() []byte {
	rdByte, err := utils.RandBytes(8)
	if err != nil {
		log.Fatal(err)
	}
	key, err := utils.CreateNewHash(rdByte)
	if err != nil {
		log.Fatal(err)
	}
	return key
}
