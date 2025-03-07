package singleton

import "sync"

// SingletonStruct is a struct that will be used as a singleton
type SingletonStruct struct {
	data string
}

var instance *SingletonStruct // Singleton instance of SingletonStruct
var syncOnce sync.Once        // sync.Once instance to make sure that the singleton is created only once

// GetInstance is a function that will return the singleton instance
func GetInstance() *SingletonStruct {
	// Using sync.Once to make sure that the singleton is created only once
	syncOnce.Do(func() {
		instance = &SingletonStruct{
			data: "Singleton Data",
		}
	})
	// Returning the singleton instance of SingletonStruct. This will be the same instance every time this function is called
	return instance
}
