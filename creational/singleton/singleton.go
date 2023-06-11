package singleton

import (
	"log"
	"sync"
)

var locker = &sync.Mutex{}

type Singleton struct {
}

var singleInstance *Singleton

func getSingleton() *Singleton {
	if singleInstance != nil {
		log.Println("Instance already created")
		return singleInstance
	}
	locker.Lock()
	singleInstance = &Singleton{}
	locker.Unlock()
	log.Println("Instance created")
	return singleInstance
}
