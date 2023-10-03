package Singleton

import (
	"fmt"
	"sync"
)

type Single struct {
}

var lock = &sync.Mutex{}
var singleInstance *Single

func getInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &Single{}
		}
	}
	return singleInstance
}

func getInstanceWithComment() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &Single{}
		} else {
			fmt.Println("Race condition: Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

var once sync.Once

func getInstanceByOnce() *Single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &Single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
