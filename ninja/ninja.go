package ninja

import (
	"fmt"
	"sync"
	"time"
)

func NinjaMain() {

	evilNinjas := []string{"Tommy", "Jhony", "Boby", "Andy"}

	wg := &sync.WaitGroup{}
	wg.Add(len(evilNinjas))

	start := time.Now()

	defer func() {
		fmt.Println("time spent: ", time.Since(start))
		fmt.Printf("Round %s\n", time.Since(start).Round(time.Second))
	}()

	for _, evilNinja := range evilNinjas {
		go attack(evilNinja, wg)
	}

	wg.Wait()

}

func attack(target string, wg *sync.WaitGroup) {
	fmt.Println("Throwing ninja star at -> ", target)
	time.Sleep(time.Second)
	wg.Done()

}
