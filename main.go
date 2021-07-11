package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/sanjayshr/gocourse/utils"
)

func main() {

	// Mutex example
	rand.Seed(time.Now().UnixNano())
	wg := &sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go utils.Save(fmt.Sprintf("user_id_%d", rand.Intn(100)), rand.Intn(100), wg)
	}
	wg.Wait()
	fmt.Println(utils.InMemCache.M)
}
