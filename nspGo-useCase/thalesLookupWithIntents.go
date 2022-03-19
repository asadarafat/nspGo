package nspgousecase

import (
	"fmt"
	"strconv"

	"github.com/schollz/progressbar/v3"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/sjson"

	nspgointentmanger "local.com/nspgo/nspGo-intentManger"
	nspgorestconf "local.com/nspgo/nspGo-restConf"
	nspgosession "local.com/nspgo/nspGo-session"
)

func ThalesLookupWithIntent() {

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

	minIteration := 61002
	maxIteration := 70001
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
	// im.NspIntentManagerPost(p.IpAdressNspOs, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, "data/ibn:ibn", bytesJson)

	// rc.NspRestconf8545Del(p.IpAdressNspOs, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, "data/ibn:ibn/intent="+serviceId+",service", bytesJson)
	// rc.NspRestconf8545Post(p.IpAdressNspOs, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, "data/ibn:ibn", bytesJson)
	// rc.NspRestconf8545Get(p.IpAdressNspOs, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, "data/ibn:ibn/intent=service1000003,service", bytesJson)
	// rc.NspRestconf8545Del(p.IpAdressNspOs, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, "data/ibn:ibn/intent=service1000003,service", bytesJson)

	s.RevokeRestToken()

}
