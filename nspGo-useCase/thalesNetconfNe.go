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

func ThalesNetconfNe() {
	// init class
	t := nspgotools.Tools{}

	// log level 4: info
	// log level 5: debug
	t.InitLogger("./logs/nspGo-useCase-ThalesNetconfNe.log", 4)

	template := (`
		<config>
		<configure xmlns="urn:nokia.com:sros:ns:yang:sr:conf">
			<service>
				<sdp>
					<sdp-id>1</sdp-id>
					<admin-state>enable</admin-state>
					<delivery-type>mpls</delivery-type>
					<ldp>true</ldp>
					<far-end>
						<ip-address>99.99.99.1</ip-address>
					</far-end>
				</sdp>{% for n in range(1,4091) %}
				<vpls>
					<service-name>service-{{n}}</service-name>
					<description>This Is PW-Labels-01 PlaceHolder-TiMOS-B-21.10.R1</description>
					<service-id>{{n}}</service-id>
					<customer>1</customer>
					<spoke-sdp>
						<sdp-bind-id>1:{{n}}</sdp-bind-id>
						<description>This Is MPLS-Params-01 PlaceHolder-1-10.2.31.2</description>
					</spoke-sdp>
					<sap>
						<sap-id>1/1/c1/1:{{n}}</sap-id>
						<description>This Is MPLS-Params-02 PlaceHolder-1-7750 SR-1</description>
					</sap>
				</vpls>{% endfor %}
			</service>
		</configure>
		</config>
		`)
	t.LoadTemplateJinja(template)

	// Write the generated payload to file, then get the size
	pathToPayload := "./nspGo-useCase/netconf-payload.xml"
	f, err := os.Create(pathToPayload)
	if err != nil {
		log.Error("Failed to create netconf payload file")
	}
	defer f.Close()
	f.WriteString((t.JinjaOutput))
	f.Sync()
	file, err := os.Stat(pathToPayload)
	if err != nil {
		return
	}

	// log.Info("Payload Size(bytes): ", file.Size())

	// templateManual := (`
	// 	<config>
	// 	<configure xmlns="urn:nokia.com:sros:ns:yang:sr:conf">
	// 		<service>
	// 			<sdp>
	// 				<sdp-id>1</sdp-id>
	// 				<admin-state>enable</admin-state>
	// 				<delivery-type>mpls</delivery-type>
	// 				<ldp>true</ldp>
	// 				<far-end>
	// 					<ip-address>99.99.99.1</ip-address>
	// 				</far-end>
	// 			</sdp>
	// 			<vpls>
	// 				<service-name>service-1</service-name>
	// 				<description>This Is PW-Labels-01 PlaceHolder-TiMOS-B-21.10.R1</description>
	// 				<service-id>1</service-id>
	// 				<customer>1</customer>
	// 				<spoke-sdp>
	// 					<sdp-bind-id>1:1</sdp-bind-id>
	// 					<description>This Is MPLS-Params-01 PlaceHolder-1-10.2.31.2</description>
	// 				</spoke-sdp>
	// 				<sap>
	// 					<sap-id>1/1/c1/1:1</sap-id>
	// 					<description>This Is MPLS-Params-02 PlaceHolder-1-7750 SR-1</description>
	// 				</sap>
	// 			</vpls>
	// 			<vpls>
	// 				<service-name>service-2</service-name>
	// 				<description>This Is PW-Labels-01 PlaceHolder-TiMOS-B-21.10.R1</description>
	// 				<service-id>2</service-id>
	// 				<customer>1</customer>
	// 				<spoke-sdp>
	// 					<sdp-bind-id>1:2</sdp-bind-id>
	// 					<description>This Is MPLS-Params-01 PlaceHolder-1-10.2.31.2</description>
	// 				</spoke-sdp>
	// 				<sap>
	// 					<sap-id>1/1/c1/1:2</sap-id>
	// 					<description>This Is MPLS-Params-02 PlaceHolder-1-7750 SR-1</description>
	// 				</sap>
	// 			</vpls>
	// 		</service>
	// 	</configure>
	// 	</config>
	// 	`)

	// fmt.Println(t.JinjaOutput)
	// t.NetconfClientEditConfig("10.2.31.2", "admin", "admin", templateManual)
	// log.Info("Start Time")

	// t.NetconfClientEditConfig("10.2.31.2", "admin", "admin", t.JinjaOutput)
	// t.NetconfClientEditCommit("10.2.31.2", "admin", "admin")

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
	// log.Info(inventoryJson)
	value := gjson.Get(inventoryJson, "nsp-inventory:output.data.#.active-ip-addresses.0") // get bof ip address
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

	listOfExecutionTime := []time.Duration{}
	// iteration := int(math.Ceil(10000000 / float64(file.Size())))
	iteration := 1
	log.Info("Iteration: ", iteration)

	for i := 1; i <= iteration; i++ {

		startConcurrentIteration := time.Now()
		log.Info("Start Time Concurrent Request: ", startConcurrentIteration)
		log.Info("Number of Concurrent Requests: ", len(listOfNeId))
		var waitingGroupNeListIteration sync.WaitGroup
		waitingGroupNeListIteration.Add(len(listOfNeId))

		for j := 0; j < len(listOfNeId); j++ {
			go func(j int) {
				// fmt.Println(listOfNeId[j])
				// t.NetconfClientEditConfig(listOfNeId[j], "admin", "admin", t.JinjaOutput)
				t.NetconfClientEditCommit(listOfNeId[j], "admin", "admin")
				waitingGroupNeListIteration.Done()
			}(j)
		}
		waitingGroupNeListIteration.Wait()
		listOfExecutionTime = append(listOfExecutionTime, (time.Since(startConcurrentIteration)))
		log.Info("####################################")
		log.Info("Iteration: ", i)
		log.Info("Number of Targeted NE: ", len(listOfNeId))
		log.Info("Finished Concurrent Request")
		log.Info("Payload Size Per NE (KiloBytes): ", float64(file.Size()))
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

	log.Info("##################################")
	log.Info("##################################")
	log.Info("Test-Case : Direct Netconf to NE Call")
	log.Info("Number of Targeted NE: ", len(listOfNeId))
	log.Info("Total Iteration: ", (float64(iteration)))
	log.Info("Payload Size Per NE Per Interation (MegaBytes): ", float64(file.Size())/1000000)
	log.Info("Average Elapsed Time Per Interation (seconds): ", totalElapsed.Seconds()/float64(iteration))
	log.Info("Average Elapsed Time Per Interation(minutes): ", totalElapsed.Minutes()/float64(iteration))
	log.Info("Min Elapsed Time Per Interation: ", min.Seconds())
	log.Info("Max Elapsed Time Per Interation: ", max.Seconds())
	log.Info("##################################")
	log.Info("Total Payload Size Per NE (MegaBytes): ", (float64(iteration) * float64(file.Size()) / 1000000))
	log.Info("Total Elapsed Time(seconds) : ", totalElapsed.Seconds())
	log.Info("Total Elapsed Time : ", totalElapsed)
	// log.Info("Time Per Iteration :", listOfExecutionTime)
	log.Info("##################################")
	log.Info("##################################")

	p.RevokeRestToken()
}
