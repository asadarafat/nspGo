package main

import (
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	nspgoipoptim "local.com/nspgo/nspGo-ipOptim"
	nspgosession "local.com/nspgo/nspGo-session"
	nspgotopoviewer "local.com/nspgo/nspGo-topoViewer"
)

func main() {
	// graph1 := nsptopoviewer.CyGraph{}

	// content, err := ioutil.ReadFile("../../ietfNetwork.json")
	// if err != nil {
	// 	log.Fatal("Error when opening file: ", err)
	// }
	// filePath, _ := os.Getwd()
	// filePath = (filePath + "../../../vis-library/colajs-asad-graph/data-cytoMarshall.json")
	// graph1.DumpIetfNetworkToCyGraph(content, nsptopoviewer.IetfNetworkStruct{}, filePath)

	s := nspgosession.Session{}

	s.LogLevel = 5
	s.InitLogger()
	s.LoadConfig()

	s.EncodeUserName()
	log.Info(s.EncodeUserName())

	s.GetRestToken()

	// // Get IETF from NSP Topology
	nspgoipoptim := nspgoipoptim.IpOptim{}
	nspgoipoptim.LogLevel = 5
	nspgoipoptim.InitLogger()
	var dummyPayload []byte
	ietfRawFile := nspgoipoptim.IpoV4IetfTeNetworksGet(s.IpAdressNspOs, s.Token, s.Proxy.Enable, s.Proxy.ProxyAddress, dummyPayload)

	// Draw MultiLayer Topology
	nextUiGo := nspgotopoviewer.NextUiGo{}
	nextUiGo.LogLevel = 5

	nextUiGo.NextUiUnmarshalIetfNetworkModel([]byte(ietfRawFile), nspgotopoviewer.IetfNetworkStruct{})

	s.RevokeRestToken()

	// Load NetSupFile and Make Physical Topology
	filePath, _ := os.Getwd()
	filePath = (filePath + "/nspGo-topoViewer/rawTopoFile/netSupPhysical.json")
	netSupFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	nextUiGo.NextUiUnmarshalNetSupPhysicalModel(netSupFile, nspgotopoviewer.NetSupPhysicalStruct{})

	nextUiGo.NextUiAppendInterLayerLinks(nextUiGo.Topology.Nodes)
	nextUiGo.NextUiHttpServer(nextUiGo.NextUiMarshal())

}
