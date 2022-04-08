package nspgousecase

import (
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"

	nspgorestconf "local.com/nspgo/nspGo-restConf"
	nspgosession "local.com/nspgo/nspGo-session"
	nspgotools "local.com/nspgo/nspGo-tools"
)

func ThalesRestConf4kSvc() {
	t := nspgotools.Tools{}

	// log level 4: info
	// log level 5: debug
	t.InitLogger("./logs/nspGo-useCase-ThalesRestconf8545-4000svc.log", 4)

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

	if len(listOfNeId) != 0 {
		log.Info("Number of Targeted NE: ", len(listOfNeId))
	} else {
		log.Error("Restconf Gateway is unavailable")
	}

	// get RestConf payload
	// pathToPayload := "./nspGo-restConf/resconf-payload.json"
	// pathToPayload := "./nspGo-restConf/resconf-payload-100-svcs.json"
	// pathToPayload := "./nspGo-restConf/resconf-payload-200-svcs.json"
	pathToPayload := "./nspGo-restConf/resconf-payload-250-svcs.json"

	// pathToPayload := "./nspGo-restConf/resconf-payload-400-svcs.json"
	// pathToPayload := "./nspGo-restConf/resconf-payload-700-svc.json"
	// pathToPayload := "./nspGo-restConf/resconf-payload-1k-svc.json"
	// pathToPayload := "./nspGo-restConf/resconf-payload-2k-svc.json"
	// pathToPayload := "./nspGo-restConf/resconf-payload-4k-svc.json"

	file, err := os.Stat(pathToPayload)
	if err != nil {
		return
	}
	log.Info("Payload Size(bytes): ", file.Size())

	rc.ReadRestConfPayload(pathToPayload)
	restconfPayloadCreate := rc.Payload

	restconfAsync := true

	listOfExecutionTime := []time.Duration{}
	// iteration := int(math.Ceil(10000000 / float64(file.Size())))
	iteration := 16
	log.Info("Iteration: ", iteration)

	for i := 1; i <= iteration; i++ {

		startConcurrentIteration := time.Now()
		log.Info("Start Time Concurrent Request: ", startConcurrentIteration)
		log.Info("Number of Concurrent Requests: ", len(listOfNeId))
		var waitingGroupNeListIteration sync.WaitGroup
		waitingGroupNeListIteration.Add(len(listOfNeId))

		for j := 0; j < len(listOfNeId); j++ {
			go func(j int) {
				rc.PatchRestConfMdc(p.IpAdressIprc, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, listOfNeId[j], "/root", restconfPayloadCreate, restconfAsync)
				waitingGroupNeListIteration.Done()
			}(j)
		}
		waitingGroupNeListIteration.Wait()
		listOfExecutionTime = append(listOfExecutionTime, (time.Since(startConcurrentIteration)))
		log.Info("####################################")
		log.Info("Iteration: ", i)
		log.Info("Number of Targeted NE: ", len(listOfNeId))
		log.Info("Finished Concurrent Request")
		log.Info("RestConf Async: ", restconfAsync)
		log.Info("Payload Size Per NE (Bytes): ", float64(file.Size()))
		log.Info("Payload Size Per NE (KiloBytes): ", float64(file.Size())/1000)
		log.Info("Payload Size Per NE (MegaBytes): ", float64(file.Size())/1000000)
		log.Info("Elapsed Time(Iteration): ", time.Since(startConcurrentIteration))
		log.Info("####################################")

	}
	//Find Total Elapsed Time
	var totalElapsed time.Duration
	for _, v := range listOfExecutionTime {
		totalElapsed += v
	}
	//Find Min Max in Total Elapsed List
	var max time.Duration = listOfExecutionTime[0]
	var min time.Duration = listOfExecutionTime[0]
	for _, value := range listOfExecutionTime {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}

	log.Info(" ##################################")
	log.Info(" ##################################")
	log.Info(" Test-Case : RestConf MDC (8548) Call to get 4000 svc")
	log.Info(" Number of Targeted NE: ", len(listOfNeId))
	log.Info(" Payload File: ", pathToPayload)
	log.Info(" RestConf Async: ", restconfAsync)
	log.Info(" Total Iteration: ", (float64(iteration)))
	log.Info(" Payload Size Per NE Per Interation (MegaBytes): ", float64(file.Size())/1000000)
	log.Info(" Average Elapsed Time Per Interation (seconds): ", totalElapsed.Seconds()/float64(iteration))
	log.Info(" Average Elapsed Time Per Interation (minutes): ", totalElapsed.Minutes()/float64(iteration))
	log.Info(" Min Elapsed Time Per Interation: ", min.Seconds())
	log.Info(" Max Elapsed Time Per Interation: ", max.Seconds())
	log.Info(" ##################################")
	log.Info(" Total Payload Size Per NE (MegaBytes): ", (float64(iteration) * float64(file.Size()) / 1000000))
	log.Info(" Total Elapsed Time (seconds) : ", totalElapsed.Seconds())
	log.Info(" Total Elapsed Time : ", totalElapsed)
	// log.Info("Time Per Iteration :", listOfExecutionTime)
	log.Info(" ##################################")
	log.Info(" ##################################")

	p.RevokeRestToken()

}
