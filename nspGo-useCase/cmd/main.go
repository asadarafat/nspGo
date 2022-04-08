package main

import (
	"time"

	log "github.com/sirupsen/logrus"
	nspgousecase "local.com/nspgo/nspGo-useCase"
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}
func main() {

	// nspgousecase.ThalesLookupWithIntentCreateIntents()
	// nspgousecase.ThalesLookupWithIntentGetIntents()
	// nspgousecase.ThalesLookupWithResourceManagerObtain(1, 10)
	// nspgousecase.ThalesLookupWithResourceManagerRelease()

	// nspgousecase.ThalesLookupWithGoCache(1, 1000000)

	nspgousecase.ThalesRestConf4kSvc()

	// nspgousecase.ThalesRestConfMdc()
	// nspgousecase.ThalesRestConf8545()

	// doEvery(10*time.Second, nspgousecase.ThalesRestConf4kSvc)

	ticker := time.NewTicker(10 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				log.Info("Tick at", t)
			}
		}
	}()
	time.Sleep(60 * time.Second)
	ticker.Stop()
	done <- true
	log.Info("Ticker stopped")

}
