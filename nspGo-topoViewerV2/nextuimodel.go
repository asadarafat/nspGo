package nspgotopoviewernew

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	nspgotools "local.com/nspgo/nspGo-tools"
)

type NextUiGo struct {
	Topology NextTopology
	LogLevel uint32
	DataHtml TopologyDataHtml
}

type NextTopology struct {
	Nodes []NextNode `json:"nodes"`
	Links []NextLink `json:"links"`
}

type NextNode struct {
	Id          string `json:"id"`
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	NetworkUuid string `json:"networkUuid"`
	LayerName   string `json:"layerName"`
	Group       string `json:"group"`

	Type string `json:"type"`
}

type NextLink struct {
	Id             string `json:"id"`
	Uuid           string `json:"uuid"`
	Source         string `json:"source"`
	SourceEndpoint string `json:"source_endpoint"`
	Target         string `json:"target"`
	TargetEndpoint string `json:"target_endpoint"`
	NetworkUuid    string `json:"networkUuid"`
	Group          string `json:"group"`
	Type           string `json:"type"`
}

type TopologyDataHtml struct {
	Name string
	Data string
}

func (nextGo *NextUiGo) InitLogger() {
	// init logConfig
	toolLogger := nspgotools.Tools{}
	toolLogger.InitLogger("./logs/nspGo-nextUi.log", nextGo.LogLevel)
}

// func InitLogger() {
// 	// init logConfig
// 	toolLogger := nspgotools.Tools{}
// 	toolLogger.InitLogger("./logs/nspGo-restconf.log", 5)
// }

func (nextGo *NextUiGo) ReadRawTopoJsonFile(file string) []byte {
	rawFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("unable to read file: %v", err)
	}
	log.Info("Raw Topology Loaded")
	return rawFile
	//fmt.Println(string(body))
}

func (nextGo *NextUiGo) NextUiUnmarshalIetfNetworkModel(rawTopoFile []byte, ietfModel IetfNetworkStruct) {
	json.Unmarshal(rawTopoFile, &ietfModel)
	log.Debugf("Lenght of Data Network Array: ", (len(ietfModel.Response.Data.Network)))
	var node NextNode
	var link NextLink

	for i := 0; i < len(ietfModel.Response.Data.Network); i++ {

		for j := 0; j < len(ietfModel.Response.Data.Network[i].Link); j++ {
			var SourceNodeID = ietfModel.Response.Data.Network[i].Link[j].Source.SourceNode
			var DestNodeID = ietfModel.Response.Data.Network[i].Link[j].Destination.DestNode

			for k := 0; k < len(ietfModel.Response.Data.Network[i].Node); k++ {
				if ietfModel.Response.Data.Network[i].Node[k].NodeID == SourceNodeID {
					SourceNodeName := ietfModel.Response.Data.Network[i].Node[k].TeNodeAugment.Te.TeNodeID.DottedQuad.String
					node.Id = ("L3-" + SourceNodeName)
					node.Name = (SourceNodeName)
					SourceNodeUuid := ietfModel.Response.Data.Network[i].Node[k].NodeID
					node.Uuid = SourceNodeUuid
					node.LayerName = "L3-Layer"
					node.Group = "IgpLayer"
					node.NetworkUuid = ietfModel.Response.Data.Network[i].NetworkID
					//  aarafat-tag: this will result duplicated nodes to be appended
					nextGo.Topology.Nodes = append(nextGo.Topology.Nodes, node)

					link.Source = ("L3-" + SourceNodeName)
					link.SourceEndpoint = ("L3-" + SourceNodeName)
				}

				if ietfModel.Response.Data.Network[i].Node[k].NodeID == DestNodeID {
					DestNodeName := ietfModel.Response.Data.Network[i].Node[k].TeNodeAugment.Te.TeNodeID.DottedQuad.String
					node.Id = ("L3-" + DestNodeName)
					node.Name = (DestNodeName)
					DestNodeUuid := ietfModel.Response.Data.Network[i].Node[k].NodeID
					node.Uuid = DestNodeUuid
					node.LayerName = "L3-Layer"
					node.Group = "IgpLayer"
					node.NetworkUuid = ietfModel.Response.Data.Network[i].NetworkID
					//  aarafat-tag: this will result duplicated nodes to be appended
					nextGo.Topology.Nodes = append(nextGo.Topology.Nodes, node)

					link.Target = ("L3-" + DestNodeName)
					link.TargetEndpoint = ("L3-" + DestNodeName)

				}

			}
			LinkUuid := ietfModel.Response.Data.Network[i].Link[j].LinkID
			link.Id = ("from-" + link.Source + "-to-" + link.Target)
			link.Uuid = LinkUuid

			link.NetworkUuid = ietfModel.Response.Data.Network[i].NetworkID
			nextGo.Topology.Links = append(nextGo.Topology.Links, link)

		}
	}
	// remove duplicated node from Nodelist
	nextGo.RemovedDuplicatedNode(nextGo.Topology.Nodes, "L3-Layer")
}

func contains(nodesList []NextNode, nodeId string, nodeNetworkUuid string, layerName string) bool {
	for _, node := range nodesList {
		if (node.LayerName == layerName) && (node.Id == nodeId) && (node.NetworkUuid == nodeNetworkUuid) {
			return true
		}
	}
	return false
}

func (nextGo *NextUiGo) RemovedDuplicatedNode(oldNodesList []NextNode, layerName string) {
	newNodesList := []NextNode{}
	for _, node := range oldNodesList {
		fmt.Println(node)
		if contains(newNodesList, node.Id, node.NetworkUuid, layerName) == false {
			newNodesList = append(newNodesList, node)
		}
	}
	nextGo.Topology.Nodes = nil
	nextGo.Topology.Nodes = newNodesList
}

func (nextGo *NextUiGo) NextUiUnmarshalNetSupL2Model(rawTopoFile []byte, netSupModel NetSupL2Struct) {
	json.Unmarshal(rawTopoFile, &netSupModel)
	var node NextNode
	var link NextLink
	for i := 0; i < len(netSupModel.Response.Data); i++ {
		// log.Debugf("sourceNode: ", netSupModel.Response.Data[i].Endpoints[0].ParentNeID)
		// log.Debugf("destNode: ", netSupModel.Response.Data[i].Endpoints[1].ParentNeID)

		// append pyhsical source Node in pyhsical Topology
		SourceNodeName := netSupModel.Response.Data[i].Endpoints[0].ParentNeID
		node.Id = ("L2-" + SourceNodeName)
		node.Name = (SourceNodeName)
		node.LayerName = "L2-Layer"
		node.Group = "PhysicalLayer"
		nextGo.Topology.Nodes = append(nextGo.Topology.Nodes, node)

		// append pyhsical destination Node in pyhsical Topology
		DestNodeName := netSupModel.Response.Data[i].Endpoints[1].ParentNeID
		node.Id = ("L2-" + DestNodeName)
		node.Name = (DestNodeName)
		node.LayerName = "L2-Layer"
		node.Group = "PhysicalLayer"
		nextGo.Topology.Nodes = append(nextGo.Topology.Nodes, node)

		// append pyhsical link in pyhsical Topology
		link.Source = ("L2-" + SourceNodeName)
		link.SourceEndpoint = netSupModel.Response.Data[i].Endpoints[0].Name
		link.Target = ("L2-" + DestNodeName)
		link.TargetEndpoint = netSupModel.Response.Data[i].Endpoints[1].Name
		// link.Id = strconv.Itoa(i)
		link.Id = netSupModel.Response.Data[i].ID
		nextGo.Topology.Links = append(nextGo.Topology.Links, link)
	}
	nextGo.RemovedDuplicatedNode(nextGo.Topology.Nodes, "L2-Layer")

}

func (nextGo *NextUiGo) NextUiAppendInterLayerLinks(Nodes []NextNode, layerSourceName string, layerDestinationName string) {
	// u, err := json.Marshal(NextNode{Id: nextGo.Topology.NextNodes[0].Name})
	var link NextLink
	var linkId int = 1

	for i := 0; i < len(Nodes); i++ {

		if !strings.Contains(Nodes[i].Id, layerSourceName+"-") { // find L3 layer nodes //layerSourceName=L2
			link.Source = (Nodes[i].Id) // L3 nodes
			link.SourceEndpoint = layerDestinationName
			for j := 0; j < len(Nodes); j++ {
				// if Nodes[j].Name == ("L2-" + Nodes[i].Name) { // find L2 Layer node with same name with L3 Layer node
				if Nodes[j].Name == (Nodes[i].Name) { // find L2 Layer node with same name with L3 Layer node

					link.Target = (Nodes[j].Id)
					link.TargetEndpoint = layerSourceName

					link.Id = strconv.Itoa(linkId + j)
				}
			}
			link.Group = ("interTopoLayerLink")
			nextGo.Topology.Links = append(nextGo.Topology.Links, link)
		}
	}
}

func (nextGo *NextUiGo) NextUiMarshal() []byte {
	// u, err := json.Marshal(NextNode{Id: nextGo.Topology.NextNodes[0].Name})
	jsonBytes, err := json.Marshal(NextTopology{
		Nodes: nextGo.Topology.Nodes,
		Links: nextGo.Topology.Links})
	if err != nil {
		panic(err)
	}
	log.Debugf("NextUiMarshal Result:", string(jsonBytes))

	return jsonBytes
}

func (nextGo *NextUiGo) NextUiHttpServer(marshaledNextuiTopo []byte) {

	type nodeData struct {
		Name string
		// Name []byte
	}
	std1 := nodeData{string(marshaledNextuiTopo)}
	tmp1 := template.New("Template_1")
	tmp1, _ = tmp1.Parse("var data = {{.Name}}")

	filePath, _ := (os.Getwd())
	file, err := os.Create(filePath + "/nspGo-topoViewerV2/assets/static/js/nextDataAsad.js")

	log.Debugf("fileOutput: ", file)

	if err == nil {
		tmp1.Execute(file, std1)
	}

	// read the file again to replace "&#34;" with single quoute "
	// aaraafat-tag: code need to be optimised, no need to read/write again file nextDataAsad.js
	input, err := ioutil.ReadFile(filePath + "/nspGo-topoViewerV2/html-template/static/js/nextDataAsad.js")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	output := bytes.Replace(input, []byte("&#34;"), []byte(`"`), -1)

	if err = ioutil.WriteFile(filePath+"/nspGo-topoViewerV2/html-template/static/js/nextDataAsad.js", output, 0666); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// log.Debugf("filePath: ", filePath+"/nspGo-topoViewerV2/html-template/static/")
	http.Handle("/", http.FileServer(http.Dir(filePath+"/nspGo-topoViewerV2/html-template")))
	http.ListenAndServe(":8080", nil)
}

func (nextGo *NextUiGo) NextUiGenerateNextDataJs(marshaledNextuiTopo []byte) {

	type nodeData struct {
		Name string
		// Name []byte
	}

	// define nextData.js template
	std1 := nodeData{string(marshaledNextuiTopo)}
	tmp1 := template.New("Template_1")
	tmp1, _ = tmp1.Parse("var data = {{.Name}}")

	// write nextData.js template
	filePath, _ := (os.Getwd())
	file, err := os.Create(filePath + "/nspGo-topoViewerV2/assets/js/nextData.js")
	log.Debugf("fileOutput: ", file)
	if err == nil {
		tmp1.Execute(file, std1)
	}

	log.Debugf("filePathPrefix: ", filePath)
	log.Debugf("filePathFull: ", filePath+"/nspGo-topoViewerV2/assets/js/nextData.js")
	log.Debugf("fileOutput: ", file)

	// if err == nil {
	// 	// tmp1.Execute(file, std1)
	// }

	// read the file again to replace "&#34;" with single quoute "
	// aaraafat-tag: code need to be optimised, no need to read/write again file nextDataAsad.js
	input, err := ioutil.ReadFile(filePath + "/nspGo-topoViewerV2/assets/js/nextData.js")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	output := bytes.Replace(input, []byte("&#34;"), []byte(`"`), -1)

	if err = ioutil.WriteFile(filePath+"/nspGo-topoViewerV2/assets/js/nextData.js", output, 0666); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	if err = ioutil.WriteFile(filePath+"/nspGo-topoViewerV2/assets/js/nextData-master.js", output, 0666); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
