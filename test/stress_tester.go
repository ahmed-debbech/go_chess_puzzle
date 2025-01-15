package main

import (
	"log"
	"net/http"
	"fmt"
	"sync"
)

const (
	host = "https://pichess.org"
	MAX = 100
	//host = "https://google.com"
)

func main(){
	log.Println("Stress Teser Starting...")

	var wg sync.WaitGroup

	sch := make(chan bool, MAX)
	fch := make(chan bool, MAX)
	success := 0
	fail := 0

	go func(){
		for {
			select {
			case _= <- sch:
				success++
			case _= <- fch:
				fail++
			}
		}
	}()

	for i := 1; i<=MAX; i++{
		wg.Add(1)
		go func(k int) {
			defer wg.Done()

			_, err := http.Get(fmt.Sprintf("%s/%s", host ,""))
			
			if err != nil {
				fch <- true
				return
			}
			sch <- true
		}(i)
	
	}

	wg.Wait()
	close(fch)
	close(sch)

	log.Println("SUCCESS:", success, "FAIL:", fail)
}