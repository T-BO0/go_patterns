package singleton

import "sync"

type SingletonStruct struct {
	data string
}

var instance *SingletonStruct
var syncOnce sync.Once

func GetInstance() *SingletonStruct {
	syncOnce.Do(func() {
		instance = &SingletonStruct{
			data: "Singleton Data",
		}
	})
	return instance
}
