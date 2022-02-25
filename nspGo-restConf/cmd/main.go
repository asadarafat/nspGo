package main

import (
	"os"
	"sync"
	"time"

	// "github.com/rifflock/lfshook"
	// "github.com/sirupsen/logrus"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"

	nspgorestconf "local.com/nspgo/nspGo-restConf"
	nspgosession "local.com/nspgo/nspGo-session"
)

// var log *logrus.Logger

// func NewLogger() *logrus.Logger {
// 	if log != nil {
// 		return log
// 	}

// 	pathMap := lfshook.PathMap{
// 		logrus.InfoLevel: "./nspGo-restConf/resconf-inventory-payload.log",
// 	}

// 	log = logrus.New()
// 	log.Hooks.Add(lfshook.NewHook(
// 		pathMap,
// 		&logrus.TextFormatter{
// 			FullTimestamp: true},
// 	))
// 	return log
// }

func main() {

	// init class session
	p := nspgosession.Session{}
	p.LoadConfig()

	log.Info("nsp.nspOsIP :", p.IpAdressNspOs)
	log.Info("nsp.nspIprcIP :", p.IpAdressIprc)
	log.Info("nsp.Username :", p.Username)
	log.Info("nsp.Password :", p.Password)
	log.Info("nsp.linetoken :", p.Token)

	p.EncodeUserName()
	log.Info(p.EncodeUserName())

	p.GetRestToken()
	log.Info("nsp.linetoken_NEW :", p.Token)

	// init class RestConf
	rc := nspgorestconf.RestConf{}

	//get list NE
	rc.ReadRestConfPayload("./nspGo-restConf/resconf-inventory-payload.json")
	restconfPayloadInventory := rc.Payload
	inventoryJson := rc.NspRestconfInventory(p.IpAdressIprc, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, "operations/nsp-inventory:find", restconfPayloadInventory)
	value := gjson.Get(inventoryJson, "nsp-inventory:output.data.#.ne-id")
	// log.Info(value.String())

	listOfNeId := []string{}
	value.ForEach(func(key, value gjson.Result) bool {
		//println(value.String())
		listOfNeId = append(listOfNeId, value.String())
		return true // keep iterating
	})
	// log.Info("listOfNeId: ", listOfNeId)

	// get payload
	pathToPayload := "./nspGo-restConf/resconf-payload.json"
	file, err := os.Stat(pathToPayload)
	if err != nil {
		return
	}
	log.Info("Payload Size(bytes): ", file.Size())

	rc.ReadRestConfPayload(pathToPayload)
	restconfPayloadCreate := rc.Payload

	// Running Restconf Request In Sequence
	//measure the start time
	startRestconfSequence := time.Now()
	log.Info("Start Sequence: ", startRestconfSequence)

	log.Info("Running Sequence: ", len(listOfNeId))
	value.ForEach(func(key, value gjson.Result) bool {
		//println(value.String())
		rc.PatchRestConfMdc(p.IpAdressIprc, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, value.String(), "/root", restconfPayloadCreate, false)
		return true // keep iterating
	})

	log.Info("Finished Sequence")
	log.Info("Elapsed Time For Sequence: ", time.Since(startRestconfSequence))

	startRestconfConcurrent := time.Now()
	log.Info("Start Sequence: ", startRestconfConcurrent)
	// Running Restconf Request In Concurrent
	log.Info("Running Concurrency: ", len(listOfNeId))
	var waitingGroupNeList sync.WaitGroup
	waitingGroupNeList.Add(len(listOfNeId))

	for j := 0; j < len(listOfNeId); j++ {
		go func(j int) {
			// fmt.Println(listOfNeId[j])
			rc.PatchRestConfMdc(p.IpAdressIprc, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, listOfNeId[j], "/root", restconfPayloadCreate, false)
			waitingGroupNeList.Done()
		}(j)
	}
	waitingGroupNeList.Wait()

	log.Info("Finished Concurrent")
	log.Info("Elapsed Time For Concurrent: ", time.Since(startRestconfConcurrent))

	p.RevokeRestToken()
}
