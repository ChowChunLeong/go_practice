package service

import "sync"

// Singleton 是一个单例结构体
type Singleton struct {
	Data string
}

var instance *Singleton
var once sync.Once

// GetInstance 是获取 Singleton 实例的函数
func GetSingletonInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			Data: "This is a singleton instance",
		}
	})
	return instance
}

// func UpdateSingletonInstance(newSingletonData string){

// 	instance
// 	return "update_singleton_instance_data"
// }
