package ramstore

import (
	"fmt"
)

type RamStore struct{
	store map[string]string
}

var ramStoreInstance *RamStore = nil

func newRamStore() *RamStore{
	fmt.Println("Creating new RamStore")
	return &RamStore{ store: make(map[string]string) }
}

func GetRamStoreInstance() *RamStore{
	if ramStoreInstance == nil {
		ramStoreInstance = newRamStore()
	}
	return ramStoreInstance
}

func Set(pid string, hash string){
	ramStoreInstance.store[pid] = hash
}