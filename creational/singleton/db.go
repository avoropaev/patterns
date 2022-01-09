package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once
type db struct {}
var dbInstance *db

func getInstance() *db {
	if dbInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating DB instance now.")
				dbInstance = &db{}
			})
	} else {
		fmt.Println("DB instance already created.")
	}

	return dbInstance
}
