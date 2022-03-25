package nspgousecase

import (
	"fmt"
	"strconv"
	"time"

	"github.com/schollz/progressbar/v3"
	log "github.com/sirupsen/logrus"

	nspgorestconf "local.com/nspgo/nspGo-restConf"
	nspgosession "local.com/nspgo/nspGo-session"
)

func ThalesLookupWithResourceManagerObtain(minSvcId int, maxSvcId int) {
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
	rc.LogLevel = 4
	rc.InitLogger()

	urlPathPwLabel := "data/nsp-resource-pool:resource-pools/numeric-resource-pools=PW-Label,thales/obtain-value-from-pool"
	payload := `{
		"nsp-resource-pool:input": {
			"total-number-of-resources": 2,
			"all-or-nothing": false,
			"owner": "CUST1",
			"reference": "ref-1",
			"confirmed": true
		}
	}`
	bytesPayload := []byte(payload)

	// im.NspIntentManagerGet(s.IpAdressNspOs, s.Token, s.Proxy.Enable, s.Proxy.ProxyAddress, urlPath, bytesPayload)
	// rc.NspRestconf8545Post(s.IpAdressNspOs, s.Token, s.Proxy.Enable, s.Proxy.ProxyAddress, urlPathPwLabel, bytesPayload)

	// minSvcId := 1
	// maxSvcId := 200000

	bar := progressbar.Default(int64(maxSvcId - minSvcId))
	listOfLookupTimePerIteration := []time.Duration{}

	for i := minSvcId; i <= maxSvcId; i++ {
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
		startLookupTime := time.Now()
		rc.NspRestconf8545Post(s.IpAdressNspOs, s.Token, s.Proxy.Enable, s.Proxy.ProxyAddress, urlPathPwLabel, bytesPayload)
		listOfLookupTimePerIteration = append(listOfLookupTimePerIteration, (time.Since(startLookupTime)))
		bar.Add(1)
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

func ThalesLookupWithResourceManagerRelease() {
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
	rc.LogLevel = 5
	rc.InitLogger()

	urlPathPwLabel := "data/nsp-resource-pool:resource-pools/numeric-resource-pools=PW-Label,thales/release-by-owner"
	payload := `{
		"nsp-resource-pool:input": {
			"owner": "CUST1"
		}
	}`
	bytesPayload := []byte(payload)

	listOfLookupTimePerIteration := []time.Duration{}
	startLookupTime := time.Now()

	rc.NspRestconf8545Post(s.IpAdressNspOs, s.Token, s.Proxy.Enable, s.Proxy.ProxyAddress, urlPathPwLabel, bytesPayload)

	listOfLookupTimePerIteration = append(listOfLookupTimePerIteration, (time.Since(startLookupTime)))

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
	log.Info("###############################################################")
	log.Info("Releasing All Reserved PW-Label Resources")
	log.Info("###############################################################")
	log.Info("Total Elapsed Time to Release All Reserved PW-Label Resources(seconds): ", totalElapsed)
	s.RevokeRestToken()
}
