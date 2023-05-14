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
			log.Printf("port :%v", cfg.Port)
			log.Printf("rdbs user :%v", cfg.Rdbs.User)
			log.Printf("rdbs password :%v", cfg.Rdbs.Password)
			log.Printf("rdbs port :%v", cfg.Rdbs.Port)
			time.Sleep(30 * time.Second)
		}
	}()



	wg.Wait()
}