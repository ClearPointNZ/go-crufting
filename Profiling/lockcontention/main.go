package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

const (
	pprofAddress = "localhost:6060"
)

var (
	mutex     sync.Mutex
	waitGroup sync.WaitGroup
)

func main() {

	// Expose the pprof HTTP debugging endpoints:
	go func() {
		log.Printf("Starting PPROF (http://%s/debug/pprof\n", pprofAddress)

		if err := http.ListenAndServe(pprofAddress, nil); err != nil {
			log.Fatalf("Error serving pprof: %s", err)
		}
	}()

	// Start some background lock contention:
	go lockUnlock(2 * time.Second)
	go lockUnlock(3 * time.Second)
	go lockUnlock(4 * time.Second)

	time.Sleep(time.Minute)

	// Wait until everything has finished:
	// waitGroup.Wait()
}

// lockUnlock will attempt to lock and unlock the mutex on a schedule:
func lockUnlock(interval time.Duration) {

	log.Printf("LockUnlock (%s)", interval.String())

	// Tell the wait group to wait for us to finish:
	waitGroup.Add(1)

	// Do this 5 times:
	for i := 0; i < 5; i++ {
		log.Printf("Locking (%s)", interval.String())
		mutex.Lock()

		time.Sleep(interval)

		log.Printf("Unlocking (%s)", interval.String())
		mutex.Unlock()
	}

	// Tell the wait group we're finished:
	waitGroup.Done()
}
