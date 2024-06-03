package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{} // mutal exclution
var waitGroup = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		waitGroup.Add(1)
		go dbCall(i)
	}
	waitGroup.Wait()
	fmt.Printf("\nTotal Execution Time: %v", time.Since(t0))
	fmt.Printf("\nThe Results are: %v", results)
}

func dbCall(i int) {
	//Simnulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	waitGroup.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}
