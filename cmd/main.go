package main

import ("log"
		"sync"
		"time"
		"example.com/m/app")

func main(){
	wg := sync.WaitGroup{}
	cfg, err := app.NewConfig()
	if err != "" {
		panic(err)
	}

	wg.Add(1)
	go func(){
		defer wg.Done()
		for {
		
			log.Printf("prt :%v", cfg.Port)
			time.Sleep(30 * time.Second)
		}
	}()

	log.Printf("prt :%v", cfg.Port)
	wg.Wait()
}