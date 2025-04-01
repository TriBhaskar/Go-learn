package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var rwm = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() { t0 := time.Now()
	// for i:=0; i<len(dbData); i++{
	// 	wg.Add(1)
	// 	go dbCall(i)
	// }
	// wg.Wait()
	// fmt.Printf("\n Total time taken: %v\n", time.Since(t0))
	// fmt.Println("The results are: ", results)

	for i:=0; i<1000; i++{
		wg.Add(1)
		go performaceTest()
	}
	wg.Wait()
	fmt.Printf("\n Total time taken: %v\n", time.Since(t0))
}

func performaceTest() {
	var res int
	for i:=0; i<100000000; i++{
		res += i
	}
	wg.Done()
}

func dbCall(i int) {

	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	// fmt.Println("The result from the database is :", dbData[i])
	// m.Lock()
	// results = append(results, dbData[i])
	save(dbData[i])
	log()
	// m.Unlock()
	wg.Done()
}

func save(result string){
	rwm.Lock()
	results = append(results, result)
	rwm.Unlock()
}

func log(){
	rwm.RLock()
	fmt.Println("The current results are: ", results)
	rwm.RUnlock()
}
