package nspgousecase

import (
	"fmt"
	"strconv"
	"time"

	"github.com/schollz/progressbar/v3"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"

	nspgointentmanger "local.com/nspgo/nspGo-intentManger"
	nspgorestconf "local.com/nspgo/nspGo-restConf"
	nspgosession "local.com/nspgo/nspGo-session"
)

func ThalesLookupWithIntentCreateIntents() {

	const json = `{
		"ibn:intent": {
			"target": "service1000001",
			"intent-type": "service",
			"intent-type-version": 1,
			"required-network-state": "active",
			"ibn:intent-specific-data": {
				"service:service": {
					"pwlable1": "pwlable1-1000001",
					"pwlable2": "pwlable2-2000001"
				}
			}
		}
	}`

	// init class Session
	s := nspgosession.Session{}
	s.LoadConfig()

	s.EncodeUserName()
	log.Debug(s.EncodeUserName())

	s.GetRestToken()
	log.Debug("nsp.linetoken_NEW :", s.Token)
	fmt.Println(s.Token)

	// init class RestConf
	rc := nspgorestconf.RestConf{}
	rc.LogLevel = 3
	rc.InitLogger()

	// init class IntentManager
	im := nspgointentmanger.IntentManager{}
	im.LogLevel = 5
	im.InitLogger()

	minIteration := 1000000
	maxIteration := 1000001
	// batch := 100

	bar := progressbar.Default(int64(maxIteration - minIteration))

	for i := minIteration; i <= maxIteration; i++ {
		// var waitingGroupNeListIteration sync.WaitGroup
		// waitingGroupNeListIteration.Add(maxIteration - minIteration)

		// for j := i; j < i+batch; j += batch {
		// 	go func(j int) {
		// println("test" + strconv.Itoa(1000000+j))

		if i%10000 == 0 { //
			s.RevokeRestToken()

			log.Debug("nsp.nspOsIP :", s.IpAdressNspOs)
			log.Debug("nsp.nspIprcIP :", s.IpAdressIprc)
			log.Debug("nsp.Username :", s.Username)
			log.Debug("nsp.Password :", s.Password)
			log.Debug("nsp.linetoken :", s.Token)

			s.EncodeUserName()
			log.Debug(s.EncodeUserName())

			s.GetRestToken()
			log.Debug("nsp.linetoken_NEW :", s.Token)
			fmt.Println(s.Token)

		}

		serviceId := "service" + strconv.Itoa(1000000+i)
		pwlable1Id := "pwlable01Id-" + strconv.Itoa(1000000+i)
		pwlable2Id := "pwlable02Id-" + strconv.Itoa(2000000+i)

		// serviceId := "service" + strconv.Itoa(1000000+j)
		// pwlable1Id := "pwlable01Id-" + strconv.Itoa(1000000+j)
		// pwlable2Id := "pwlable02Id-" + strconv.Itoa(2000000+j)

		setJson________, _ := sjson.Set(json, "ibn:intent.target", serviceId)
		setJsonPwlable1, _ := sjson.Set(setJson________, "ibn:intent.ibn:intent-specific-data.service:service.pwlable1", pwlable1Id)
		setJsonPwlable2, _ := sjson.Set(setJsonPwlable1, "ibn:intent.ibn:intent-specific-data.service:service.pwlable2", pwlable2Id)

		bytesJson := []byte(setJsonPwlable2)
		im.NspIntentManagerPost(s.IpAdressNspOs, s.Token, s.Proxy.Enable, s.Proxy.ProxyAddress, "data/ibn:ibn", bytesJson)
		// 	waitingGroupNeListIteration.Done()
		// }(j)

		bar.Add(1)
		// }
		// time.Sleep(1 * time.Microsecond)
	}

	s.RevokeRestToken()

}

func ThalesLookupWithIntentGetIntents() {
	// init class Session
	s := nspgosession.Session{}
	s.LoadConfig()

	s.EncodeUserName()
	log.Debug(s.EncodeUserName())

	s.GetRestToken()
	log.Debug("nsp.linetoken_NEW :", s.Token)
	fmt.Println(s.Token)

	// init class IntentManager
	im := nspgointentmanger.IntentManager{}
	im.LogLevel = 4
	im.InitLogger()

	minSvcId := 1
	maxSvcId := 10

	payload := []byte("")

	listOfLookupTimePerIteration := []time.Duration{}

	for i := minSvcId; i <= maxSvcId; i++ {
		startLookupTime := time.Now()
		log.Debug("Lookup Time Start ", startLookupTime)

		urlPath := "data/ibn:ibn/intent=service" + strconv.Itoa(1000000+i) + ",service"
		// im.NspIntentManagerGet(s.IpAdressNspOs, s.Token, s.Proxy.Enable, s.Proxy.ProxyAddress, urlPath, payload)
		log.Info(gjson.Get(im.NspIntentManagerGet(s.IpAdressNspOs, s.Token, s.Proxy.Enable, s.Proxy.ProxyAddress, urlPath, payload), "ibn:intent.intent-specific-data.service:service"))

		listOfLookupTimePerIteration = append(listOfLookupTimePerIteration, (time.Since(startLookupTime)))

	}

	//Find Total Elapsed Time at MidPoint
	var totalElapsedMidPoint time.Duration
	for w := 0; w <= len(listOfLookupTimePerIteration)/2; w++ {
		totalElapsedMidPoint += listOfLookupTimePerIteration[w]
	}

	//Find Total Elapsed Time
	var totalElapsed time.Duration
	for _, v := range listOfLookupTimePerIteration {
		totalElapsed += v
	}

	//Find Min Max in Total Elapsed List
	var max time.Duration = listOfLookupTimePerIteration[0]
	var min time.Duration = listOfLookupTimePerIteration[0]
	for _, value := range listOfLookupTimePerIteration {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	log.Info("Min Elapsed Time Per Iteration (second): ", min.Seconds())
	log.Info("Max Elapsed Time Per Iteration(second): ", max.Seconds())
	log.Info("Total Iteration: ", maxSvcId-minSvcId+1)
	log.Info("Total Elapsed Time MidPoint at "+strconv.Itoa(len(listOfLookupTimePerIteration)/2)+" Iteration (seconds): ", totalElapsedMidPoint)
	log.Info("Total Elapsed Time(seconds): ", totalElapsed)
	s.RevokeRestToken()
}
