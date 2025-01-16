package main

import (
	"log"
	"net/http"
	"fmt"
	"sync"
	"os"
	"os/signal"
)

const (
	host = "https://pichess.org"
	MAX = 5000
	//host = "http://localhost:5530"
)

var (
	success = 0
	fail = 0
)

func Dump(){
	var wg sync.WaitGroup

	sch := make(chan bool, MAX)
	fch := make(chan bool, MAX)

	defer close(fch)
	defer close(sch)

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
			
			log.Print(k)
			if err != nil {
				fch <- true
				return
			}
			sch <- true
		}(i)
	
	}

	wg.Wait()


	log.Println("SUCCESS:", success, "FAIL:", fail)
}

func Serial(){

	success := 0
	fail := 0

	for i := 1; i<=MAX; i++{
		_, err := http.Get(fmt.Sprintf("%s/%s", host ,""))
		
		log.Print(i)
		if err != nil {
			fail++
			
		}else{
			success++
		}
	}
	log.Println("SUCCESS:", success, "FAIL:", fail)

}

func main(){
	log.Println("Stress Tester Starting...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		for  _ = range c {
			log.Println("SUCCESS:", success, "FAIL:", fail)
			os.Exit(1)
		}
	}()

	if os.Args[1] == "dump" {
		log.Println("dump mode")
		Dump()
	}else{
		Serial()
	}
}